# Cardano IBC Incubator
This project is working towards a bridge implementation to allow exchange of information from a Cardano blockchain to Cosmos SDK based blockchains.

It follows the [inter-blockchain communication protocol](https://github.com/cosmos/ibc) and is trying to achieve full compliance with the parts of the specification identified necessary for the developed framework.

## :heavy_exclamation_mark: Disclaimer
Please be aware that this is an incubator project and by this means it is neither complete nor sufficiently tested at the current point in time to be used for production grade operation of a bridge. So the use of the source code and software artifacts in this repository are subject to your own discretion and risk.

:heavy_exclamation_mark: The software withing this repository is provided to you on an "as is" and "as available" basis.

While we strive for high functionality and user satisfaction and endeavour to maintain reliability and accuracy, unforeseen issues may arise due to the experimental nature of this project.

## :eyes: Overview
This repository is subdivided into three main folders:
- `cardano`: Contains all Cardano related source code that are part of the bridge as well as some facilities for bringing up a local Cardano blockchain for test and development purposes. It also contains the Aiken based Tendermint Light Client and IBC primitives implementation.
- `cosmos`: Contains all Cosmos SDK related source code including the Cardano light client (or thin client) implementation running on the Cosmos chain. The folder was scaffolded via [Ignite CLI](https://docs.ignite.com/) with [Cosmos SDK 0.50](https://github.com/cosmos/cosmos-sdk).
- `relayer`: Contains all relayer related source code. Forked from https://github.com/cosmos/relayer

## :rocket: Getting Started

### Prerequisites

The following components are required to run the project:

- [Docker](https://docs.docker.com/get-docker/)
- [Aiken](https://aiken-lang.org/installation-instructions)
- [deno](https://docs.deno.com/runtime/manual/getting_started/installation)
- [jq](https://jqlang.github.io/jq/download/)
- [golang](https://golang.org/doc/install)
- [Rust & Cargo](https://www.rust-lang.org/tools/install)

### Check if you have all the prerequisites installed

```sh
cd caribic
cargo run check
```

### Run the services

```sh
cd caribic
cargo run start
```

### Stopping the services

```sh
cd caribic
cargo run stop
```

### Using the faucet to create and fund accounts in the test environment

Sidechain:
```sh
curl -X POST "http://localhost:4500/" -H  "accept: application/json" -H  "Content-Type: application/json" -d "{  \"address\": \"cosmos1ycel53a5d9xk89q3vdr7vm839t2vwl08pl6zk6\",  \"coins\": [    \"10token\",\"10stake\"  ]}"
```

or access to `http://localhost:4500`

To seed the Cardano addresses, you can use the caribic `config.json`. The file will be generated after running caribic for the first time. Per default, the config file is located at `<USER_HOME>/.caribic/config.json`.

### Sending tokens from Cosmos to Cardano and vice versa
```sh
docker exec -it relayer sh # Access to relayer container
cd /root && ./scripts/xtransfer.sh # Start submit transfer packet
```

After running `xtransfer.sh`, the relayer will capture the packet, relay a message to Cardano, call Ack on Cosmos, and by that complete the cycle.

```sh
2024-03-04T09:26:53.779140Z	info	Successful transaction	{"provider_type": "cardano", "chain_id": "cardano", "gas_used": 0, "height": 0, "msg_types": ["/ibc.core.channel.v1.MsgRecvPacket"], "tx_hash": "a35bc010a9e5e78c88469707aa10c3501bf19e51e0539b4720d70479d44fc3bc"}
...
2024-03-04T09:27:01.748158Z	info	Successful transaction	{"provider_type": "cosmos", "chain_id": "sidechain", "packet_src_channel": "channel-7", "packet_dst_channel": "channel-7", "gas_used": 55261, "fees": "", "fee_payer": "cosmos1ycel53a5d9xk89q3vdr7vm839t2vwl08pl6zk6", "height": 8573, "msg_types": ["/ibc.core.channel.v1.MsgAcknowledgement"], "tx_hash": "D162CC2356A09F09C80D616987FE4BE965FDEA7C3C93AC0F2D1D5BE4589C8A46"}
```

You can query balance using this endpoint:
#### Cosmos:
```sh
http://localhost:1317/cosmos/bank/v1beta1/balances/cosmos1ycel53a5d9xk89q3vdr7vm839t2vwl08pl6zk6
```
Notice that you will have voucher token with prefix: "ibc/"
Example:
```json
{
  "balances": [
    {
      "denom": "ibc/018463FA736C852FA78B23CE6CAE123B9182D18658E0F323B130BB4B1FBB6A52",
      "amount": "13578"
    }
  ]
}
```

#### Cardano:
```sh
http://localhost:1442/matches/addr_test1vqj82u9chf7uwf0flum7jatms9ytf4dpyk2cakkzl4zp0wqgsqnql?unspent&order=most_recent_first
```
Notice that you will have UTXO, asset with amount 2000:
Example:
```json
[
  {
    "transaction_index": 0,
    "transaction_id": "4ceee14cffdf8a03bba53e058bc02f0ed5e3cc1169d1e45963c02b780694b1af",
    "output_index": 2,
    "address": "addr_test1vqj82u9chf7uwf0flum7jatms9ytf4dpyk2cakkzl4zp0wqgsqnql",
    "value": {
      "coins": 1150770,
      "assets": {
        "901a270744d7eee7a2ef5e0199a29ca2636b3ede7e6fa520aba1a1c1.84916548b2860f827f717b20796c9ddd4742325677e9534cd5e92c8ca260c553": 2000
      }
    },
    "datum_hash": null,
    "script_hash": null,
    "created_at": {
      "slot_no": 4202,
      "header_hash": "3d2e1690468685cf5c95364b7200812f7252994d6a9620be0cc1f74991656020"
    },
    "spent_at": null
  }
]
```

### Register a new stake pool on the local Cardano blockchain
```sh
cd cardano/chains && ./regis-spo.sh <name>
```

Example:

```sh
cd cardano/chains && ./regis-spo.sh alice
```

### Retire a stake pool on the local Cardano blockchain
This will sent a tx to retire your pool in the next epoch:

```sh
cd cardano/chains && ./deregis-spo.sh <name>
```

Example:

```sh
cd cardano/chains && ./deregis-spo.sh alice
```

### Register a validator on Cosmos
This script will connect to your current docker and regis a new validator

```sh
Run this to check we only have 1 validator: curl -X GET "http://localhost:1317/cosmos/base/tendermint/v1beta1/validatorsets/latest" -H  "accept: application/json"

Run this to regis new validator: cd cosmos/scripts/ && ./regis-spo.sh

Run this to check we now have 2 validators: curl -X GET "http://localhost:1317/cosmos/base/tendermint/v1beta1/validatorsets/latest" -H  "accept: application/json"

```

### Unregister a validator on Cosmos
Stop the running script above, then wait for about 100 blocks (~2 mins), then check we only have 1 validator:

```sh
curl -X GET "http://localhost:1317/cosmos/base/tendermint/v1beta1/validatorsets/latest" -H  "accept: application/json"
```

### Test timeout packet
After successful create clients, connections, channels, terminate that terminal(A).

```sh
Access this url to check current balance in Cardano: http://localhost:1442/matches/addr_test1vqj82u9chf7uwf0flum7jatms9ytf4dpyk2cakkzl4zp0wqgsqnql?unspent&order=most_recent_first

Access this url to check current balance in Cosmos: http://localhost:1317/cosmos/bank/v1beta1/balances/cosmos1ycel53a5d9xk89q3vdr7vm839t2vwl08pl6zk6

```
Update script `/scripts/xtransfer.sh`, `timeout-time-offset` from `1h` to `10s`

Open another terminal(B) and run:
```sh
docker exec -it relayer sh
cd /root && ./scripts/xtransfer.sh
```

Recheck you current balance, notice that your balance will be deduct.
```sh
Access this url to check current balance in Cardano: http://localhost:1442/matches/addr_test1vqj82u9chf7uwf0flum7jatms9ytf4dpyk2cakkzl4zp0wqgsqnql?unspent&order=most_recent_first

Access this url to check current balance in Cosmos: http://localhost:1317/cosmos/bank/v1beta1/balances/cosmos1ycel53a5d9xk89q3vdr7vm839t2vwl08pl6zk6

```

In the terminal A, run this to execute timeout
```sh
cd /root && ./bin/rly start demo --processor legacy
```

After seeing something like `/ibc.core.channel.v1.MsgTimeout`, recheck you current balance, notice that your token will be return back.

### Test Crosschain Swap
**Prerequisite:**
  - `cargo run start` has been finished successfully.

1. Run:
  ```sh
  chains/osmosis/osmosis/scripts/setup_crosschain_swaps.sh
  ```
  to transfer Cardano mock token to Osmosis via IBC, create swap pool with this token and deploy contracts relates to crosschain swap. 

2. Copy address of `crosschain_swaps` contract in result of the previous command to variable `CROSSCHAIN_SWAPS_ADDRESS` of file `swap.sh`.

3. Run:
  ```sh
  swap.sh
  ```
  This command will send mock token in Cardano to Osmosis via IBC Packet Forward Middleware, swap this token to `uosmo` on created pool and send swapped token back to Cardano via Packet Forward Middleware again.

### Cardano developer ecosystem components used

The current implementation leverages a number of frameworks maintained by the Cardano developer community. We list them here for appreciation and transparency. Without tools like those listed and others, projects like this would not be possible:

- [Pallas](https://github.com/txpipe/pallas)
- [Lucid](https://github.com/spacebudz/lucid)
- [Ogmios](https://github.com/cardanosolutions/ogmios)
- [Kupo](https://github.com/cardanosolutions/kupo)
- [db-sync](https://github.com/IntersectMBO/cardano-db-sync)
- [gOuroboros](https://github.com/blinklabs-io/gouroboros)
- [Mithril](https://github.com/input-output-hk/mithril)

## :blue_heart: Contributing
All contributions are welcome! Please feel free to open a new thread on the issue tracker or submit a new pull request.

Please read [Contributing](CONTRIBUTING.md) in advance. Thank you for contributing!

## :books: Additional Documents
- [Code of Conduct](CODE_OF_CONDUCT.md)
- [Security](SECURITY.md)
