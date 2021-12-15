use std::{
    fs::File,
    io::{BufRead, BufReader},
};

macro_rules! scan {
    ( $string:expr, $sep:expr, $( $x:ty ),+ ) => {{
        let mut iter = $string.split($sep);
        ($(iter.next().and_then(|word| word.parse::<$x>().ok()),)*)
    }}
}

fn main() {
    match File::open("./values") {
        Ok(file) => {
            let buffer = BufReader::new(file);
            let mut h = 0;
            let mut d = 0;

            for line in buffer.lines() {
                let l = line.unwrap();
                let (so, io) = scan!(l, char::is_whitespace, String, i32);
                let (s, i) = (so.unwrap(), io.unwrap());

                match s.as_ref() {
                    "forward" => h += i,
                    "down" => d += i,
                    "up" => d -= i,
                    _ => {}
                };
            }

            println!("Final value: {}", h * d);
        }
        Err(err) => println!("[ERROR] {}", err),
    }
}
