use std::fs::File;

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

fn find_next_free(map: &DiskMap, mut pos: usize) -> usize {
    while BlockType::FreeSpace != map[pos] {
        pos += 1;
    }

    pos
}

fn find_file(map: &DiskMap, mut pos: usize) -> usize {
    while BlockType::FreeSpace == map[pos] {
        pos -= 1;
    }

    pos
}

fn compact(map: &mut DiskMap) {
    let mut start_pos = 0;
    let mut end_pos = map.len() - 1;

    loop {
        start_pos = find_next_free(map, start_pos);
        end_pos = find_file(map, end_pos);
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
        } else {
            return sum;
        }
    }

    sum
}

fn part1(mut map: DiskMap) -> usize {
    compact(&mut map);
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
    /*
    println!(
        "Part 2: {}",
        part2(util::lines_iter(File::open("input/input_09.txt").unwrap()))
    );
    */
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
}
