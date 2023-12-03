use std::collections::HashMap;
use std::fs::File;
use std::io::{BufRead, BufReader, Lines, Result};
use std::path::Path;


fn main() {
    let mut sum: u32 = 0;
    let number_words: HashMap<&str, u32> = HashMap::from([
        ("one", 1),
        ("two", 2),
        ("three", 3),
        ("four", 4),
        ("five", 5),
        ("six", 6),
        ("seven", 7),
        ("eight", 8),
        ("nine", 9),
        ("1", 1),
        ("2", 2),
        ("3", 3),
        ("4", 4),
        ("5", 5),
        ("6", 6),
        ("7", 7),
        ("8", 8),
        ("9", 9),
    ]);
    let lines = read_line("./input2.txt").unwrap();
    for line in lines {
        if let Ok(text_line) = line {
            let number = get_number(&text_line, &number_words);
            sum += number;
        }
    }
    println!("Sum is {}", sum);
}

fn read_line<T: AsRef<Path>>(filename: T) -> Result<Lines<BufReader<File>>> {
    let file = File::open(filename)?;
    let bufreader = BufReader::new(file);
    Ok(bufreader.lines())
}

fn get_number(line: &str, number_words: &HashMap<&str, u32>) -> u32 {
        let mut position_digits = HashMap::new();
        for (number_word, value) in number_words.iter() {
            let indexes: Vec<(usize, &str)> = line.match_indices(number_word).collect();
            for (index, _) in indexes {
                position_digits.insert(index, value);
            }
        }
        let (_, &first_digit) = position_digits.iter().min().unwrap();
        let (_, &last_digit) = position_digits.iter().max().unwrap();
        let number = first_digit * 10 + last_digit;
        println!("{:?}, {:?}", position_digits, number);
        number
}
