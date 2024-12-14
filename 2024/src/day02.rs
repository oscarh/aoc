use std::fs::File;
use std::io::{self, BufRead};

fn is_safe(report: &str) -> bool {
    let mut levels = report.split_whitespace();
    let mut prev = levels.next().unwrap().parse::<i32>().unwrap();
    let mut was_increasing = None;
    for current in levels.map(|level| level.parse::<i32>().unwrap()) {
        let difference = current - prev;

        let is_increasing = current > prev;
        if let Some(was_increasing) = was_increasing {
            if was_increasing != is_increasing {
                return false;
            }
        }
        was_increasing = Some(is_increasing);
        let abs_difference = i32::abs(difference);
        if !(1..=3).contains(&abs_difference) {
            return false;
        }

        prev = current;
    }

    true
}

fn part1() -> usize {
    let input = File::open("input/input_02.txt").unwrap();
    io::BufReader::new(input)
        .lines()
        .map(|report| report.unwrap())
        .filter(|report| is_safe(report))
        .count()
}

pub fn run() {
    println!("Part1: {}", part1());
}
