mod game;

use crate::game::Game;
use std::fs::File;
use std::io::{BufRead, BufReader};
use std::process::exit;

const INPUT_PATH: &str = "input";

fn main() {
    let file = File::open(INPUT_PATH)
        .expect("could not read file");
    let reader = BufReader::new(file);

    for line in reader.lines() {
        if line.is_ok() {
            let line: String = line.unwrap();

            // star1
            let mut game = Game::new();
            game.analyze(line.as_str());
            println!("{}", &game);
        } else {
            println!("could not read line");
            exit(1)
        }
    }
}
