#![cfg_attr(feature = "ci_run", deny(warnings))]

extern crate clap;
extern crate ton_node;
#[macro_use]
extern crate log;
extern crate ton_block;
extern crate ton_types;
extern crate ton_vm as tvm;
extern crate ed25519_dalek;
extern crate log4rs;
extern crate reqwest;
extern crate http;
extern crate parking_lot;
extern crate iron;
extern crate serde;
#[macro_use]
extern crate serde_json;
extern crate serde_derive;
extern crate router;
extern crate base64;
extern crate adnl;
extern crate ton_block_json;

mod types;

use clap::{App, Arg};
use std::env;
use std::path::{Path, PathBuf};
use std::sync::Arc;
use std::thread;
use std::time::Duration;
use ton_node::error::NodeResult;
use adnl::config::AdnlServerConfig;
use ton_node::node_engine::{DocumentsDb, MessagesReceiver};
use ton_node::node_engine::ton_node_engine::TonNodeEngine;
use ton_node::node_engine::ton_node_handlers::init_ton_node_handlers;
use ed25519_dalek::{Keypair};
use ton_node::node_engine::config::NodeConfig;
use std::fs;
use types::{ArangoHelper, KafkaProxyMsgReceiver};

fn main() {
    run().expect("Error run node");
}

fn run() -> Result<(), ()> {

    println!("TON Startup Edition Prototype {}\nCOMMIT_ID: {}\nBUILD_DATE: {}\nCOMMIT_DATE: {}\nGIT_BRANCH: {}",
            env!("CARGO_PKG_VERSION"),
            env!("BUILD_GIT_COMMIT"),
            env!("BUILD_TIME") ,
            env!("BUILD_GIT_DATE"),
            env!("BUILD_GIT_BRANCH"));

    let args = App::new(env!("CARGO_PKG_NAME"))
        .version(env!("CARGO_PKG_VERSION"))
        .arg(
            Arg::with_name("workdir")
                .help("Path to working directory")
                .long("workdir")
                .takes_value(true)
                .max_values(1),
        )
        .arg(
            Arg::with_name("config")
                .help("configuration file name")
                .long("config")
                .required(true)
                .takes_value(true)
                .max_values(1),
        )
        // .arg(
        //     Arg::with_name("localhost")
        //         .help("Localhost connectivity only")
        //         .long("localhost"),
        // )
        .arg(
            Arg::with_name("automsg")
                .help("Auto generate message timeout")
                .long("automsg")
                .takes_value(true)
                .max_values(1),
        )
        .get_matches();

    if args.is_present("workdir") {
        if let Some(workdir) = args.value_of("workdir") {
            env::set_current_dir(Path::new(workdir)).unwrap()
        }
    }

    log4rs::init_file("./log_cfg.yml", Default::default()).expect("Error initialize logging configuration. config: log_cfg.yml");

    info!(target: "node", "TON Node Startup Edition Prototype {}\nCOMMIT_ID: {}\nBUILD_DATE: {}\nCOMMIT_DATE: {}\nGIT_BRANCH: {}",
        env!("CARGO_PKG_VERSION"),
        env!("BUILD_GIT_COMMIT"),
        env!("BUILD_TIME") ,
        env!("BUILD_GIT_DATE"),
        env!("BUILD_GIT_BRANCH"));

    let err = start_node(args.value_of("config").unwrap_or_default());
    log::error!(target: "node", "{:?}", err);

    Ok(())
}

fn start_node(config: &str) -> NodeResult<()> {

    let json = fs::read_to_string(Path::new(config))?;

    let (config, public_keys) = get_config_params(&json);

    let keypair = fs::read(Path::new(&config.private_key))
        .expect(&format!("Error reading key file {}", config.private_key));
    let private_key = Keypair::from_bytes(&keypair).unwrap();

    let adnl_config = AdnlServerConfig::from_json_config(&config.adnl);

    let db: Box<dyn DocumentsDb> = Box::new(ArangoHelper::from_config(&config.document_db_config())?);
    let receivers: Vec<Box<dyn MessagesReceiver>> = vec!(
            Box::new(KafkaProxyMsgReceiver::from_config(&config.kafka_msg_recv_config())?));

    let ton = TonNodeEngine::with_params( 
        config.shard_id_config().shard_ident(),
        false,
        config.port,
        config.node_index, 
        config.poa_validators,
        config.poa_interval,
        private_key,
        public_keys,
        config.boot,
        adnl_config,
        receivers,
        Some(db),
        PathBuf::from("./"),
    )?;

    init_ton_node_handlers(&ton);
    let ton = Arc::new(ton);
    TonNodeEngine::start(ton.clone())?;

    loop {
        thread::sleep(Duration::from_secs(1));
    }
}

pub fn get_config_params(json: &str) -> (NodeConfig, Vec<ed25519_dalek::PublicKey>) {
    match NodeConfig::parse(json) {
        Ok(config) => match config.import_keys() {
            Ok(keys) => (config, keys),
            Err(err) => {
                log::error!(target: "node", "{}", err);
                panic!("{}", err)
            }
        },
        Err(err) => {
            log::error!(target: "node", "Error parsing configuration file. {}", err);
            panic!("Error parsing configuration file. {}", err)
        },
    }
}
