use color_print::cwrite;
use std::collections::{HashMap, HashSet};
use std::fs::File;

use std::fmt::Display;

use crate::util;

#[derive(Debug, Eq, Hash, PartialEq, Copy, Clone, Ord, PartialOrd)]
struct Coordinate {
    x: i64,
    y: i64,
}

impl Coordinate {
    fn new(x: i64, y: i64) -> Self {
        Self { x, y }
    }
}

struct Map {
    rows: i64,
    columns: i64,
    positions: HashMap<Coordinate, char>,
    frequencies: HashMap<char, HashSet<Coordinate>>,
    antinodes: HashSet<Coordinate>,
}

struct MapPoint {
    coordinate: Coordinate,
    frequency: char,
}

struct MapIterator {
    columns: i64,
    rows: i64,
    x: i64,
    y: i64,
    positions: HashMap<Coordinate, char>,
}

impl Iterator for MapIterator {
    type Item = MapPoint;

    fn next(&mut self) -> Option<Self::Item> {
        if self.y == self.rows {
            return None;
        }

        let coord = Coordinate {
            x: self.x,
            y: self.y,
        };
        let point = MapPoint {
            coordinate: coord,
            frequency: *self.positions.get(&coord).unwrap_or(&'.'),
        };

        if self.x < self.columns - 1 {
            self.x += 1;
        } else {
            self.y += 1;
            self.x = 0;
        }
        Some(point)
    }
}

impl Map {
    fn new() -> Self {
        Self {
            rows: 0,
            columns: 0,
            positions: HashMap::new(),
            frequencies: HashMap::new(),
            antinodes: HashSet::new(),
        }
    }

    fn iter(&self) -> MapIterator {
        MapIterator {
            columns: self.columns,
            rows: self.rows,
            x: 0,
            y: 0,
            positions: self.positions.clone(),
        }
    }
}

impl Display for Map {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let mut row = 0;
        for point in self.iter() {
            if point.coordinate.y != row {
                writeln!(f)?;
                row = point.coordinate.y;
            }
            if self.antinodes.contains(&point.coordinate) {
                cwrite!(f, "<red>{}</red>", point.frequency)?
            } else {
                write!(f, "{}", point.frequency)?
            }
        }
        Ok(())
    }
}

fn calc_delta(a: &Coordinate, b: &Coordinate) -> (i64, i64) {
    (a.x - b.x, a.y - b.y)
}

fn add_delta(c: &Coordinate, delta_x: i64, delta_y: i64) -> Coordinate {
    Coordinate {
        x: c.x + delta_x,
        y: c.y + delta_y,
    }
}

fn within_map(c: &Coordinate, map: &Map) -> bool {
    c.x >= 0 && c.y >= 0 && c.x < map.columns && c.y < map.rows
}

fn add_antinodes(map: &mut Map, resonant_harmonics: bool) {
    for coordinates in map.frequencies.values().clone() {
        for a in coordinates {
            for b in coordinates {
                if a == b {
                    continue;
                }

                map.antinodes.insert(*a);
                let (delta_x, delta_y) = calc_delta(a, b);
                let mut from = *a;
                loop {
                    let node = add_delta(&from, delta_x, delta_y);
                    if within_map(&node, map) {
                        map.antinodes.insert(node);
                    } else {
                        break;
                    }
                    if !resonant_harmonics {
                        break;
                    }
                    from = node;
                }
            }
        }
    }
}

fn parse(lines: impl Iterator<Item = String>) -> Map {
    let mut map = Map::new();
    for (row, line) in lines.enumerate() {
        map.rows = row as i64 + 1;
        for (column, ch) in line.chars().enumerate() {
            map.columns = column as i64 + 1;
            if ch != '.' {
                let coord = Coordinate::new(column as i64, row as i64);
                map.positions.insert(coord, ch);
                map.frequencies.entry(ch).or_default().insert(coord);
            }
        }
    }
    map
}

fn part1(lines: impl Iterator<Item = String>) -> usize {
    let mut map = parse(lines);
    add_antinodes(&mut map, false);
    map.antinodes.len()
}

fn part2(lines: impl Iterator<Item = String>) -> usize {
    let mut map = parse(lines);
    add_antinodes(&mut map, true);
    map.antinodes.len()
}

pub fn run() {
    let example_data = &b"............\n\
                          ........0...\n\
                          .....0......\n\
                          .......0....\n\
                          ....0.......\n\
                          ......A.....\n\
                          ............\n\
                          ............\n\
                          ........A...\n\
                          .........A..\n\
                          ............\n\
                          ............"[..];
    println!(
        "Part 1 (example data): {}",
        part1(util::lines_iter(example_data))
    );
    println!(
        "Part 1: {}",
        part1(util::lines_iter(File::open("input/input_08.txt").unwrap()))
    );
    println!(
        "Part 2 (example data): {}",
        part2(util::lines_iter(example_data))
    );
    println!(
        "Part 2: {}",
        part2(util::lines_iter(File::open("input/input_08.txt").unwrap()))
    );
}
