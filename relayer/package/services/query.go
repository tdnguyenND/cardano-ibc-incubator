package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/cardano/relayer/v1/constant"
	"slices"

	"github.com/cardano/relayer/v1/package/mithril/dtos"
	"github.com/cardano/relayer/v1/package/services/helpers"
	"github.com/cardano/relayer/v1/relayer/chains/cosmos/mithril"
	"os"
)

func (gw *Gateway) QueryIBCHeader(ctx context.Context, h int64, cs *mithril.ClientState) (*mithril.MithrilHeader, error) {
	cardanoTxsSetSnapshot, err := gw.MithrilService.GetCardanoTransactionsSetSnapshot()
	if err != nil {
		return nil, err
	}
	snapshotIdx := slices.IndexFunc(cardanoTxsSetSnapshot, func(c dtos.CardanoTransactionSetSnapshot) bool { return c.Beacon.ImmutableFileNumber == uint64(h) })
	if snapshotIdx == -1 {
		return nil, errors.New(fmt.Sprintf("Could not find snapshot with height %d", h))
	}

	snapshot := &cardanoTxsSetSnapshot[snapshotIdx]
	snapshotCertificate, err := gw.MithrilService.GetCertificateByHash(snapshot.CertificateHash)
	if err != nil {
		return nil, err
	}
	if cs.CurrentEpoch < snapshot.Beacon.Epoch {
		//fmt.Printf("Client State has Current epoch: %v, ", cs.CurrentEpoch)
		//fmt.Printf("Snapshot has epoch: %v \n", snapshot.Beacon.Epoch)
		return gw.QueryIBCGenesisCertHeader(ctx, int64(cs.CurrentEpoch+1))
	}

	mithrilStakeDistributionList, err := gw.MithrilService.GetListMithrilStakeDistributions()
	if err != nil {
		return nil, err
	}

	mithrilStakeDistributionIdx := slices.IndexFunc(mithrilStakeDistributionList, func(c dtos.MithrilStakeDistribution) bool { return c.Epoch == snapshot.Beacon.Epoch })
	if mithrilStakeDistributionIdx == -1 {
		return nil, errors.New(fmt.Sprintf("Could not find stake distribution with epoch %d", snapshot.Beacon.Epoch))
	}
	mithrilStakeDistribution := mithrilStakeDistributionList[mithrilStakeDistributionIdx]
	mithrilDistributionCertificate, err := gw.MithrilService.GetCertificateByHash(mithrilStakeDistribution.CertificateHash)
	if err != nil {
		return nil, err
	}

	mithrilHeader := mithril.MithrilHeader{
		MithrilStakeDistribution:            helpers.ConvertMithrilStakeDistribution(mithrilStakeDistribution, *mithrilDistributionCertificate),
		MithrilStakeDistributionCertificate: helpers.ConvertMithrilStakeDistributionCertificate(mithrilStakeDistribution, *mithrilDistributionCertificate),
		TransactionSnapshot: &mithril.CardanoTransactionSnapshot{
			SnapshotHash:    snapshot.Hash,
			MerkleRoot:      snapshot.MerkleRoot,
			CertificateHash: snapshot.CertificateHash,
			Epoch:           snapshot.Beacon.Epoch,
			Height: &mithril.Height{
				MithrilHeight: snapshot.Beacon.ImmutableFileNumber,
			},
		},
		TransactionSnapshotCertificate: helpers.ConvertMithrilStakeDistributionCertificate(dtos.MithrilStakeDistribution{
			Hash:            snapshot.Hash,
			Epoch:           snapshot.Beacon.Epoch,
			CertificateHash: snapshot.CertificateHash,
			CreatedAt:       snapshot.CreatedAt,
		}, *snapshotCertificate),
	}

	return &mithrilHeader, nil
}

