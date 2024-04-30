package main

import (
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
	// Test case 1: All correct answers
	problems := []problem{
		{question: "5+5", answer: "10"},
		{question: "1+1", answer: "2"},
		{question: "8+3", answer: "11"},
		{question: "what 2+2, sir?", answer: "4"},
	}
	mockInputs := []string{"10", "2", "11", "4"}
	score := processQuestions(problems, getMockUserInput(mockInputs))
	if score != 4 {
		t.Errorf("processQuestions() with all correct answers should return 3, but got %d", score)
	}

	// Test case 2: All incorrect answers
	problems = []problem{
		{question: "5+5", answer: "11"},
		{question: "1+1", answer: "3"},
		{question: "8+3", answer: "12"},
		{question: "what 2+2, sir?", answer: "5"},
	}
	mockInputs = []string{"10", "2", "11", "4"}
	score = processQuestions(problems, getMockUserInput(mockInputs))
	if score != 0 {
		t.Errorf("processQuestions() with all incorrect answers should return 0, but got %d", score)
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
