package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
)

type Person struct {
	Name            string  `json:"name"`
	TechnicalSkills float64 `json:"Technical Skills"`
	SoftSkills      float64 `json:"Soft Skills"`
	BusinessSkills  float64 `json:"Business Skills"`
	CreativeSkills  float64 `json:"Creative Skills"`
	AcademicSkills  float64 `json:"Academic Skills"`
}

type People struct {
	People []Person `json:"people"`
}

func main() {
	// Read the JSON file
	jsonBytes, err := os.ReadFile("data2.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var people People
	// Try object-with-people first
	if err := json.Unmarshal(jsonBytes, &people); err != nil || len(people.People) == 0 {
		// fallback: maybe it's a top-level array
		var arr []Person
		if err2 := json.Unmarshal(jsonBytes, &arr); err2 == nil {
			people.People = arr
		} else {
			fmt.Println("Error parsing JSON:", err2)
			return
		}
	}

	// Create CSV
	csvFile, err := os.Create("data3.csv")
	if err != nil {
		fmt.Println("Error creating CSV:", err)
		return
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)

	// Write header
	if err := writer.Write([]string{
		"Name",
		"Technical Skills",
		"Soft Skills",
		"Business Skills",
		"Creative Skills",
		"Academic Skills",
	}); err != nil {
		fmt.Println("Error writing header:", err)
		return
	}

	// Write rows
	for _, p := range people.People {
		row := []string{
			p.Name,
			formatFloat(p.TechnicalSkills),
			formatFloat(p.SoftSkills),
			formatFloat(p.BusinessSkills),
			formatFloat(p.CreativeSkills),
			formatFloat(p.AcademicSkills),
		}
		if err := writer.Write(row); err != nil {
			fmt.Println("Error writing row:", err)
			return
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		fmt.Println("Error flushing CSV:", err)
	}
	fmt.Println("data3.csv created successfully.")
}

// formatFloat makes sure numbers match expected test output
func formatFloat(v float64) string {
	// Round to 4 decimal places
	r := math.Round(v*1e4) / 1e4
	s := strconv.FormatFloat(r, 'f', 4, 64) // fixed 4 dp
	s = strings.TrimRight(s, "0")           // remove trailing zeros
	s = strings.TrimRight(s, ".")           // remove trailing dot if integer
	return s
}
