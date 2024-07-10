import { UTxO } from '@dinhbx/lucid-custom';

export type UnsignedSendPacketEscrowDto = {
  channelUTxO: UTxO;
  connectionUTxO: UTxO;
  clientUTxO: UTxO;
  spendChannelRefUTxO: UTxO;
  spendTransferModuleUTxO: UTxO;
  transferModuleUTxO: UTxO;

  encodedSpendChannelRedeemer: string;
  encodedSpendTransferModuleRedeemer: string;
  encodedUpdatedChannelDatum: string;

  transferAmount: bigint;
  senderAddress: string;
  receiverAddress: string;

  spendChannelAddress: string;
  channelTokenUnit: string;
  transferModuleAddress: string;
  denomToken: string;
};
