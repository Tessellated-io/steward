use crate::{application::APP, cellars, commands::cork_proposal::print_proposal, prelude::*};
use abscissa_core::{clap::Parser, Command, Runnable};
use steward_proto::steward::{
    cellar_v1_governance::{Function, InitiateShutdown},
    governance_call::Call,
    CellarV1Governance, GovernanceCall,
};

#[derive(Command, Debug, Parser)]
#[clap(
    long_about = "DESCRIPTION\n\nCalls initiateShutdown() on the target cellar contract at the specified block height.\nFor more information see https://github.com/PeggyJV/cellar-v1_5/blob/release/src/base/Cellar.sol"
)]
pub struct InitiateShutdownCmd {
    /// Target contract for scheduled cork.
    #[clap(short, long)]
    cellar_id: String,

    /// Block height to schedule cork.
    #[clap(short, long)]
    block_height: u64,

    /// Only print JSON output, omitting explanatory text
    #[clap(short, long)]
    quiet: bool,
}

impl Runnable for InitiateShutdownCmd {
    fn run(&self) {
        abscissa_tokio::run_with_actix(&APP, async {
            cellars::validate_cellar_id(&self.cellar_id)
                .await
                .unwrap_or_else(|err| {
                    status_err!("invalid cellar ID: {}", err);
                    std::process::exit(1);
                });

            let governance_call = GovernanceCall {
                call: Some(Call::CellarV1(CellarV1Governance {
                    function: Some(Function::InitiateShutdown(InitiateShutdown {})),
                })),
            };

            print_proposal(
                self.block_height,
                self.cellar_id.clone(),
                governance_call,
                self.quiet,
            )
        })
        .unwrap_or_else(|e| {
            status_err!("executor exited with error: {}", e);
            std::process::exit(1);
        });
    }
}
