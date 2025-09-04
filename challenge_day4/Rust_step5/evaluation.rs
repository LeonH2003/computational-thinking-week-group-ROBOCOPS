use std::fs::File;
use std::io::{self, BufRead, BufReader, Write};
use std::path::Path;
use std::fs::OpenOptions;

fn main() -> io::Result<()> {
    // Input and output paths
    let input_path = Path::new("data5.txt");
    let output_path = Path::new("data6.txt");

    // Open input
    let file = File::open(&input_path)?;
    let reader = BufReader::new(file);

    // Open output
    let mut output = OpenOptions::new()
        .write(true)
        .create(true)
        .truncate(true)  // clear existing file
        .open(&output_path)?;

    for (indexing, line) in reader.lines().enumerate() {
        let line = line?;
        let mut parts: Vec<&str> = line.split(',').collect();

        if indexing == 0 {
            // Header: add Evaluation column
            writeln!(output, "{},Evaluation", line)?;
            continue;
        }

        // Fill missing fields if any
        while parts.len() < 7 {
            parts.push("N/A");
        }

        // Calculate Evaluation from skills 1..5
        let mut total_score = 0;
        let mut num_skills = 0;
        for &skill in &parts[1..=5] {
            match skill {
                "low" => { total_score += 2; num_skills += 1; },
                "middle" => { total_score += 3; num_skills += 1; },
                "good" => { total_score += 4; num_skills += 1; },
                "super" => { total_score += 5; num_skills += 1; },
                _ => {},
            }
        }

        let evaluation = if num_skills > 0 {
            (total_score as f32) / (num_skills as f32)
        } else {
            0.0
        };

        // Write original line + evaluation
        writeln!(output, "{},{}", parts.join(","), evaluation)?;
    }

    println!("data6.txt created successfully!");
    Ok(())
}
