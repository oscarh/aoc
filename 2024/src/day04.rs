use std::fs::File;

use crate::util;

struct Matrix {
    rows: usize,
    columns: usize,
    v: Vec<Vec<char>>,
}

impl Matrix {
    fn new(lines: impl Iterator<Item = String>) -> Self {
        let v = lines
            .map(|line| line.chars().collect::<Vec<_>>())
            .collect::<Vec<Vec<_>>>();
        let rows = v.len();
        let columns = v[0].len();
        Self { rows, columns, v }
    }
}

const XMAS_LEN: usize = 4;

fn search_left(m: &Matrix, x: usize, y: usize) -> u64 {
    if x >= XMAS_LEN - 1 && m.v[y][x - 1] == 'M' && m.v[y][x - 2] == 'A' && m.v[y][x - 3] == 'S' {
        1
    } else {
        0
    }
}

fn search_diagonal_left_up(m: &Matrix, x: usize, y: usize) -> u64 {
    if x >= XMAS_LEN - 1
        && y >= XMAS_LEN - 1
        && m.v[y - 1][x - 1] == 'M'
        && m.v[y - 2][x - 2] == 'A'
        && m.v[y - 3][x - 3] == 'S'
    {
        1
    } else {
        0
    }
}

fn search_up(m: &Matrix, x: usize, y: usize) -> u64 {
    if y >= XMAS_LEN - 1 && m.v[y - 1][x] == 'M' && m.v[y - 2][x] == 'A' && m.v[y - 3][x] == 'S' {
        1
    } else {
        0
    }
}

fn search_diagonal_right_up(m: &Matrix, x: usize, y: usize) -> u64 {
    if x <= m.columns - XMAS_LEN
        && y >= XMAS_LEN - 1
        && m.v[y - 1][x + 1] == 'M'
        && m.v[y - 2][x + 2] == 'A'
        && m.v[y - 3][x + 3] == 'S'
    {
        1
    } else {
        0
    }
}

fn search_right(m: &Matrix, x: usize, y: usize) -> u64 {
    if x <= m.columns - XMAS_LEN
        && m.v[y][x + 1] == 'M'
        && m.v[y][x + 2] == 'A'
        && m.v[y][x + 3] == 'S'
    {
        1
    } else {
        0
    }
}

fn search_diagonal_right_down(m: &Matrix, x: usize, y: usize) -> u64 {
    if x <= m.columns - XMAS_LEN
        && y <= m.rows - XMAS_LEN
        && m.v[y + 1][x + 1] == 'M'
        && m.v[y + 2][x + 2] == 'A'
        && m.v[y + 3][x + 3] == 'S'
    {
        1
    } else {
        0
    }
}

fn search_down(m: &Matrix, x: usize, y: usize) -> u64 {
    if y <= m.rows - XMAS_LEN
        && m.v[y + 1][x] == 'M'
        && m.v[y + 2][x] == 'A'
        && m.v[y + 3][x] == 'S'
    {
        1
    } else {
        0
    }
}

fn search_diagonal_left_down(m: &Matrix, x: usize, y: usize) -> u64 {
    if x >= XMAS_LEN - 1
        && y <= m.rows - XMAS_LEN
        && m.v[y + 1][x - 1] == 'M'
        && m.v[y + 2][x - 2] == 'A'
        && m.v[y + 3][x - 3] == 'S'
    {
        1
    } else {
        0
    }
}

fn search_from(m: &Matrix, x: usize, y: usize) -> u64 {
    search_left(m, x, y)
        + search_diagonal_left_up(m, x, y)
        + search_up(m, x, y)
        + search_diagonal_right_up(m, x, y)
        + search_right(m, x, y)
        + search_diagonal_right_down(m, x, y)
        + search_down(m, x, y)
        + search_diagonal_left_down(m, x, y)
}

fn part1(lines: impl Iterator<Item = String>) -> u64 {
    let matrix = Matrix::new(lines);
    let mut count = 0;
    for (y, row) in matrix.v.iter().enumerate() {
        for (x, value) in row.iter().enumerate() {
            if *value == 'X' {
                count += search_from(&matrix, x, y);
            }
        }
    }
    count
}

fn search_x_mas_from(m: &Matrix, x: usize, y: usize) -> bool {
    let top_left_to_bottom_right = m.v[y - 1][x - 1] == 'M' && m.v[y + 1][x + 1] == 'S'
        || m.v[y - 1][x - 1] == 'S' && m.v[y + 1][x + 1] == 'M';

    let bottom_left_to_top_right = m.v[y - 1][x + 1] == 'M' && m.v[y + 1][x - 1] == 'S'
        || m.v[y - 1][x + 1] == 'S' && m.v[y + 1][x - 1] == 'M';

    top_left_to_bottom_right && bottom_left_to_top_right
}

fn part2(lines: impl Iterator<Item = String>) -> u64 {
    let matrix = Matrix::new(lines);
    let mut count = 0;
    for (y, row) in matrix.v.iter().enumerate() {
        // skip first and the last row
        if y == 0 || y == matrix.rows - 1 {
            continue;
        }
        for (x, value) in row.iter().enumerate() {
            // skip first and the last ones
            if x == 0 || x == matrix.columns - 1 {
                continue;
            }

            if *value == 'A' && search_x_mas_from(&matrix, x, y) {
                count += 1;
            }
        }
    }
    count
}

pub fn run() {
    let mut example_data = &b"MMMSXXMASM\n\
           MSAMXMSMSA\n\
           AMXSXMAAMM\n\
           MSAMASMSMX\n\
           XMASAMXAMM\n\
           XXAMMXXAMA\n\
           SMSMSASXSS\n\
           SAXAMASAAA\n\
           MAMMMXMMMM\n\
           MXMXAXMASX\n"[..];
    println!(
        "Part 1 (example data): {}",
        part1(util::lines_iter(&mut example_data))
    );
    println!(
        "Part 1: {}",
        part1(util::lines_iter(File::open("input/input_04.txt").unwrap()))
    );

    let mut example_data = &b"MMMSXXMASM\n\
           MSAMXMSMSA\n\
           AMXSXMAAMM\n\
           MSAMASMSMX\n\
           XMASAMXAMM\n\
           XXAMMXXAMA\n\
           SMSMSASXSS\n\
           SAXAMASAAA\n\
           MAMMMXMMMM\n\
           MXMXAXMASX\n"[..];
    println!(
        "Part 2 (example data): {}",
        part2(util::lines_iter(&mut example_data))
    );
    println!(
        "Part 2: {}",
        part2(util::lines_iter(File::open("input/input_04.txt").unwrap()))
    );
}
