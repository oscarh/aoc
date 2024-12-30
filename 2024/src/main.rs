use clap::Parser;

#[derive(Debug, Parser)]
#[command(version, about, long_about = None)]
struct Args {
    #[command(subcommand)]
    command: Command,
}

mod util;

mod day01;
mod day02;
mod day03;
mod day04;
mod day05;
mod day06;
mod day07;
mod day08;
mod day09;
mod day10;
mod day11;

#[derive(Debug, Parser)]
enum Command {
    Day01,
    Day02,
    Day03,
    Day04,
    Day05,
    Day06,
    Day07,
    Day08,
    Day09,
    Day10,
    Day11,
}

fn main() {
    let args = Args::parse();
    match args.command {
        Command::Day01 => day01::run(),
        Command::Day02 => day02::run(),
        Command::Day03 => day03::run(),
        Command::Day04 => day04::run(),
        Command::Day05 => day05::run(),
        Command::Day06 => day06::run(),
        Command::Day07 => day07::run(),
        Command::Day08 => day08::run(),
        Command::Day09 => day09::run(),
        Command::Day10 => day10::run(),
        Command::Day11 => day11::run(),
    }
}
