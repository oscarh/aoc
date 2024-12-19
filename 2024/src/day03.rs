use regex::{Match, Regex};
use std::fs::File;
use std::io::{self, BufRead};

fn to_i64(group: Option<Match<'_>>) -> i64 {
    group.unwrap().as_str().parse::<i64>().unwrap()
}

fn run_statements(re: Regex) -> impl Fn(i64, String) -> i64 {
    return move |sum, line| {
        re.captures_iter(&line)
            .map(|cap| {
                let x = to_i64(cap.get(1));
                let y = to_i64(cap.get(2));
                x * y
            })
            .fold(sum, |sum, value| sum + value)
    };
}

fn part1(lines: impl Iterator<Item = String>) -> i64
where
{
    let re = Regex::new(r"mul\(([0-9]+),([0-9]+)\)").unwrap();
    lines.fold(0, run_statements(re))
}

fn part2(lines: impl Iterator<Item = String>) -> i64 {
    let re =
        Regex::new(r"(?<enable>do\(\))|(?<disable>don't\(\))|mul\((?<x>[0-9]+),(?<y>[0-9]+)\)")
            .unwrap();
    let mut enabled = true;
    lines.fold(0, |sum, line| {
        re.captures_iter(&line)
            .map(|cap| {
                if cap.name("enable").is_some() {
                    enabled = true;
                    0
                } else if cap.name("disable").is_some() {
                    enabled = false;
                    0
                } else if enabled {
                    let x = to_i64(cap.name("x"));
                    let y = to_i64(cap.name("y"));
                    x * y
                } else {
                    0
                }
            })
            .fold(sum, |sum, value| sum + value)
    })
}

fn lines_iter<D>(input: D) -> impl Iterator<Item = String>
where
    D: std::io::Read + Sized,
{
    io::BufReader::new(input)
        .lines()
        .map(|lines_res| lines_res.unwrap())
}

pub fn run() {
    let mut part_1_example =
        &b"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"[..];
    println!(
        "Part 1 (example data): {}",
        part1(lines_iter(&mut part_1_example))
    );
    println!(
        "Part 1: {}",
        part1(lines_iter(File::open("input/input_03.txt").unwrap()))
    );
    let mut part_2_example =
        &b"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"[..];
    println!(
        "Part 2 (example data): {}",
        part2(lines_iter(&mut part_2_example))
    );
    println!(
        "Part 2: {}",
        part2(lines_iter(File::open("input/input_03.txt").unwrap()))
    );
}
