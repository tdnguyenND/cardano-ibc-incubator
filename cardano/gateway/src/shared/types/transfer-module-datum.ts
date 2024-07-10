import {Data} from '@dinhbx/lucid-custom';

export type TransferModuleDatum = {
    denom_trace: Map<string, string>;
};

export async function encodeTransferModuleDatum(transferModuleDatum: TransferModuleDatum, Lucid: typeof import('@dinhbx/lucid-custom')) {
    const {Data} = Lucid;
    const TransferModuleDatumSchema = Data.Object({
        denom_trace: Data.Map(Data.Bytes(), Data.Bytes()),
    });
    const TTransferModuleDatum = TransferModuleDatumSchema as unknown as TransferModuleDatum;
    return Data.to(transferModuleDatum, TTransferModuleDatum);
}

export async function decodeTransferModuleDatum(transferModuleDatum: string, Lucid: typeof import('@dinhbx/lucid-custom')) {
    const {Data} = Lucid;
    const TransferModuleDatumSchema = Data.Object({
        denom_trace: Data.Map(Data.Bytes(), Data.Bytes()),
    });
    const TTransferModuleDatum = TransferModuleDatumSchema as unknown as TransferModuleDatum;
    return Data.from(transferModuleDatum, TTransferModuleDatum);
}