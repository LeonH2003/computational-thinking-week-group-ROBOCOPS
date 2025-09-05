package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
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
	// Input and output file paths
	inputFile := "data2.json"
	outputFile := "data3.csv"

	// Read the JSON file
	jsonFile, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening JSON file: %v\n", err)
		return
	}
	defer jsonFile.Close()

	// Parse the JSON data
	var people People
	if err := json.NewDecoder(jsonFile).Decode(&people); err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		return
	}

	// Create the CSV file
	csvFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Error creating CSV file: %v\n", err)
		return
	}
	defer csvFile.Close()

	// Write the CSV header
	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	header := []string{"Name", "Technical Skills", "Soft Skills", "Business Skills", "Creative Skills", "Academic Skills", "Average Skills"}
	if err := writer.Write(header); err != nil {
		fmt.Printf("Error writing CSV header: %v\n", err)
		return
	}

	// Write the data rows
	for _, person := range people.People {
		averageSkills := (person.TechnicalSkills + person.SoftSkills + person.BusinessSkills +
			person.CreativeSkills + person.AcademicSkills) / 5.0

		row := []string{
			person.Name,
			fmt.Sprintf("%.2f", person.TechnicalSkills),
			fmt.Sprintf("%.2f", person.SoftSkills),
			fmt.Sprintf("%.2f", person.BusinessSkills),
			fmt.Sprintf("%.2f", person.CreativeSkills),
			fmt.Sprintf("%.2f", person.AcademicSkills),
			fmt.Sprintf("%.2f", averageSkills),
		}

		if err := writer.Write(row); err != nil {
			fmt.Printf("Error writing CSV row: %v\n", err)
			return
		}
	}

	fmt.Printf("CSV file '%s' created successfully.\n", outputFile)
}
