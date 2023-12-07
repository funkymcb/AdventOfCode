mod game;

use crate::game::Game;
use std::fs::File;
use std::io::{BufRead, BufReader};
use std::process::exit;

const INPUT_PATH: &str = "input";

fn main() {
    let file = File::open(INPUT_PATH).expect("could not read file");
    let reader = BufReader::new(file);

    let mut sum: u16 = 0;

    for line in reader.lines() {
        let mut game = Game::default();

        if line.is_ok() {
            game.init(line.unwrap().as_str());
            game.calculate_score();

            println!("{}", game);
            sum += game.score;
        } else {
            println!("could not read line");
            exit(1)
        }
    }

    println!("sum of scores: {}", sum)
}
