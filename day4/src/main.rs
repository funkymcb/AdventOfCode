mod game;

use crate::game::Game;
use std::fs::File;
use std::io::{BufRead, BufReader};
use std::process::exit;

const INPUT_PATH: &str = "input";

fn main() {
    let file = File::open(INPUT_PATH).expect("could not read file");
    let reader = BufReader::new(file);

    let mut sum_of_scores: u16 = 0;
    let mut games: Vec<Game> = Vec::new();

    // star 1
    for (i, line) in reader.lines().enumerate() {
        let mut game = Game::default();

        if line.is_ok() {
            game.init(i, line.unwrap().as_str());
            game.calculate_score();

            sum_of_scores += game.score;
            games.push(game);
        } else {
            println!("could not read line");
            exit(1)
        }
    }
    println!("sum star 1: {}", sum_of_scores);

    // star 2
    let mut i: usize = 0;
    loop {
        println!("i: {}, game: {}", i, games[i]);
        if games[i].id < 203 {
            if games[i].score > 0 {
                for j in i..(usize::try_from(games[i].count_of_wins).unwrap() + i) {
                    println!("j: {}", j);
                    games.insert(i+1, games[j].clone())
                }
            }
        } else {
            break
        }
        i = i + 1
    }

    println!("total scratchcards: {}", games.len())
}
