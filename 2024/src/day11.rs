use std::fmt::Debug;
use std::fs::File;

use std::collections::HashMap;

use crate::util;

struct StoneIter<'a> {
    next: Option<&'a Stone>,
}

impl<'a> Stones {
    fn iter(&'a self) -> StoneIter<'a> {
        StoneIter {
            next: self.head.as_deref(),
        }
    }
}

impl<'a> Iterator for StoneIter<'a> {
    type Item = &'a Stone;

    fn next(&mut self) -> Option<Self::Item> {
        self.next.inspect(|node| {
            self.next = node.next.as_deref();
        })
    }
}

struct Stone {
    value: u64,
    next: Option<Box<Stone>>,
}

struct Stones {
    head: Option<Box<Stone>>,
    len: usize,
}

fn even_digits(n: u64) -> bool {
    let no_digits = n.checked_ilog10().unwrap_or(0) + 1;
    no_digits % 2 == 0
}

fn split_value(n: u64) -> (u64, u64) {
    let str_val = n.to_string();
    let val_a = str_val[0..str_val.len() / 2].parse::<u64>().unwrap();
    let val_b = str_val[str_val.len() / 2..].parse::<u64>().unwrap();
    (val_a, val_b)
}

impl Stones {
    fn new() -> Self {
        Self { head: None, len: 0 }
    }

    fn push(&mut self, value: u64) {
        let node = Box::new(Stone {
            value,
            next: self.head.take(),
        });
        self.head = Some(node);
        self.len += 1;
    }

    fn blink(&mut self) {
        let mut curr = self.head.as_deref_mut();
        let mut skip_next = false;
        while let Some(node) = curr {
            if skip_next {
                curr = node.next.as_deref_mut();
                skip_next = false;
                continue;
            }

            if node.value == 0 {
                node.value = 1;
            } else if even_digits(node.value) {
                let (value_a, value_b) = split_value(node.value);
                let tail = Box::new(Stone {
                    value: value_b,
                    next: node.next.take(), // take the tail from the curent node
                });
                self.len += 1;

                node.value = value_a;
                // Update node to point to the new tail
                node.next = Some(tail);

                // don't know how move this to the next tail without this hack
                skip_next = true;
            } else {
                node.value *= 2024;
            }

            curr = node.next.as_deref_mut();
        }
    }

    fn len(&self) -> usize {
        self.len
    }
}

impl Debug for Stones {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "[")?;
        let mut first = true;
        for stone in self.iter() {
            if first {
                first = false;
            } else {
                write!(f, ", ")?;
            }
            write!(f, "{}", stone.value)?;
        }
        writeln!(f, "]")
    }
}

impl Drop for Stones {
    fn drop(&mut self) {
        let mut cur_node = self.head.take();
        while let Some(mut node) = cur_node {
            cur_node = node.next.take();
        }
    }
}

fn parse(mut lines: impl Iterator<Item = String>) -> Stones {
    let mut stones = Stones::new();
    lines
        .next()
        .unwrap()
        .split(' ')
        .rev()
        .map(|num| num.parse::<u64>().unwrap())
        .for_each(|num| stones.push(num));
    stones
}

fn part1(mut stones: Stones) -> usize {
    for _ in 0..25 {
        stones.blink();
    }
    stones.len()
}

#[derive(Debug)]
struct CompactStones {
    stones: HashMap<u64, u64>,
}

fn transform(stone: u64) -> impl Iterator<Item = u64> {
    let mut stones = Vec::new();

    if stone == 0 {
        stones.push(1);
    } else if even_digits(stone) {
        let (value_a, value_b) = split_value(stone);
        stones.push(value_a);
        stones.push(value_b);
    } else {
        stones.push(stone * 2024);
    }

    stones.into_iter()
}

impl CompactStones {
    fn new() -> Self {
        Self {
            stones: HashMap::new(),
        }
    }

    fn from(stones: Stones) -> Self {
        let mut cstones = Self::new();
        stones.iter().for_each(|stone| {
            cstones.inc(stone.value, 1);
        });
        cstones
    }

    fn blink(&mut self) {
        let stones: Vec<(u64, u64)> = self.stones.iter().map(|(k, v)| (*k, *v)).collect();
        self.stones = HashMap::new();

        for (stone, count) in stones {
            // these stones are removed
            transform(stone).for_each(|stone| self.inc(stone, count))
        }
    }

    fn inc(&mut self, value: u64, count: u64) {
        *self.stones.entry(value).or_default() += count
    }

    fn len(&self) -> u64 {
        self.stones.values().sum()
    }
}

fn part2(stones: Stones, times: usize) -> u64 {
    let mut stones = CompactStones::from(stones);

    for _ in 0..times {
        stones.blink();
    }

    stones.len()
}

pub fn run() {
    println!(
        "Part 1: {}",
        part1(parse(util::lines_iter(
            File::open("input/input_11.txt").unwrap()
        )))
    );
    println!(
        "Part 2: {}",
        part2(
            parse(util::lines_iter(File::open("input/input_11.txt").unwrap())),
            75
        )
    );
}

#[cfg(test)]
mod test {
    use super::*;

    const EXAMPLE_DATA: &'static [u8] = b"125 17";

    #[test]
    fn part1() {
        assert_eq!(super::part1(parse(util::lines_iter(EXAMPLE_DATA))), 55312);
    }

    #[test]
    fn part2_25_blinks() {
        assert_eq!(
            super::part2(parse(util::lines_iter(EXAMPLE_DATA)), 25),
            55312
        );
    }

    #[test]
    fn part2_6_blinks() {
        assert_eq!(super::part2(parse(util::lines_iter(EXAMPLE_DATA)), 6), 22);
    }
}
