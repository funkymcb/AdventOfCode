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

    let mut game = Game::new();
    let mut sum_of_ids: u16 = 0;

    for line in reader.lines() {
        if line.is_ok() {
            let line: String = line.unwrap();

            // star1
            game.get_max_cube_counts(line.as_str());
            game.check_validity();

            if game.valid {
                sum_of_ids += game.id as u16
            }

            println!("{}", &game);
        } else {
            println!("could not read line");
            exit(1)
        }
    }

    println!("{}", sum_of_ids)
}
