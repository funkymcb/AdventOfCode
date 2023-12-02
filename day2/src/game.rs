use core::fmt;
use std::process::exit;
use regex::Regex;

const GAME_REGEX: &str = r"(Game (?P<id>\d+))";

#[derive(Debug)]
pub struct Game {
    pub id: u8,
    pub red_cube_count: u8,
    pub green_cubes_count: u8,
    pub blue_cubes_count: u8,
}

impl Game {
    pub fn new() -> Game {
        Game {
            id: 0,
            red_cube_count: 0,
            green_cubes_count: 0,
            blue_cubes_count: 0
        }
    }

    pub fn analyze(&mut self, line: &str) {
       let re = Regex::new(GAME_REGEX).unwrap();
       let Some(caps) = re.captures(line) else {
           println!("no line matched regex");
           exit(1)
       };

       self.id = caps["id"].parse::<u8>().unwrap();
    }
}

impl fmt::Display for Game {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "game id: {}\n red cubes: {}\n green cubes: {}\n blue cubes: {}",
           self.id,
           self.red_cube_count,
           self.green_cubes_count,
           self.blue_cubes_count
        )
    }
}
