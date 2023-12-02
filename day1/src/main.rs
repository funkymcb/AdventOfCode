use phf::phf_map;
use std::fs::File;
use std::io::{BufRead, BufReader};
use std::process::exit;

const INPUT_PATH: &str = "input";

const NUMBER_STRINGS: phf::Map<&str, &str> = phf_map! {
    "one" => "o1e",
    "two" => "t2o",
    "three" => "t3e",
    "four" => "f4r",
    "five" => "f5e",
    "six" => "s6x",
    "seven" => "s7n",
    "eight" => "e8t",
    "nine" => "n9e",
};

fn main() {
    let file = File::open(INPUT_PATH)
        .expect("could not read file");
    let reader = BufReader::new(file);

    let mut sum_one: u32 = 0;
    let mut sum_two: u32 = 0;

    for line in reader.lines() {
        if line.is_ok() {
            let line = line.unwrap();
            // star 1
            let calibration_value_one: u32 = read_calibration_value(&line); 
            sum_one += calibration_value_one;

            // star 2
            let fixed_line: String = fix_spelling_digits(&line);
            let calibration_value_two: u32 = read_calibration_value(&fixed_line); 
            sum_two += calibration_value_two;

        } else {
            println!("could not read line");
            exit(1)
        }
    }

    println!("(star 1) the sum of all calibration values is: {}", sum_one);
    println!("(star 2) the sum of all fixed calibration values is: {}", sum_two);
}

fn read_calibration_value(line: &String) -> u32 {
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

fn fix_spelling_digits(line: &String) -> String {
    let mut fixed_line = line.to_string();

    for (s, d) in NUMBER_STRINGS.into_iter() {
        fixed_line = fixed_line.replace(s, d);
    }

    fixed_line
}
