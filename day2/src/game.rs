use core::fmt;
use regex::Regex;

const GAME_ID_REGEX: (&str, &str) = ("id", r"(Game (?P<id>\d+))");
const RED_CUBE_REGEX: (&str, &str) = ("red", r"(?P<red>\d+) red");
const GREEN_CUBE_REGEX: (&str, &str) = ("green", r"(?P<green>\d+) green");
const BLUE_CUBE_REGEX: (&str, &str) = ("blue", r"(?P<blue>\d+) blue");

const RED_CUBE_MAX: u8 = 12;
const GREEN_CUBE_MAX: u8 = 13;
const BLUE_CUBE_MAX: u8 = 14;

#[derive(Debug)]
pub struct Game {
    pub id: u8,
    pub red_cube_max: u8,
    pub green_cube_max: u8,
    pub blue_cube_max: u8,
    pub valid: bool
}

impl Game {
    pub fn new() -> Game {
        Game {
            id: 0,
            red_cube_max: 0,
            green_cube_max: 0,
            blue_cube_max: 0,
            valid: false
        }
    }

    pub fn get_max_cube_counts(&mut self, line: &str) {
        let id: u8 = capture_regex(line, GAME_ID_REGEX);
        let red_max: u8 = capture_regex(line, RED_CUBE_REGEX);
        let green_max: u8 = capture_regex(line, GREEN_CUBE_REGEX);
        let blue_max: u8 = capture_regex(line, BLUE_CUBE_REGEX);

        self.id = id;
        self.red_cube_max = red_max;
        self.green_cube_max = green_max;
        self.blue_cube_max = blue_max;
    }

    pub fn check_validity(&mut self) {
        if self.red_cube_max > RED_CUBE_MAX || self.green_cube_max > GREEN_CUBE_MAX || self.blue_cube_max > BLUE_CUBE_MAX {
            self.valid = false
        } else {
                    self.valid = true
        }
    }
}

impl fmt::Display for Game {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "game id: {}\n red cubes: {}\n green cubes: {}\n blue cubes: {}\n valid game: {}",
               self.id,
               self.red_cube_max,
               self.green_cube_max,
               self.blue_cube_max,
               self.valid
              )
    }
}

fn capture_regex(line: &str, regex: (&str, &str)) -> u8 {
    let re = Regex::new(regex.1).unwrap();

    let caps = re.captures_iter(line);
    let mut max: u8 = 0;

    for c in caps {
        let n = c[regex.0].parse::<u8>().unwrap();
        if n > max {
            max = n
        }
    }

    max
}
