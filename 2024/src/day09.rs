use std::fs::File;
use std::option::Option;

use crate::util;

type DiskMap = Vec<BlockType>;

#[derive(Debug, Copy, Clone, PartialEq, Eq)]
enum BlockType {
    File(usize),
    FreeSpace,
}

fn push_file(disk: &mut Vec<BlockType>, id: usize, mut size: u8) {
    while size > 0 {
        disk.push(BlockType::File(id));
        size -= 1;
    }
}

fn push_free_blocks(disk: &mut Vec<BlockType>, mut size: u8) {
    while size > 0 {
        disk.push(BlockType::FreeSpace);
        size -= 1;
    }
}

fn parse(line: &str) -> DiskMap {
    let mut disk = Vec::new();
    let mut bt = BlockType::File(0);
    let mut next_id = 0;
    for ch in line.chars() {
        let size = ch.to_digit(10).unwrap();
        match bt {
            BlockType::File(id) => {
                push_file(&mut disk, id, size as u8);
                bt = BlockType::FreeSpace;
                next_id = id + 1;
            }
            BlockType::FreeSpace => {
                push_free_blocks(&mut disk, size as u8);
                bt = BlockType::File(next_id);
            }
        }
    }
    disk
}

fn find_next_free_block(map: &DiskMap, mut pos: usize) -> usize {
    while BlockType::FreeSpace != map[pos] {
        pos += 1;
    }

    pos
}

fn find_file_block(map: &DiskMap, mut pos: usize) -> usize {
    while BlockType::FreeSpace == map[pos] {
        pos -= 1;
    }

    pos
}

fn compact(map: &mut DiskMap) {
    let mut start_pos = 0;
    let mut end_pos = map.len() - 1;

    loop {
        start_pos = find_next_free_block(map, start_pos);
        end_pos = find_file_block(map, end_pos);
        if start_pos >= end_pos {
            return;
        }

        map[start_pos] = map[end_pos];
        map[end_pos] = BlockType::FreeSpace;
    }
}

fn checksum(map: &DiskMap) -> usize {
    let mut sum = 0;
    for (pos, block) in map.iter().enumerate() {
        if let BlockType::File(id) = block {
            sum += pos * id;
        }
    }

    sum
}

fn find_fitting_free_area(map: &DiskMap, size: usize) -> Option<usize> {
    let mut start_pos = 0;
    loop {
        while start_pos < map.len() && map[start_pos] != BlockType::FreeSpace {
            start_pos += 1;
        }

        if start_pos >= map.len() {
            return None;
        }

        let mut free_size = 1;
        while start_pos + free_size < map.len()
            && BlockType::FreeSpace == map[start_pos + free_size]
        {
            free_size += 1;
        }

        if free_size >= size {
            return Some(start_pos);
        }

        // move on
        start_pos += 1;
    }
}

fn find_file(map: &DiskMap, mut pos: usize) -> (usize, usize) {
    let file_id;
    loop {
        match map[pos] {
            BlockType::FreeSpace => {
                if pos == 0 {
                    return (0, 0);
                } else {
                    pos -= 1
                }
            }
            BlockType::File(id) => {
                file_id = id;
                break;
            }
        }
    }

    let mut size = 1;
    while pos >= size {
        match map[pos - size] {
            BlockType::FreeSpace => break,
            BlockType::File(id) => {
                if file_id != id {
                    break;
                }
            }
        }
        size += 1;
    }
    let start_pos = pos + 1 - size;

    (start_pos, size)
}

fn compact_v2(map: &mut DiskMap) {
    let mut file_pos = map.len();
    let mut file_size;

    while file_pos > 0 {
        (file_pos, file_size) = find_file(map, file_pos - 1);
        if file_size == 0 {
            return;
        }
        let free_pos = match find_fitting_free_area(map, file_size) {
            Some(free_pos) => free_pos,
            None => continue,
        };

        if free_pos >= file_pos {
            continue;
        }

        for i in 0..file_size {
            map[free_pos + i] = map[file_pos + i];
            map[file_pos + i] = BlockType::FreeSpace;
        }
    }
}

fn part1(mut map: DiskMap) -> usize {
    compact(&mut map);
    checksum(&map)
}

fn part2(mut map: DiskMap) -> usize {
    compact_v2(&mut map);
    checksum(&map)
}

pub fn run() {
    println!(
        "Part 1: {}",
        part1(parse(
            &util::lines_iter(File::open("input/input_09.txt").unwrap())
                .next()
                .unwrap()
        ))
    );
    println!(
        "Part 2: {}",
        part2(parse(
            &util::lines_iter(File::open("input/input_09.txt").unwrap())
                .next()
                .unwrap()
        ))
    );
}

#[cfg(test)]
mod test {
    use super::*;

    const EXAMPLE_DATA: &'static [u8] = b"2333133121414131402";

    #[test]
    fn part1() {
        assert_eq!(
            super::part1(parse(&util::lines_iter(EXAMPLE_DATA).next().unwrap())),
            1928
        );
    }

    #[test]
    fn part2() {
        assert_eq!(
            super::part2(parse(&util::lines_iter(EXAMPLE_DATA).next().unwrap())),
            2858
        );
    }
}
