use std::fs::File;

use crate::util;

fn add(a: u64, b: u64) -> u64 {
    a + b
}

fn mul(a: u64, b: u64) -> u64 {
    a * b
}

fn concat(a: u64, b: u64) -> u64 {
    format!("{a}{b}").parse::<u64>().unwrap()
}

type Op = fn(u64, u64) -> u64;
static OPS: [Op; 3] = [add, mul, concat];

fn calc(ops: &[Op], result: u64, acc: u64, digits: &[u64]) -> bool {
    if acc > result {
        return false;
    }

    if digits.is_empty() {
        return result == acc;
    }

    if digits.is_empty() {
        return result == acc;
    }

    ops.iter()
        .any(|op| calc(ops, result, op(acc, digits[0]), &digits[1..]))
}

type Equation = (u64, Vec<u64>);

fn parse(lines: impl Iterator<Item = String>) -> impl Iterator<Item = Equation> {
    lines.map(|line| {
        let mut parts = line.split(':');
        let result = parts.next().unwrap().parse::<u64>().unwrap();
        let digits: Vec<u64> = parts
            .next()
            .unwrap()
            .trim()
            .split(' ')
            .map(|v| v.parse::<u64>().unwrap())
            .collect();
        (result, digits)
    })
}

fn part1(lines: impl Iterator<Item = String>) -> u64 {
    parse(lines).fold(0, |acc, (result, digits)| {
        if calc(&OPS[0..=1], result, digits[0], &digits[1..]) {
            acc + result
        } else {
            acc
        }
    })
}

fn part2(lines: impl Iterator<Item = String>) -> u64 {
    parse(lines).fold(0, |acc, (result, digits)| {
        if calc(&OPS, result, digits[0], &digits[1..]) {
            acc + result
        } else {
            acc
        }
    })
}

pub fn run() {
    let example_data = &b"190: 10 19\n\
                          3267: 81 40 27\n\
                          83: 17 5\n\
                          156: 15 6\n\
                          7290: 6 8 6 15\n\
                          161011: 16 10 13\n\
                          192: 17 8 14\n\
                          21037: 9 7 18 13\n\
                          292: 11 6 16 20"[..];
    println!(
        "Part 1 (example data): {}",
        part1(util::lines_iter(example_data))
    );
    println!(
        "Part 1: {}",
        part1(util::lines_iter(File::open("input/input_07.txt").unwrap()))
    );
    println!(
        "Part 2 (example data): {}",
        part2(util::lines_iter(example_data))
    );
    println!(
        "Part 2: {}", // 2151 is too high
        part2(util::lines_iter(File::open("input/input_07.txt").unwrap()))
    );
}
