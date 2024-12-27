use std::fs::File;

use crate::util;

static OPS: [fn(u64, u64) -> u64; 2] = [|a, b| -> u64 { a + b }, |a, b| -> u64 { a * b }];

fn validate(result: u64, acc: u64, values: &[u64]) -> bool {
    if values.is_empty() {
        return result == acc;
    }

    OPS.iter()
        .any(|op| validate(result, op(acc, values[0]), &values[1..]))
}

fn part1(lines: impl Iterator<Item = String>) -> u64 {
    let mut sum = 0;
    for line in lines {
        let mut parts = line.split(':');
        let result = parts.next().unwrap().parse::<u64>().unwrap();
        let values: Vec<u64> = parts
            .next()
            .unwrap()
            .trim()
            .split(' ')
            .map(|v| v.parse::<u64>().unwrap())
            .collect();
        if validate(result, values[0], &values[1..]) {
            sum += result;
        }
    }
    sum
}

/*
fn part2(lines: impl Iterator<Item = String>) -> usize {
    0
}
*/

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
    /*
    println!(
        "Part 2 (example data): {}",
        part2(util::lines_iter(example_data))
    );
    println!(
        "Part 2: {}", // 2151 is too high
        part2(util::lines_iter(File::open("input/input_07.txt").unwrap()))
    );
    */
}
