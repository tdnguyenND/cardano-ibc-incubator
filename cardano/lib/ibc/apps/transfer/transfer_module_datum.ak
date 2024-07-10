use aiken/dict.{Dict}

pub type TransferModuleDatum {
    denom_trace: Dict<ByteArray, ByteArray>,
}
