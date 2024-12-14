use clap::Parser;

mod day01;
mod day02;

#[derive(Debug, Parser)]
enum Command {
    Day01,
    Day02,
}

#[derive(Debug, Parser)]
#[command(version, about, long_about = None)]
struct Args {
    #[command(subcommand)]
    command: Command,
}

fn main() {
    let args = Args::parse();
    match args.command {
        Command::Day01 => day01::run(),
        Command::Day02 => day02::run(),
    }
}
