use sorted_vec::SortedVec;
use std::collections::HashMap;
use std::fs::File;
use std::io::{self, BufRead};

fn part1() {
    let input = File::open("input/input_01.txt").unwrap();
    let mut group_1 = SortedVec::new();
    let mut group_2 = SortedVec::new();
    for line in io::BufReader::new(input).lines() {
        let line = line.unwrap();
        let vec = line
            .split_whitespace()
            .map(|id| id.parse::<i32>().unwrap())
            .collect::<Vec<i32>>();
        group_1.push(vec[0]);
        group_2.push(vec[1]);
    }
    let mut sum = 0;
    for (g1, g2) in group_1.iter().zip(group_2.iter()) {
        sum += i32::abs(g1 - g2);
    }

    println!("Part 1: {}", sum);
}

fn part2() {
    let input = File::open("input/input_01.txt").unwrap();
    let mut group_1 = Vec::new();
    let mut group_2_count = HashMap::new();
    for line in io::BufReader::new(input).lines() {
        let line = line.unwrap();
        let vec = line
            .split_whitespace()
            .map(|id| id.parse::<i32>().unwrap())
            .collect::<Vec<i32>>();
        group_1.push(vec[0]);
        *group_2_count.entry(vec[1]).or_insert(0) += 1;
    }

    let mut sum = 0;
    for id in group_1 {
        sum += id * *group_2_count.entry(id).or_default();
    }

    println!("Part 2: {}", sum);
}

pub fn run() {
    part1();
    part2();
}
