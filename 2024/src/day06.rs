use std::collections::HashSet;
use std::fs::File;

use crate::util;

#[derive(Copy, Clone, Default)]
enum Direction {
    #[default]
    Up,
    Left,
    Right,
    Down,
}

impl From<Direction> for char {
    fn from(value: Direction) -> char {
        match value {
            Direction::Up => '^',
            Direction::Right => '>',
            Direction::Down => 'v',
            Direction::Left => '<',
        }
    }
}

impl From<char> for Direction {
    fn from(value: char) -> Self {
        match value {
            '^' => Self::Up,
            '>' => Self::Right,
            '<' => Self::Left,
            'v' => Self::Down,
            _ => panic!("Not a direction"),
        }
    }
}

#[derive(Debug, Clone, Eq, Hash, PartialEq, Default)]
struct Pos {
    x: usize,
    y: usize,
}

#[derive(Default)]
struct Guard {
    pos: Pos,
    dir: Direction,
}

struct Map {
    rows: usize,
    columns: usize,
    obstructions: HashSet<Pos>,
    guard: Guard,
    visited: HashSet<Pos>,
}

impl std::fmt::Debug for Map {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        for y in 0..self.rows {
            for x in 0..self.columns {
                let pos = Pos { x, y };
                let ch = if self.obstructions.contains(&pos) {
                    '#'
                } else if self.guard.pos == pos {
                    self.guard.dir.into()
                } else if self.visited.contains(&pos) {
                    'X'
                } else {
                    '.'
                };

                write!(f, "{}", ch)?;
            }
            writeln!(f)?;
        }
        Ok(())
    }
}

fn parse_map(lines: impl Iterator<Item = String>) -> Map {
    let mut obstructions: HashSet<Pos> = HashSet::new();
    let mut guard = Guard::default();
    let mut rows = 0;
    let mut columns = 0;
    for (y, line) in lines.enumerate() {
        rows += 1;
        columns = line.len();
        for (x, ch) in line.chars().enumerate() {
            match ch {
                '#' => {
                    let _ = obstructions.insert(Pos { x, y });
                }
                '.' => (),
                dir => {
                    guard.pos = Pos { x, y };
                    guard.dir = Direction::from(dir);
                }
            }
        }
    }

    Map {
        rows,
        columns,
        obstructions,
        guard,
        visited: HashSet::new(),
    }
}

fn next_pos(map: &Map) -> Option<Pos> {
    let pos = &map.guard.pos;
    match map.guard.dir {
        Direction::Up => {
            if map.guard.pos.y > 0 {
                Some(Pos {
                    x: pos.x,
                    y: pos.y - 1,
                })
            } else {
                None
            }
        }
        Direction::Right => {
            if map.guard.pos.x < map.columns - 1 {
                Some(Pos {
                    x: pos.x + 1,
                    y: pos.y,
                })
            } else {
                None
            }
        }
        Direction::Down => {
            if map.guard.pos.y < map.rows - 1 {
                Some(Pos {
                    x: pos.x,
                    y: pos.y + 1,
                })
            } else {
                None
            }
        }
        Direction::Left => {
            if map.guard.pos.x > 0 {
                Some(Pos {
                    x: pos.x - 1,
                    y: pos.y,
                })
            } else {
                None
            }
        }
    }
}

fn next_dir(dir: &Direction) -> Direction {
    match dir {
        Direction::Up => Direction::Right,
        Direction::Right => Direction::Down,
        Direction::Down => Direction::Left,
        Direction::Left => Direction::Up,
    }
}

fn move_util_outside(map: &mut Map) -> usize {
    map.visited.insert(map.guard.pos.clone());
    while let Some(next) = next_pos(map) {
        if map.obstructions.contains(&next) {
            map.guard.dir = next_dir(&map.guard.dir);
        } else {
            map.visited.insert(next.clone());
            map.guard.pos = next;
        }
    }

    // 5079 == too low
    map.visited.len()
}

fn part1(lines: impl Iterator<Item = String>) -> usize {
    let mut map = parse_map(lines);
    move_util_outside(&mut map)
}

pub fn run() {
    let example_data = &b"....#.....\n\
                          .........#\n\
                          ..........\n\
                          ..#.......\n\
                          .......#..\n\
                          ..........\n\
                          .#..^.....\n\
                          ........#.\n\
                          #.........\n\
                          ......#..."[..];
    println!(
        "Part 1 (example data): {}",
        part1(util::lines_iter(example_data))
    );
    println!(
        "Part 1: {}",
        part1(util::lines_iter(File::open("input/input_06.txt").unwrap()))
    );
    /*
    println!(
        "Part 2 (example data): {}",
        part2(util::lines_iter(example_data))
    );
    println!(
        "Part 2: {}",
        part2(util::lines_iter(File::open("input/input_06.txt").unwrap()))
    );
    */
}
