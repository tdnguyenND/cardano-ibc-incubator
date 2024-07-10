import { UTxO } from '@dinhbx/lucid-custom';

export type UnsignedSendPacketBurnDto = {
  channelUTxO: UTxO;
  connectionUTxO: UTxO;
  clientUTxO: UTxO;
  spendChannelRefUTxO: UTxO;
  spendTransferModuleUTxO: UTxO;
  transferModuleUTxO: UTxO;
  mintVoucherRefUtxo: UTxO;
  senderVoucherTokenUtxo: UTxO;

  encodedSpendChannelRedeemer: string;
  encodedSpendTransferModuleRedeemer: string;
  encodedMintVoucherRedeemer: string;
  encodedUpdatedChannelDatum: string;

  channelTokenUnit: string;
  voucherTokenUnit: string;

  senderAddress: string;
  receiverAddress: string;
  transferAmount: bigint;
  denomToken: string;
};
