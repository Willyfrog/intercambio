package main

import (
	"flag"
	"fmt"
)

func main() {
	input := flag.String("input", "./donations.csv", "Path to CSV file, default is ./donations.csv")
	output := flag.String("output", "./output.csv", "Path to ouput results, default is ./output.csv")
	emailTemplate := flag.String("template", "{.EmailFrom}, {.EmailTo}, {.Address}", "Template to fill in with the data")
	flag.Parse()
	fmt.Println("input: ", *input)
	fmt.Println("output: ", *output)
	fmt.Println("Template: ", *emailTemplate)
}
