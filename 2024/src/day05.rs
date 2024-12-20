use core::cmp::Ordering;
use std::collections::{HashMap, HashSet};
use std::fs::File;

use crate::util;

type Rules = HashMap<u64, HashSet<u64>>;

fn parse_rule(rules: &mut Rules, rule: &str) {
    let rule = rule.split('|').collect::<Vec<&str>>();
    let key = rule[0].parse::<u64>().unwrap();
    let before = rule[1].parse::<u64>().unwrap();
    rules.entry(key).or_default().insert(before);
}

fn parse_update(update: &str) -> Vec<u64> {
    let mut vec = Vec::new();
    for element in update.split(',') {
        vec.push(element.parse::<u64>().unwrap());
    }
    vec
}

fn order(rules: &Rules) -> impl FnMut(&u64, &u64) -> Ordering + use<'_> {
    |a, b| {
        if rules.contains_key(a) && rules[a].contains(b) {
            Ordering::Less
        } else if rules.contains_key(b) && rules[b].contains(a) {
            Ordering::Greater
        } else {
            Ordering::Equal
        }
    }
}

fn ordered(rules: &Rules) -> impl FnMut(&u64, &u64) -> bool + use<'_> {
    let mut order_fn = order(rules);
    move |a, b| {
        order_fn(a, b) == Ordering::Less
    }
}

fn get_middle_value<T>(slice: &[T]) -> &T {
    let middle_pos = (slice.len() - 1) / 2;
    &slice[middle_pos]
}

fn part1(lines: impl Iterator<Item = String>) -> u64 {
    let mut sort_rules = HashMap::new();
    let mut rules_finished = false;
    let mut count = 0;
    for line in lines {
        if line.is_empty() {
            rules_finished = true;
            continue;
        }

        if !rules_finished {
            parse_rule(&mut sort_rules, &line)
        } else {
            let update = parse_update(&line);
            if update.is_sorted_by(ordered(&sort_rules)) {
                count += get_middle_value(&update);
            }
        }
    }
    count
}

fn part2(lines: impl Iterator<Item = String>) -> u64 {
    let mut sort_rules = HashMap::new();
    let mut rules_finished = false;
    let mut count = 0;
    for line in lines {
        if line.is_empty() {
            rules_finished = true;
            continue;
        }

        if !rules_finished {
            parse_rule(&mut sort_rules, &line)
        } else {
            let mut update = parse_update(&line);
            if !update.is_sorted_by(ordered(&sort_rules)) {
                update.sort_by(order(&sort_rules));
                count += get_middle_value(&update);
            }
        }
    }
    count
}

pub fn run() {
    let example_data = &b"47|53\n\
           97|13\n\
           97|61\n\
           97|47\n\
           75|29\n\
           61|13\n\
           75|53\n\
           29|13\n\
           97|29\n\
           53|29\n\
           61|53\n\
           97|53\n\
           61|29\n\
           47|13\n\
           75|47\n\
           97|75\n\
           47|61\n\
           75|61\n\
           47|29\n\
           75|13\n\
           53|13\n\
           \n\
           75,47,61,53,29\n\
           97,61,53,29,13\n\
           75,29,13\n\
           75,97,47,61,53\n\
           61,13,29\n\
           97,13,75,29,47"[..];
    println!(
        "Part 1 (example data): {}",
        part1(util::lines_iter(example_data))
    );
    println!(
        "Part 1: {}",
        part1(util::lines_iter(File::open("input/input_05.txt").unwrap()))
    );
    println!(
        "Part 2 (example data): {}",
        part2(util::lines_iter(example_data))
    );
    println!(
        "Part 2: {}",
        part2(util::lines_iter(File::open("input/input_05.txt").unwrap()))
    );
}
