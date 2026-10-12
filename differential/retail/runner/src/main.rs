//! Runs one Rust decoder over the retail harness inputs and writes its
//! verdicts; see `differential/README.md`.

mod alloc;
mod contract;
mod valentine;
mod verdict;

use std::path::Path;
use std::process::ExitCode;

fn main() -> ExitCode {
    let args: Vec<String> = std::env::args().collect();
    let [_, decoder, out] = args.as_slice() else {
        eprintln!("usage: retail-runner <rust|valentine> <out-dir>");
        return ExitCode::from(2);
    };
    match run(decoder, Path::new(out)) {
        Ok(()) => ExitCode::SUCCESS,
        Err(error) => {
            eprintln!("retail-runner: {error}");
            ExitCode::FAILURE
        }
    }
}

fn run(decoder: &str, out: &Path) -> Result<(), String> {
    let decode = match decoder {
        "rust" => contract::run,
        "valentine" => valentine::run,
        other => return Err(format!("unknown decoder {other:?}")),
    };
    let mut verdicts = Vec::new();
    for input in verdict::read_inputs(out)? {
        decode(&input, &mut verdicts);
    }
    verdict::write(out, decoder, &verdicts)
}
