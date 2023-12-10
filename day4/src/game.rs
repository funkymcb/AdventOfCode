use core::fmt;

#[derive(Clone, Debug, Default)]
pub struct Game {
    pub id: usize,
    pub winning_numbers: Vec<u8>,
    pub scratched_numbers: Vec<u8>,
    pub count_of_wins: u8,
    pub score: u16
}

impl Game {
    pub fn init(&mut self, index: usize, line: &str) {
        let n: &str = line.split(':').nth(1).unwrap();
        let w: &str = n.split('|').nth(0).unwrap();
        let s: &str = n.split('|').nth(1).unwrap();

        let winning_tokens: Vec<&str> = w.trim().split_whitespace().collect();
        let scratched_tokens: Vec<&str> = s.trim().split_whitespace().collect();

        let winning_numbers: Vec<u8> = winning_tokens.into_iter().filter_map(|x| x.parse().ok()).collect();
        let scratched_numbers: Vec<u8> = scratched_tokens.into_iter().filter_map(|x| x.parse().ok()).collect();

        self.id = index;
        self.winning_numbers = winning_numbers;
        self.scratched_numbers = scratched_numbers;
    }

    pub fn calculate_score(&mut self) {
        let mut count: u8 = 0;

        for w in &self.winning_numbers {
            for s in &self.scratched_numbers {
                if w == s {
                    count += 1;
                    if self.score == 0 {
                        self.score = 1
                    } else {
                        self.score = self.score * 2
                    }
                }
            }
        }

        self.count_of_wins = count;
    }
}

impl fmt::Display for Game {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "id: {}\n winning numbers: {:?}\n count of wins {}\n score: {}",
               self.id,
               self.winning_numbers,
               self.count_of_wins,
               self.score
               )
    }
}
