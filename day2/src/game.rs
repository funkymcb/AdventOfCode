use core::fmt;

#[derive(Debug)]
pub struct Game {
    pub id: u8,
    pub red_cube_count: u8,
    pub green_cubes_count: u8,
    pub blue_cubes_count: u8,
}

impl fmt::Display for Game {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "game id: {}\n red cubes: {}\n green cubes: {}\n blue cubes: {}",
               self.id, self.red_cube_count, self.green_cubes_count, self.blue_cubes_count)
    }
}
