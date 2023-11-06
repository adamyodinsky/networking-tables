// main.go

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/schollz/progressbar/v3"
)

func main() {
	// Getting user input
	peopleArr, tablesCapacity, combinationsLength := getUserInput(os.Stdin)

	// Initializing people scores
	peopleScores := initPeopleScores(peopleArr)

	// Running the algorithm
	runningAlgo(
		combinationsLength,
		peopleArr,
		peopleScores,
		tablesCapacity,
	)
}

func getUserInput(reader io.Reader) ([]string, []int, int) {
	scanner := bufio.NewScanner(reader)

	fmt.Println("Enter the names of the people: ")
	scanner.Scan()
	peopleArr := strings.Split(scanner.Text(), " ")

	fmt.Println("Enter the capacity of the tables: ")
	scanner.Scan()
	tablesCapacity := splitToInt(scanner.Text(), " ")

	fmt.Println("Enter the number of brute force combinations: ")
	scanner.Scan()
	combinationsLength, _ := strconv.Atoi(scanner.Text())

	return peopleArr, tablesCapacity, combinationsLength
}

func splitToInt(s, delim string) []int {
	ss := strings.Split(s, delim)
	ii := make([]int, len(ss))
	for i, v := range ss {
		ii[i], _ = strconv.Atoi(v)
	}
	return ii
}

func initPeopleScores(peopleArr []string) map[string]map[string]int {
	peopleScores := make(map[string]map[string]int)
	for _, person := range peopleArr {
		peopleScores[person] = make(map[string]int)
		for _, otherPerson := range peopleArr {
			if person != otherPerson {
				peopleScores[person][otherPerson] = 0
			}
		}
	}
	return peopleScores
}

func bestCombination(
	combinationsLength int,
	peopleArr []string,
	peopleScores map[string]map[string]int,
	tablesCapacity []int,
) map[string]interface{} {

	bar := progressbar.Default(int64(combinationsLength))

	bestCombination := make(map[string]interface{})
	bestCombination["score"] = math.MaxInt64

	for i := 0; i < combinationsLength; i++ {
		bar.Add(1)
		tables := []map[string]interface{}{}
		peopleArrCopy := make([]string, len(peopleArr))
		copy(peopleArrCopy, peopleArr)
		peopleScoresCopy := deepCopy(peopleScores)

		for _, tableLength := range tablesCapacity {
			table := []string{}
			tableScore := 0

			for j := 0; j < tableLength; j++ {
				index := rand.Intn(len(peopleArrCopy))
				table = append(table, peopleArrCopy[index])
				peopleArrCopy = append(peopleArrCopy[:index], peopleArrCopy[index+1:]...)
			}

			for _, person := range table {
				for _, otherPerson := range table {
					if person != otherPerson {
						if peopleScoresCopy[person][otherPerson] > 0 {
							peopleScoresCopy[person][otherPerson] += 10
						} else {
							peopleScoresCopy[person][otherPerson] += 1
						}
						tableScore += peopleScoresCopy[person][otherPerson]
					}
				}
			}

			tables = append(tables, map[string]interface{}{
				"table": table,
				"score": tableScore,
			})
		}

		combinationScore := 0
		for _, table := range tables {
			combinationScore += table["score"].(int)
		}

		if combinationScore < bestCombination["score"].(int) {
			bestCombination = map[string]interface{}{
				"tables":        tables,
				"score":         combinationScore,
				"people_scores": peopleScoresCopy,
			}
		}
	}

	return bestCombination
}

func deepCopy(m map[string]map[string]int) map[string]map[string]int {
	copy := make(map[string]map[string]int)
	for k, v := range m {
		copy[k] = make(map[string]int)
		for k2, v2 := range v {
			copy[k][k2] = v2
		}
	}
	return copy
}

func runningAlgo(
	combinationsLength int,
	peopleArr []string,
	peopleScores map[string]map[string]int,
	tablesCapacity []int,
) {
	iterations := len(tablesCapacity)
	finalScore := 0

	for i := 0; i < iterations; i++ {

		fmt.Println("Iteration", i+1)
		combination := bestCombination(
			combinationsLength,
			peopleArr,
			peopleScores,
			tablesCapacity,
		)
		peopleScores = combination["people_scores"].(map[string]map[string]int)
		fmt.Println("Score:", combination["score"])
		tablesJSON, _ := json.MarshalIndent(combination["tables"], "", "    ")
		fmt.Println("Tables:", string(tablesJSON))
		finalScore += combination["score"].(int)
	}

	fmt.Println("Final score:", finalScore)
}
