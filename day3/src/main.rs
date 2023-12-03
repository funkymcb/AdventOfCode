// FAIL: thats not working. no time to fix it :/
use std::fs::File;
use std::io::{BufRead, BufReader};

const INPUT_PATH: &str = "input";
const SYMBOLS: &str = "!\"§$%&/()[]{}=?+*#'-_:;,@€";
const SCHEMA_SIZE: usize = 140;

fn main() {
    let file = File::open(INPUT_PATH).expect("could not read file");
    let reader = BufReader::new(file);

    let mut engine_schema = [['\0'; SCHEMA_SIZE]; SCHEMA_SIZE];
 
    for (i, line) in reader.lines().enumerate() {
        for (j, char) in  line.unwrap().chars().enumerate(){
            engine_schema[i][j] = char;
        }
    }

    let sum: u32 = sum_of_part_numbers(engine_schema);
    println!("{}", sum)
}

fn sum_of_part_numbers(schema: [[char; SCHEMA_SIZE]; SCHEMA_SIZE]) -> u32 {
    let mut clone = schema.clone();
    let mut part_numbers: Vec<u32> = Vec::new();

    for (i, row) in clone.iter_mut().enumerate() {
        let mut n: u32 = 0;
        for (j, col) in row.iter_mut().enumerate() {
            if col.is_digit(10) {
                n = n*10+schema[i][j].to_digit(10).unwrap();

                if j > 0 && j < SCHEMA_SIZE {
                    let is_part: bool = check_for_adjacent_symbol(schema, i, j);
                
                    if j == SCHEMA_SIZE - 1 {
                        if is_part {
                            part_numbers.push(n)
                        }
                        continue
                    }

                    let next_char: char =  schema[i][j+1];
                    if !next_char.is_digit(10) || j == SCHEMA_SIZE - 1 {
                        if is_part {
                            part_numbers.push(n)
                        }
                    }
                }
            } else {
                n = 0
            } 
        }
        break
    }

    print!("{:?} ", part_numbers);
    part_numbers.iter().sum()
}

fn check_for_adjacent_symbol(schema: [[char; SCHEMA_SIZE]; SCHEMA_SIZE], i: usize, j: usize) -> bool {
    let mut is_part: bool = false;

    // check row below
    if SYMBOLS.contains(schema[i+1][j-1]) || SYMBOLS.contains(schema[i+1][j]) || SYMBOLS.contains(schema[i+1][j+1]) {
        is_part = true
    }

    is_part
}
