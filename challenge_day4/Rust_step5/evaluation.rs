use std::fs::File;
use std::io::{self, BufRead, BufReader, Write};
use std::path::Path;
use std::fs::OpenOptions;

fn main() -> io::Result<()> {
    let input_path = Path::new("data5.txt");
    let output_path = Path::new("data6.txt");

    let file = File::open(&input_path)?;
    let reader = BufReader::new(file);

    let mut output = OpenOptions::new()
        .write(true)
        .create(true)
        .truncate(true)
        .open(&output_path)?;

    for (i, line) in reader.lines().enumerate() {
        let line = line?;

        if i == 0 {
            // header: just append Evaluation
            writeln!(output, "{},Evaluation", line)?;
            continue;
        }

        // Split naively on commas (works if your input has no quoted commas)
        // We will NOT write these parts back (to avoid injecting "N/A")
        let parts: Vec<&str> = line.split(',').collect();

        // Compute evaluation based on columns 1..=5 if present
        let mut total = 0u32;
        let mut n = 0u32;
        for idx in 1..=5 {
            if let Some(raw) = parts.get(idx) {
                let skill = raw.trim().to_ascii_lowercase();
                match skill.as_str() {
                    "low" =>    { total += 2; n += 1; }
                    "middle" => { total += 3; n += 1; }
                    "good" =>   { total += 4; n += 1; }
                    "super" =>  { total += 5; n += 1; }
