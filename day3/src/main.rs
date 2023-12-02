use std::fs::File;
use std::io::{BufRead, BufReader};
use std::process::exit;

const INPUT_PATH: &str = "input";

fn main() {
    let file = File::open(INPUT_PATH).expect("could not read file");
    let reader = BufReader::new(file);

    for line in reader.lines() {
        if line.is_ok() {

        } else {
            println!("could not read line");
            exit(1)
        }
    }
}
