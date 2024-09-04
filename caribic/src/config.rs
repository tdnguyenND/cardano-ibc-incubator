use crate::logger::{error, get_verbosity, log, verbose, Verbosity};
use crate::utils::{download_file, unzip_file, IndicatorMessage};
use console::style;
use dirs::home_dir;
use fs_extra::dir::create_all;
use lazy_static::lazy_static;
use serde::{Deserialize, Serialize};
use std::fs;
use std::io::{stdin, stdout, Write};
use std::path::Path;
use std::process;
use std::sync::Mutex;

#[derive(Debug, Serialize, Deserialize, Clone)]
pub struct Config {
    pub project_root: String,
    pub use_mithril: bool,
    pub local_osmosis: bool,
    pub cardano: Cardano,
}

#[derive(Debug, Serialize, Deserialize, Clone)]
pub struct Cardano {
    pub services: Services,
    pub bootstrap_addresses: Vec<BootstrapAddress>,
}

#[derive(Debug, Serialize, Deserialize, Clone)]
pub struct BootstrapAddress {
    pub address: String,
    pub amount: i64,
}

#[derive(Debug, Serialize, Deserialize, Clone)]
pub struct Services {
    pub db_sync: bool,
    pub kupo: bool,
    pub ogmios: bool,
    pub cardano_node: bool,
    pub postgres: bool,
}

pub async fn create_config_file(config_path: &str) -> Config {
    let mut default_config = Config::default();

    if get_verbosity() == Verbosity::Verbose
        || get_verbosity() == Verbosity::Info
        || get_verbosity() == Verbosity::Standard
    {
        println!("Config file not found at: {}", config_path);
        let mut input = String::new();
        log(&format!(
            "Do you want to create it now? ({}es/no): ",
            style("y").bold().underlined()
        ));
        stdout().flush().unwrap();
        stdin().read_line(&mut input).unwrap();

        if let Some(home_path) = home_dir() {
            if input.trim().eq_ignore_ascii_case("yes")
                || input.trim().eq_ignore_ascii_case("y")
                || input.trim().is_empty()
            {
                let default_project_root = format!(
                    "{}/.caribic/cardano-ibc-incubator",
                    home_path.as_path().display()
                );
                log(&format!(
                    "Enter the project root path for 'cardano-ibc-incubator' (default: {}):",
                    default_project_root
                ));

                let mut project_root = String::new();
                stdin().read_line(&mut project_root).unwrap();
                let mut project_root = if project_root.trim().is_empty() {
                    default_project_root
                } else {
                    project_root.trim().to_string()
                };

                if project_root.starts_with("~") {
                    project_root = project_root.replace("~", home_path.to_str().unwrap());
                }
                let project_root_path = Path::new(&project_root);
                let parent_dir = project_root_path.parent().unwrap();

                if !project_root_path.exists() {
                    verbose(&format!(
                        "cardano-ibc-incubator folder does not exist. It will be downloaded to: {}",
                        project_root_path.display(),
                    ));
                    fs::create_dir_all(parent_dir).expect("Failed to create project root folder.");
                    let github_url = "https://github.com/cardano-foundation/cardano-ibc-incubator/archive/refs/heads/main.zip";
                    download_file(
                        github_url,
                        &parent_dir.join("cardano-ibc-incubator-main.zip"),
                        Some(IndicatorMessage {
                            message: "Downloading cardano-ibc-incubator project".to_string(),
                            step: "Step 1/2".to_string(),
                            emoji: "📥 ".to_string(),
                        }),
                    )
                    .await
                    .expect("Failed to download cardano-ibc-incubator project");

                    log(&format!(
                        "{} 📦 Extracting cardano-ibc-incubator project...",
                        style("Step 2/2").bold().dim()
                    ));

                    unzip_file(
                        parent_dir.join("cardano-ibc-incubator-main.zip").as_path(),
                        project_root_path,
                    )
                    .expect("Failed to unzip cardano-ibc-incubator project");
                    fs::remove_file(parent_dir.join("cardano-ibc-incubator-main.zip"))
                        .expect("Failed to cleanup cardano-ibc-incubator-main.zip");
                }

                default_config.project_root = project_root;
                verbose(&format!(
                    "Project root path set to: {}",
                    default_config.project_root
                ));
            } else {
                error("Config file not found. Exiting.");
                process::exit(0);
            }
        } else {
            error("Failed to resolve home directory. Exiting.");
            process::exit(0);
        }
    } else {
        error("No config file has been found. Creating a new config does not work with log levels warning, error or quite.");
        process::exit(0);
    }

    verbose(&format!("caribic config file: {:#?}", default_config));

    default_config
}

impl Config {
    fn default() -> Self {
        Config {
            project_root: "~/.caribic".to_string(),
            use_mithril: false,
            local_osmosis: true,
            cardano: Cardano {
                services: Services {
                    db_sync: true,
                    kupo: true,
                    ogmios: true,
                    cardano_node: true,
                    postgres: true,
                },
                bootstrap_addresses: vec![                    
                    BootstrapAddress {
                        address: "addr_test1qrwuz99eywdpm9puylccmvqfu6lue968rtt36nzeal7czuu4wq3n84h8ntp3ta30kyxx8r0x2u4tgr5a8y9hp5vjpngsmwy0wg".to_string(),
                        amount: 30000000000,
                    },
                    BootstrapAddress {
                        address: "addr_test1vz8nzrmel9mmmu97lm06uvm55cj7vny6dxjqc0y0efs8mtqsd8r5m".to_string(),
                        amount: 30000000000,
                    },
                    BootstrapAddress {
                        address: "addr_test1vqj82u9chf7uwf0flum7jatms9ytf4dpyk2cakkzl4zp0wqgsqnql".to_string(),
                        amount: 30000000000,
                    }
                ],
            },
        }
    }

    async fn load_from_file(config_path: &str) -> Self {
        if Path::new(config_path).exists() {
            let file_content =
                fs::read_to_string(config_path).expect("Failed to read config file.");
            serde_json::from_str(&file_content).unwrap_or_else(|_| {
                eprintln!("Failed to parse config file, using default config.");
                Config::default()
            })
        } else {
            verbose("Config file not found, creating default config.");
            let default_config = create_config_file(config_path).await;
            let parent_dir = Path::new(config_path).parent().unwrap();
            create_all(parent_dir, false).expect("Failed to create config dir.");
            let json_content = serde_json::to_string_pretty(&default_config)
                .expect("Failed to serialize default config.");
            fs::write(Path::new(config_path), json_content)
                .expect("Failed to write default config file.");
            default_config
        }
    }
}

lazy_static! {
    static ref CONFIG: Mutex<Config> = Mutex::new(Config::default());
}

pub async fn init(config_path: &str) {
    let mut config = CONFIG.lock().unwrap();
    *config = Config::load_from_file(config_path).await;
}

pub fn get_config() -> Config {
    CONFIG.lock().unwrap().clone()
}
