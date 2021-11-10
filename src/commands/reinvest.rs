use std::{convert::TryFrom, ops::Add, path, sync::Arc, time::Duration};

use abscissa_core::{Application, Command, Clap, Runnable};
use chrono::Utc;
use ethers::prelude::*;
use num_bigint::{BigInt, ToBigInt};
use num_traits::Zero;
use signatory::FsKeyStore;

use crate::{
    cellar_uniswap_wrapper::{UniswapV3CellarAddParams, UniswapV3CellarState, UniswapV3CellarTickInfo},
    erc20::Erc20State,
    gas::CellarGas,
    prelude::*,
    uniswap_pool::PoolState,
};

/// Cellars reinvest command
#[derive(Command, Debug, Clap)]
pub struct ReinvestCommand {
    #[clap(short = 'i', long)]
    pub cellar_id: u32,
}

impl Runnable for ReinvestCommand {
    fn run(&self) {
        let config = APP.config();
        let cellar = config.cellars.get(0).expect("Could not get cellar config");

        let keystore = path::Path::new(&config.keys.keystore);
        let keystore = FsKeyStore::create_or_open(keystore).expect("Could not open keystore");

        let name = &config
            .keys
            .rebalancer_key
            .parse()
            .expect("Could not parse name");
        let key = keystore.load(name).expect("Could not load key");

        let key = key
            .to_pem()
            .parse::<k256::elliptic_curve::SecretKey<k256::Secp256k1>>()
            .expect("Could not parse key");

        let wallet: LocalWallet = Wallet::from(key);

        let eth_host = config.ethereum.rpc.clone();
        let address = wallet.address();

        abscissa_tokio::run(&APP, async {
            let client = Provider::<Http>::try_from(eth_host)
                .unwrap()
                .interval(Duration::from_secs(3000u64));

            let client = SignerMiddleware::new(client, wallet);
            let gas = CellarGas::etherscan_standard().await.unwrap();

            // MyContract expects Arc, create with client
            let client = Arc::new(client);

            let mut contract_state = UniswapV3CellarState::new(cellar.cellar_address, client.clone());
            contract_state.gas_price = Some(gas);
   

            contract_state
                .reinvest()
                .await
                .unwrap();
        })
        .unwrap_or_else(|e| {
            status_err!("executor exited with error: {}", e);
            std::process::exit(1);
        });
    }
}