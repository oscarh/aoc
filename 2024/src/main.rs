use clap::Parser;

#[derive(Debug, Parser)]
#[command(version, about, long_about = None)]
struct Args {
    #[command(subcommand)]
    command: Command,
}

mod day01;
mod day02;
mod day03;

#[derive(Debug, Parser)]
enum Command {
    Day01,
    Day02,
    Day03,
}

fn main() {
    let args = Args::parse();
    match args.command {
        Command::Day01 => day01::run(),
        Command::Day02 => day02::run(),
        Command::Day03 => day03::run(),
    }
}
