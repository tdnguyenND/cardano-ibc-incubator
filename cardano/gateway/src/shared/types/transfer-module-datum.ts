import {Data} from '@dinhbx/lucid-custom';

export type TransferModuleDatum = {
    denom_trace: Map<string, string>;
};

export async function encodeTransferModuleDatum(transferModuleDatum: TransferModuleDatum, Lucid: typeof import('@dinhbx/lucid-custom')) {
    const {Data} = Lucid;

    const TransferModuleDatumSchema = Data.Object({
        denom_trace: Data.Map(Data.String(), Data.String()),
    });
    type TTransferModuleDatum = Data.Static<typeof TransferModuleDatumSchema>;
    const TTransferModuleDatum = TransferModuleDatumSchema as unknown as TransferModuleDatum;
    return Data.to(transferModuleDatum, TTransferModuleDatum);
}

export async function decodeTransferModuleDatum(transferModuleDatum: string, Lucid: typeof import('@dinhbx/lucid-custom')) {
    const {Data} = Lucid;
    const TransferModuleDatumSchema = Data.Object({
        denom_trace: Data.Map(Data.String(), Data.String()),
    });
    type TTransferModuleDatum = Data.Static<typeof TransferModuleDatumSchema>;
    const TTransferModuleDatum = TransferModuleDatumSchema as unknown as TransferModuleDatum;
    Data.from(transferModuleDatum, TTransferModuleDatum);
    return Data.from(transferModuleDatum, TTransferModuleDatum);
}