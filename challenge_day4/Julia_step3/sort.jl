using DataFrames, CSV, Statistics, DelimitedFiles

# Read CSV file into a DataFrame
people_df = CSV.File("data3.csv") |> DataFrame

# Function to classify a numeric score into skill level
function classify_score(score, quartiles)
    if score <= quartiles[1]
        return "low"
    elseif score <= quartiles[2]
        return "middle"
    elseif score <= quartiles[3]
        return "good"
    else
        return "super"
    end
end

# Iterate over each skill column (skip the first "Name" column)
for col_name in names(people_df)[2:end]
    col_data = people_df[!, col_name]

    # Compute quartiles from all numeric values
    quartiles = quantile(col_data, [0.25, 0.5, 0.75])

    # Replace each value with its category
    new_col = map(x -> classify_score(x, quartiles), col_data)
    people_df[!, col_name] = new_col
end

# Save to CSV (optional)
CSV.write("data4.csv", people_df)

# Save to TXT (comma-separated)
open("data4.txt", "w") do io
    # Write header
    println(io, join(names(people_df), ","))

    # Write each row
    for row in eachrow(people_df)
        println(io, join(row, ","))
    end
end
