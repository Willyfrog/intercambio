package main

import (
	"flag"
	"fmt"

	"github.com/willyfrog/intercambio/algorithm"
	input "github.com/willyfrog/intercambio/input"
	"github.com/willyfrog/intercambio/output"
)

func main() {
	inputFile := flag.String("input", "./donations.csv", "Path to CSV file, default is ./donations.csv")
	outputFile := flag.String("output", "./output.csv", "Path to ouput results, default is ./output.csv")
	emailTemplate := flag.String("template", "{.EmailFrom}, {.EmailTo}, {.Address}", "Template to fill in with the data")
	flag.Parse()
	fmt.Println("input: ", *inputFile)
	fmt.Println("output: ", *outputFile)
	fmt.Println("Template: ", *emailTemplate)
	origin, err := input.LoadCSV(*inputFile)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Loaded %v participants\n", len(origin))
	results, err := algorithm.Run(origin)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created %v matches\n", len(results))
	err = output.WriteCSV(results, *outputFile)
	if err != nil {
		panic(err)
	}

}
