use std::fs::File;
use std::io::{BufRead, BufReader, Lines, Result};
use std::path::Path;

const RED_CUBES: i32 = 12;
const GREEN_CUBES: i32 = 13;
const BLUE_CUBES: i32 = 14;

fn main() {
    let mut sum = 0;
    let mut sum_multiplied_min = 0;
    let lines = read_line("./input2.txt").unwrap();
    for line in lines {
        if let Ok(text_line) = line {
            let parse_result = parse_line(&text_line);
            if parse_result.is_valid {
                sum += parse_result.game_id;
            }
            sum_multiplied_min += parse_result.multiply_min;
        }
    }
    println!("\n\nSum: {} , Sum multiplied min: {}", sum, sum_multiplied_min);
}

fn parse_line(line: &String) -> ParseResult {
    let mut is_valid = true;
    let mut min_red = 0;
    let mut min_green = 0;
    let mut min_blue = 0;
    let column = line.split(':').collect::<Vec<&str>>();
    let game_id = column[0].split(' ').collect::<Vec<&str>>()[1]
        .parse::<i32>()
        .unwrap();
    println!("\n================");
    println!("\nGame {}", game_id);
    let grabs = column[1]
        .split(';')
        .map(|game| game.trim())
        .collect::<Vec<&str>>();
    for grab in grabs.into_iter() {
        println!("----- Grab -----");
        let cubes = grab
            .split(',')
            .map(|cube| cube.trim())
            .collect::<Vec<&str>>();
        for cube in cubes.iter() {
            let number_color = cube
                .split(' ')
                .map(|token| token.trim())
                .collect::<Vec<&str>>();
            let number = number_color[0].parse::<i32>().unwrap();
            let is_invalid: bool;
            match number_color[1] {
                "blue" => { 
                    is_invalid = number > BLUE_CUBES;
                    if number > min_blue {
                        min_blue = number;
                    }
                }
                "red" => {
                    is_invalid = number > RED_CUBES;
                    if number > min_red {
                        min_red = number;
                    }
                }
                "green" => {
                    is_invalid = number > GREEN_CUBES;
                    if number > min_green {
                        min_green = number;
                    }
                }
                _ => { is_invalid = true; }
            };
            println!("{:?}", number_color);
            if is_invalid {
                is_valid = false;
            }
        }
    }
    let multiply_min = min_red * min_green * min_blue;
    println!("min. red: {}, min. green: {}, min. blue {}, multiplied: {}", min_red, min_green, min_blue, multiply_min);
    match is_valid {
        true => println!("\n ✔ OK"),
        false => println!("\n ❌ Invalid; do not check further")
    }
    ParseResult { game_id, is_valid, multiply_min }
}

struct ParseResult {
    game_id: i32,
    is_valid: bool,
    multiply_min: i32
}

fn read_line<T: AsRef<Path>>(filename: T) -> Result<Lines<BufReader<File>>> {
    let file = File::open(filename)?;
    let bufreader = BufReader::new(file);
    Ok(bufreader.lines())
}
