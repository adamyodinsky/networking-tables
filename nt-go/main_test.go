// main_test.go

package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestMainFlow(t *testing.T) {
	// Simulate user input
	input := bytes.NewBufferString("Alice Bob Charlie\n2 1\n3\n")
	peopleArr, tablesCapacity, combinationsLength := getUserInput(input)

	// Check the returned values
	if !strings.EqualFold(strings.Join(peopleArr, " "), "Alice Bob Charlie") {
		t.Errorf("Expected peopleArr to be 'Alice Bob Charlie', got '%v'", strings.Join(peopleArr, " "))
	}

	if !equal(tablesCapacity, []int{2, 1}) {
		t.Errorf("Expected tablesCapacity to be [2 1], got '%v'", tablesCapacity)
	}

	if combinationsLength != 3 {
		t.Errorf("Expected combinationsLength to be 3, got '%v'", combinationsLength)
	}

	// Run the algorithm
	peopleScores := initPeopleScores(peopleArr)
	run(combinationsLength, peopleArr, peopleScores, tablesCapacity)
}

func TestIntegration(t *testing.T) {
	// Generate peopleArr dynamically
	peopleArr := make([]string, 26)
	for i := 0; i < 26; i++ {
		peopleArr[i] = string(rune('a' + i))
	}

	// Hard-code tablesCapacity and combinationsLength values
	tablesCapacity := []int{4, 4, 4, 4, 4, 3, 3}
	combinationsLength := 10000

	// Initialize peopleScores
	peopleScores := initPeopleScores(peopleArr)

	// Run the algorithm
	run(combinationsLength, peopleArr, peopleScores, tablesCapacity)
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
