package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

// Problem statement: https://github.com/gophercises/quiz

type problem struct {
	question string
	answer   string
}

// readCSVFile reads the CSV file and returns all the records
func readCSVFile(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Failed to open CSV file: %s", filename)
		return nil, err
	}

	// Ensures that the file is closed when the function returns,
	// regardless of how it returns (e.g., even if there's an error).
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Failed to parse CSV file: %s", filename)
		return nil, err
	}

	return records, nil
}

func parseLines(lines [][]string) []problem {
	values := make([]problem, len(lines))
	for i, line := range lines {
		values[i] = problem{question: line[0], answer: line[1]}
	}
	return values
}

// processQuestions processes the questions from the CSV records
func processQuestions(problems []problem, getUserInput func() string) int {
	score := 0
	totalQuestions := len(problems)

	for i, record := range problems {
		// Print the first value
		fmt.Printf("Problem %d: %s = ", i+1, record.question)

		// Get user input
		userInput := getUserInput()

		if strings.TrimSpace(userInput) == record.answer {
			score++
		}
	}

	fmt.Printf("You scored %d out of %d\n", score, totalQuestions)
	return score
}

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}

func main() {
	filename := flag.String("csv", "problems.csv", "a CSV file in the format of [question, answer]")
	help := flag.Bool("h", false, "Show usage information")

	flag.Parse()

	if *help {
		flag.PrintDefaults()
		return
	}

	records, err := readCSVFile(*filename)
	if err != nil {
		exit("")
	}

	problems := parseLines(records)

	getUserInput := func() string {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		return input
	}

	processQuestions(problems, getUserInput)
}
