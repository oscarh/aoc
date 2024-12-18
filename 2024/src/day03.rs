use regex::Regex;
use std::fs::File;
use std::io::{self, BufRead};

fn run_statements(re: Regex) -> impl Fn(i64, std::io::Result<String>) -> i64 {
    return move |sum, line| {
        re.captures_iter(&line.unwrap())
            .map(|cap| {
                let x = cap.get(1).unwrap().as_str().parse::<i64>().unwrap();
                let y = cap.get(2).unwrap().as_str().parse::<i64>().unwrap();
                x * y
            })
            .fold(sum, |sum, value| sum + value)
    };
}

fn part1<D>(input: io::BufReader<D>) -> i64
where
    D: std::io::Read + Sized,
{
    let re = Regex::new("mul\\(([0-9]+),([0-9]+)\\)").unwrap();
    input.lines().fold(0, run_statements(re))
}

fn part2<D>(_input: io::BufReader<D>) -> usize {
    0
}

pub fn run() {
    let mut example =
        &b"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"[..];
    println!(
        "Part 1 (example data): {}",
        part1(io::BufReader::new(&mut example))
    );
    println!(
        "Part 1: {}",
        part1(io::BufReader::new(
            File::open("input/input_03.txt").unwrap()
        ))
    );
    println!(
        "Part 2: {}",
        part2(io::BufReader::new(
            File::open("input/input_03.txt").unwrap()
        ))
    );
}
