use core::fmt;
use std::error::Error;

#[derive(Debug)]
pub struct DynError {
    details: String
}

impl DynError {
    pub fn new(msg: &str) -> DynError {
        DynError { details: msg.to_string() }
    }
}

impl fmt::Display for DynError {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{}", self.details)
    }
}

impl Error for DynError {
    fn description(&self) -> &str {
        &self.details
    }
}
