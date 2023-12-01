use std::fs::File;
use std::io::{BufRead, BufReader};

const INPUT_PATH: &str = "input";

fn main() {
    let file = File::open(INPUT_PATH)
        .expect("could not read file");
    let reader = BufReader::new(file);

    let mut sum: u32 = 0;

    for line in reader.lines() {
        if line.is_ok() {
            let calibration_value: u32 = read_calibration_value(line.unwrap());
            sum += calibration_value;
        } else {
            println!("could not read line")
        }
    }

    println!("the sum of all calibration values is: {}", sum)
}

fn read_calibration_value(line: String) -> u32 {
    let mut digits: Vec<u32> = Vec::new();

    for c in line.chars() {
        if c.is_alphabetic() {
            continue;
        }
        
        let digit = c.to_digit(10).unwrap();
        digits.push(digit);
    }

    // duplicate value if there is only one digit
    if digits.len() == 1 {
        digits.push(digits[0]);
    }

    let calibration_value = u32::from(digits[0] * 10  + digits.last().unwrap());

    calibration_value
}
