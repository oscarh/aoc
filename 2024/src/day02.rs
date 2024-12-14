use std::fs::File;
use std::io::{self, BufRead};

fn level_is_safe(prev_level: i32, level: i32, was_increasing: Option<bool>) -> bool {
    let diff = level - prev_level;

    if !(1..=3).contains(&i32::abs(diff)) {
        return false;
    }

    let is_increasing = level > prev_level;

    if let Some(was_increasing) = was_increasing {
        if is_increasing != was_increasing {
            return false;
        }
    }

    true
}

fn report_is_safe(levels: &[i32], dampener: bool) -> bool {
    let mut prev = levels[0];
    let mut increasing = None;
    for (count, level) in levels[1..].iter().enumerate() {
        let idx = count + 1;
        if !level_is_safe(prev, *level, increasing) {
            if dampener {
                // try to remove the value before the previous one
                // this can caused the "direction" to become false
                if idx >= 2 {
                    let mut modified_report = Vec::new();
                    modified_report.extend_from_slice(&levels[0..idx - 2]);
                    modified_report.extend_from_slice(&levels[idx - 1..]);
                    if report_is_safe(&modified_report, false) {
                        return true;
                    }
                }

                // try remove the previous value and test this
                let mut modified_report = Vec::new();
                modified_report.extend_from_slice(&levels[0..idx - 1]);
                modified_report.extend_from_slice(&levels[idx..]);
                if report_is_safe(&modified_report, false) {
                    return true;
                }

                let mut modified_report = Vec::new();
                modified_report.extend_from_slice(&levels[0..idx]);
                modified_report.extend_from_slice(&levels[idx + 1..]);
                if report_is_safe(&modified_report, false) {
                    return true;
                }
            }
            return false;
        }
        increasing = Some(*level > prev);
        prev = *level;
    }

    true
}

fn is_safe(report: &str, dampner: bool) -> bool {
    let levels = report
        .split_whitespace()
        .map(|level| level.parse::<i32>().unwrap())
        .collect::<Vec<i32>>();
    report_is_safe(&levels[..], dampner)
}

fn part1() -> usize {
    let input = File::open("input/input_02.txt").unwrap();
    io::BufReader::new(input)
        .lines()
        .map(|report| report.unwrap())
        .filter(|report| is_safe(report, false))
        .count()
}

fn part2() -> usize {
    let input = File::open("input/input_02.txt").unwrap();
    io::BufReader::new(input)
        .lines()
        .map(|report| report.unwrap())
        .filter(|report| is_safe(report, true))
        .count()
}

pub fn run() {
    println!("Part1: {}", part1());
    println!("Part2: {}", part2());
}
