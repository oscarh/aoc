use std::io::{self, BufRead};

pub fn lines_iter<D>(input: D) -> impl Iterator<Item = String>
where
    D: std::io::Read + Sized,
{
    io::BufReader::new(input)
        .lines()
        .map(|lines_res| lines_res.unwrap())
}