func (gw *Gateway) QueryNewMithrilClient() (*mithril.ClientState, *mithril.ConsensusState, error) {
	currentEpochSettings, err := gw.MithrilService.GetEpochSetting()
	if err != nil {
		return nil, nil, err
	}
	mithrilStakeDistributionsList, err := gw.MithrilService.GetListMithrilStakeDistributions()
	if err != nil {
		return nil, nil, err
	}
	if len(mithrilStakeDistributionsList) == 0 {
		return nil, nil, fmt.Errorf("GetListMithrilStakeDistributions returned empty list")
	}
	mithrilDistribution := mithrilStakeDistributionsList[0]
	fcCertificateMsd, err := gw.MithrilService.GetCertificateByHash(mithrilDistribution.CertificateHash)
	if err != nil {
		return nil, nil, err
	}
	certificateList, err := gw.MithrilService.GetListCertificates()
	if err != nil {
		return nil, nil, err
	}
	latestCertificateMsd := dtos.CertificateOverall{}
	idx := slices.IndexFunc(certificateList, func(c dtos.CertificateOverall) bool { return c.Epoch == mithrilDistribution.Epoch })
	if idx == -1 {
		return nil, nil, fmt.Errorf("could not find certificate with epoch %d", mithrilDistribution.Epoch)
	}
	latestCertificateMsd = certificateList[idx]

	listSnapshots, err := gw.MithrilService.GetListSnapshots()
	if err != nil {
		return nil, nil, err
	}
	if len(listSnapshots) == 0 {
		return nil, nil, fmt.Errorf("GetListSnapshots returned empty list")
	}
	latestSnapshot := listSnapshots[0]
	latestSnapshotCertificate, err := gw.MithrilService.GetCertificateByHash(latestSnapshot.CertificateHash)
	if err != nil {
		return nil, nil, err
	}

	phifFraction := helpers.FloatToFraction(currentEpochSettings.Protocol.PhiF)
	clientState := &mithril.ClientState{
		ChainId: os.Getenv(constant.CardanoChainNetworkMagic),
		LatestHeight: &mithril.Height{
			MithrilHeight: latestSnapshotCertificate.Beacon.ImmutableFileNumber,
		},
		FrozenHeight: &mithril.Height{
			MithrilHeight: 0,
		},
		CurrentEpoch:   currentEpochSettings.Epoch,
		TrustingPeriod: 0,
		ProtocolParameters: &mithril.MithrilProtocolParameters{
			K: currentEpochSettings.Protocol.K,
			M: currentEpochSettings.Protocol.M,
			PhiF: mithril.Fraction{
				Numerator:   phifFraction.Numerator,
				Denominator: phifFraction.Denominator,
			},
		},
		UpgradePath: nil,
	}
	timestamp := fcCertificateMsd.Metadata.SealedAt.UnixNano()
	consensusState := &mithril.ConsensusState{
		Timestamp:            uint64(timestamp),
		FcHashLatestEpochMsd: mithrilDistribution.CertificateHash,
		LatestCertHashMsd:    latestCertificateMsd.Hash,
		FcHashLatestEpochTs:  mithrilDistribution.CertificateHash,
		LatestCertHashTs:     latestSnapshot.CertificateHash,
	}
	return clientState, consensusState, nil
}

func (gw *Gateway) QueryIBCGenesisCertHeader(ctx context.Context, epoch int64) (*mithril.MithrilHeader, error) {
	mithrilStakeDistributionList, err := gw.MithrilService.GetListMithrilStakeDistributions()
	if err != nil {
		return nil, err
	}

	mithrilStakeDistributionIdx := slices.IndexFunc(mithrilStakeDistributionList, func(c dtos.MithrilStakeDistribution) bool { return c.Epoch == uint64(epoch) })
	if mithrilStakeDistributionIdx == -1 {
		return nil, errors.New(fmt.Sprintf("Could not find stake distribution with epoch %d", epoch))
	}
	mithrilStakeDistribution := mithrilStakeDistributionList[mithrilStakeDistributionIdx]
	mithrilDistributionCertificate, err := gw.MithrilService.GetCertificateByHash(mithrilStakeDistribution.CertificateHash)
	if err != nil {
		return nil, err
	}

	cardanoTxsSetSnapshot, err := gw.MithrilService.GetCardanoTransactionsSetSnapshot()
	if err != nil {
		return nil, err
	}

	cardanoTxsSetSnapshotReverse := slices.Clone(cardanoTxsSetSnapshot)
	slices.Reverse(cardanoTxsSetSnapshotReverse)
	firstSnapshotIdx := slices.IndexFunc(cardanoTxsSetSnapshotReverse, func(c dtos.CardanoTransactionSetSnapshot) bool { return c.Beacon.Epoch == uint64(epoch) })
	if firstSnapshotIdx == -1 {
		return nil, errors.New(fmt.Sprintf("Could not find snapshot with epoch %d", epoch))
	}
	firstSnapshot := &cardanoTxsSetSnapshotReverse[firstSnapshotIdx]
	snapshotCertificate, _ := gw.MithrilService.GetCertificateByHash(firstSnapshot.CertificateHash)
	// TODO: There is an issue that cannot get first trx snapshots with epoch if there are too many tx snapshots
	mithrilHeader := mithril.MithrilHeader{
		MithrilStakeDistribution:            helpers.ConvertMithrilStakeDistribution(mithrilStakeDistribution, *mithrilDistributionCertificate),
		MithrilStakeDistributionCertificate: helpers.ConvertMithrilStakeDistributionCertificate(mithrilStakeDistribution, *mithrilDistributionCertificate),
		TransactionSnapshot: &mithril.CardanoTransactionSnapshot{
			SnapshotHash:    firstSnapshot.Hash,
			MerkleRoot:      firstSnapshot.MerkleRoot,
			CertificateHash: firstSnapshot.CertificateHash,
			Epoch:           firstSnapshot.Beacon.Epoch,
			Height: &mithril.Height{
				MithrilHeight: firstSnapshot.Beacon.ImmutableFileNumber,
			},
		},
		TransactionSnapshotCertificate: helpers.ConvertMithrilStakeDistributionCertificate(dtos.MithrilStakeDistribution{
			Hash:            firstSnapshot.Hash,
			Epoch:           firstSnapshot.Beacon.Epoch,
			CertificateHash: firstSnapshot.CertificateHash,
			CreatedAt:       firstSnapshot.CreatedAt,
		}, *snapshotCertificate),
	}

	return &mithrilHeader, nil
}
