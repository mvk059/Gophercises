package main

import (
	"io"
	"os"
	"testing"
)

func TestParseLines(t *testing.T) {
	records := [][]string{
		{"5+5", "10"},
		{"1+1", "2"},
		{"8+3", "11"},
		{"what 2+2, sir?", "4"},
	}

	expected := []problem{
		{question: "5+5", answer: "10"},
		{question: "1+1", answer: "2"},
		{question: "8+3", answer: "11"},
		{question: "what 2+2, sir?", answer: "4"},
	}

	problems := parseLines(records)

	if len(problems) != len(expected) {
		t.Errorf("parseLines returned %d problems, expected %d", len(problems), len(expected))
	}

	for i, p := range problems {
		if p.question != expected[i].question || p.answer != expected[i].answer {
			t.Errorf("parseLines returned incorrect problem at index %d: got %+v, expected %+v", i, p, expected[i])
		}
	}
}

func TestProcessQuestions(t *testing.T) {
	testCases := []struct {
		name           string
		problems       []problem
		getUserInput   func() string
		timeLimit      int
		expectedScore  int
		expectedOutput string
	}{
		{
			name: "All correct answers",
			problems: []problem{
				{question: "5+5", answer: "10"},
				{question: "1+1", answer: "2"},
				{question: "8+3", answer: "11"},
				{question: "what 2+2, sir?", answer: "4"},
			},
			getUserInput:   getMockUserInput([]string{"10", "2", "11", "4"}),
			timeLimit:      30,
			expectedScore:  4,
			expectedOutput: "Problem 1: 5+5 = Problem 2: 1+1 = Problem 3: 8+3 = Problem 4: what 2+2, sir? = You scored 4 out of 4\n",
		},
		{
			name: "All incorrect answers",
			problems: []problem{
				{question: "5+5", answer: "11"},
				{question: "1+1", answer: "3"},
				{question: "8+3", answer: "12"},
				{question: "what 2+2, sir?", answer: "5"},
			},
			getUserInput:   getMockUserInput([]string{"10", "2", "11", "4"}),
			timeLimit:      30,
			expectedScore:  0,
			expectedOutput: "Problem 1: 5+5 = Problem 2: 1+1 = Problem 3: 8+3 = Problem 4: what 2+2, sir? = You scored 0 out of 4\n",
		},
		{
			name: "Time limit exceeded",
			problems: []problem{
				{question: "5+5", answer: "10"},
				{question: "1+1", answer: "2"},
				{question: "8+3", answer: "11"},
				{question: "what 2+2, sir?", answer: "4"},
			},
			getUserInput:   getMockUserInput([]string{"10", "2"}),
			timeLimit:      1,
			expectedScore:  2,
			expectedOutput: "Problem 1: 5+5 = Problem 2: 1+1 = \nYou scored 2 out of 4\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			originalStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			score := processQuestions(tc.problems, tc.getUserInput, tc.timeLimit)

			w.Close()
			output, _ := io.ReadAll(r)
			os.Stdout = originalStdout

			if score != tc.expectedScore {
				t.Errorf("Expected score %d, but got %d", tc.expectedScore, score)
			}

			if string(output) != tc.expectedOutput {
				t.Errorf("Expected output: %q, but got: %q", tc.expectedOutput, string(output))
			}
		})
	}
}

func getMockUserInput(inputs []string) func() string {
	index := 0
	return func() string {
		if index >= len(inputs) {
			return ""
		}
		input := inputs[index]
		index++
		return input
	}
}
