use std::fs::File;

use std::collections::HashSet;

use crate::util;

#[derive(Debug, Eq, PartialEq, Hash, Copy, Clone)]
struct Coordinate {
    x: usize,
    y: usize,
}

#[derive(Debug)]
struct TopoMap {
    heights: Vec<Vec<u8>>,
    trailheads: HashSet<Coordinate>,
}

fn parse(lines: impl Iterator<Item = String>) -> TopoMap {
    let mut map = TopoMap {
        heights: Vec::new(),
        trailheads: HashSet::new(),
    };

    for (y, line) in lines.enumerate() {
        let mut row = Vec::new();
        for (x, ch) in line.chars().enumerate() {
            let height = ch.to_digit(10).unwrap();
            if height == 0 {
                map.trailheads.insert(Coordinate { x, y });
            }
            row.push(height as u8);
        }
        map.heights.push(row);
    }

    map
}

fn neighbours(map: &TopoMap, coord: &Coordinate) -> impl Iterator<Item = Coordinate> {
    let mut vec = Vec::new();
    if coord.y > 0 {
        vec.push(Coordinate {
            x: coord.x,
            y: coord.y - 1,
        })
    }
    if coord.x > 0 {
        vec.push(Coordinate {
            x: coord.x - 1,
            y: coord.y,
        })
    }
    if coord.y < map.heights.len() - 1 {
        vec.push(Coordinate {
            x: coord.x,
            y: coord.y + 1,
        })
    }
    if coord.x < map.heights[0].len() - 1 {
        vec.push(Coordinate {
            x: coord.x + 1,
            y: coord.y,
        })
    }

    vec.into_iter()
}

impl TopoMap {
    fn is_climbable_from(&self, target: &Coordinate, from: &Coordinate) -> bool {
        let target_height = self.heights[target.y][target.x];
        let current_height = self.heights[from.y][from.x];
        target_height == current_height + 1
    }

    fn follow(&self, coord: &Coordinate, tops: &mut HashSet<Coordinate>) {
        if self.heights[coord.y][coord.x] == 9 {
            tops.insert(*coord);
        }

        neighbours(self, coord)
            .filter(|target| self.is_climbable_from(target, coord))
            .for_each(|coord| self.follow(&coord, tops));
    }

    fn start_at(&self, coord: &Coordinate) -> usize {
        let mut tops = HashSet::new();
        neighbours(self, coord)
            .filter(|target| self.is_climbable_from(target, coord))
            .for_each(|coord| self.follow(&coord, &mut tops));

        tops.len()
    }

    fn rate(&self, coord: &Coordinate) -> usize {
        if self.heights[coord.y][coord.x] == 9 {
            return 1;
        }

        neighbours(self, coord)
            .filter(|target| self.is_climbable_from(target, coord))
            .map(|coord| self.rate(&coord))
            .sum()
    }
}

fn part1(map: TopoMap) -> usize {
    map.trailheads.iter().map(|coord| map.start_at(coord)).sum()
}

fn part2(map: TopoMap) -> usize {
    map.trailheads.iter().map(|coord| map.rate(coord)).sum()
}

pub fn run() {
    println!(
        "Part 1: {}",
        part1(parse(util::lines_iter(
            File::open("input/input_10.txt").unwrap()
        )))
    );

    println!(
        "Part 2: {}",
        part2(parse(util::lines_iter(
            File::open("input/input_10.txt").unwrap()
        )))
    );
}

#[cfg(test)]
mod test {
    use super::*;

    const EXAMPLE_DATA: &'static [u8] = b"89010123\n\
                                          78121874\n\
                                          87430965\n\
                                          96549874\n\
                                          45678903\n\
                                          32019012\n\
                                          01329801\n\
                                          10456732";

    #[test]
    fn part1() {
        assert_eq!(super::part1(parse(util::lines_iter(EXAMPLE_DATA))), 36);
    }

    #[test]
    fn part2() {
        assert_eq!(super::part2(parse(util::lines_iter(EXAMPLE_DATA))), 81);
    }
}
