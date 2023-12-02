mod game;
mod error;

use crate::error::DynError;
use crate::game::Game;
use regex::Regex;
use std::fs::File;
use std::io::{BufRead, BufReader};
use std::process::exit;

const INPUT_PATH: &str = "input";
const GAME_REGEX: &str = r"(Game (?P<id>\d+))";

fn main() {
    let file = File::open(INPUT_PATH)
        .expect("could not read file");
    let reader = BufReader::new(file);

    for line in reader.lines() {
        if line.is_ok() {
            let line: String = line.unwrap();

            // star1
            let game = analyze_game(line.as_str());
            println!("{}", game.unwrap())
        } else {
            println!("could not read line");
            exit(1)
        }
    }
}

fn analyze_game(line: &str) -> Result<Game, DynError> {
   let re = Regex::new(GAME_REGEX).unwrap();
   let Some(caps) = re.captures(line) else {
       return Err(DynError::new("no line matched regex"));
   };

   Ok(Game {
       id: caps["id"].parse::<u8>().unwrap(),
       red_cube_count: 1,
       green_cubes_count: 2,
       blue_cubes_count: 3
   })
}
