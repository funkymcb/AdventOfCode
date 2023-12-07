use core::fmt;

#[derive(Debug, Default)]
pub struct Game {
    pub winning_numbers: Vec<u8>,
    pub scratched_numbers: Vec<u8>,
    pub score: u16
}

impl Game {
    pub fn init(&mut self, line: &str) {
        let n: &str = line.split(':').nth(1).unwrap();
        let w: &str = n.split('|').nth(0).unwrap();
        let s: &str = n.split('|').nth(1).unwrap();

        let winning_tokens: Vec<&str> = w.trim().split_whitespace().collect();
        let scratched_tokens: Vec<&str> = s.trim().split_whitespace().collect();

        let winning_numbers: Vec<u8> = winning_tokens.into_iter().filter_map(|x| x.parse().ok()).collect();
        let scratched_numbers: Vec<u8> = scratched_tokens.into_iter().filter_map(|x| x.parse().ok()).collect();

        self.winning_numbers = winning_numbers;
        self.scratched_numbers = scratched_numbers;
    }

    pub fn calculate_score(&mut self) {
        for w in &self.winning_numbers {
            for s in &self.scratched_numbers {
                if w == s {
                    if self.score == 0 {
                        self.score = 1
                    } else {
                        self.score = self.score * 2
                    }
                }
            }
        }
    }
}

impl fmt::Display for Game {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "winning numbers: {:?}\n score: {}",
               self.winning_numbers,
               self.score
               )
    }
}
