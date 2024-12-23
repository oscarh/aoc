use std::collections::HashSet;
use std::fs::File;

use crate::util;

#[derive(Copy, Clone, Default, Debug, Eq, PartialEq, Hash)]
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

#[derive(Debug, Copy, Clone, Eq, Hash, PartialEq, Default)]
struct Pos {
    x: i32,
    y: i32,
}

impl std::ops::Add<(i8, i8)> for Pos {
    type Output = Self;

    fn add(self, (delta_x, delta_y): (i8, i8)) -> Self::Output {
        Self {
            x: self.x + delta_x as i32,
            y: self.y + delta_y as i32,
        }
    }
}

#[derive(Default, Clone)]
struct Guard {
    pos: Pos,
    dir: Direction,
}

#[derive(Clone)]
struct Map {
    rows: usize,
    columns: usize,
    obstructions: HashSet<Pos>,
    guard: Guard,
}

impl std::fmt::Debug for Map {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        for y in 0..self.rows {
            for x in 0..self.columns {
                let pos = Pos {
                    x: x as i32,
                    y: y as i32,
                };
                let ch = if self.obstructions.contains(&pos) {
                    '#'
                } else if self.guard.pos == pos {
                    self.guard.dir.into()
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
            let pos = Pos {
                x: x as i32,
                y: y as i32,
            };
            match ch {
                '#' => {
                    let _ = obstructions.insert(pos);
                }
                '.' => (),
                dir => {
                    guard.pos = pos;
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
    }
}

fn delta(dir: Direction) -> (i8, i8) {
    match dir {
        Direction::Up => (0, -1),
        Direction::Right => (1, 0),
        Direction::Down => (0, 1),
        Direction::Left => (-1, 0),
    }
}

fn within_map(pos: Pos, map: &Map) -> bool {
    pos.x >= 0 && pos.y >= 0 && pos.x < map.columns as i32 && pos.y < map.rows as i32
}

fn next_pos(map: &Map) -> Option<Pos> {
    let next = map.guard.pos + delta(map.guard.dir);
    if within_map(next, map) {
        Some(next)
    } else {
        None
    }
}

fn next_dir(dir: Direction) -> Direction {
    match dir {
        Direction::Up => Direction::Right,
        Direction::Right => Direction::Down,
        Direction::Down => Direction::Left,
        Direction::Left => Direction::Up,
    }
}

fn move_util_outside(map: &mut Map) -> HashSet<Pos> {
    let mut visited = HashSet::new();
    visited.insert(map.guard.pos);
    while let Some(next) = next_pos(map) {
        if map.obstructions.contains(&next) {
            map.guard.dir = next_dir(map.guard.dir);
        } else {
            visited.insert(next);
            map.guard.pos = next;
        }
    }

    visited
}

fn part1(lines: impl Iterator<Item = String>) -> usize {
    let mut map = parse_map(lines);
    move_util_outside(&mut map).len()
}

fn next_collision(map: &Map) -> Option<Pos> {
    let mut pos = map.guard.pos;
    loop {
        let next = pos + delta(map.guard.dir);
        if !within_map(next, map) {
            return None;
        }

        if map.obstructions.contains(&next) {
            return Some(pos);
        }

        pos = next;
    }
}

fn creates_loop(obstruction_pos: Pos, map: &Map) -> bool {
    // create a new map with an extra obstruction
    let mut obstructions = map.obstructions.clone();
    obstructions.insert(obstruction_pos);
    let mut map = Map {
        rows: map.rows,
        columns: map.columns,
        guard: map.guard.clone(),
        obstructions,
    };

    let mut collisions = HashSet::new();
    while let Some(pos) = next_collision(&map) {
        map.guard.pos = pos;
        if !collisions.insert((map.guard.pos, map.guard.dir)) {
            return true;
        }
        map.guard.dir = next_dir(map.guard.dir);
    }
    false
}

fn part2(lines: impl Iterator<Item = String>) -> usize {
    let mut loops = 0;
    let map = parse_map(lines);
    let starting_pos = map.guard.pos;
    for pos in move_util_outside(&mut map.clone()) {
        if pos != starting_pos && creates_loop(pos, &map) {
            loops += 1;
        }
    }

    loops
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
    println!(
        "Part 2 (example data): {}",
        part2(util::lines_iter(example_data))
    );
    println!(
        "Part 2: {}", // 2151 is too high
        part2(util::lines_iter(File::open("input/input_06.txt").unwrap()))
    );
}
