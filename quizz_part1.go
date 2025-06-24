package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	content, _ := os.ReadFile("./problems.csv")

	reader := csv.NewReader(strings.NewReader(string(content)))

	answer := 0
	answerIdx := 0
	userAnswers := make(map[int]int)
	correctAnswers := make(map[int]int)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Question %v: %v \n", (answerIdx + 1), record[0])
		fmt.Scanln(&answer)
		userAnswers[answerIdx] = answer
		value, _ := strconv.Atoi(record[1])
		correctAnswers[answerIdx] = value
		answerIdx++
	}

	correctTotal := 0

	for idx, value := range correctAnswers {
		if userAnswers[idx] == value {
			correctTotal++
		}
	}

	fmt.Printf("You answered %v of %v \n", correctTotal, len(correctAnswers))
}
