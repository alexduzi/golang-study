package main

import (
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	var (
		answer      string
		answerIdx   int = 0
		userAnswers     = make(map[string]string)
		timer       *int
	)

	timer = flag.Int("timer", 30, "Timer to answer the quizz, default is 30 seconds")

	flag.Parse()

	correctAnswers := loadQuestions()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(*timer))
	defer cancel()

	go func() {
		<-ctx.Done()
		fmt.Println("\nTimes up!")
		printQuestionsSummary(userAnswers, correctAnswers)
		os.Exit(0)
	}()

	for question := range correctAnswers {
		fmt.Printf("Question %v: %v \n", (answerIdx + 1), question)
		fmt.Scanln(&answer)
		answer = strings.TrimSpace(answer)
		if answer != "" {
			userAnswers[question] = answer
		}
		answerIdx++

		if len(userAnswers) == len(correctAnswers) {
			printQuestionsSummary(userAnswers, correctAnswers)
			os.Exit(1)
		}
	}

}

func loadQuestions() map[string]string {
	correctAnswers := make(map[string]string)

	content, _ := os.ReadFile("./problems.csv")

	r := csv.NewReader(strings.NewReader(string(content)))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		correctAnswers[record[0]] = record[1]
	}

	return correctAnswers
}

func printQuestionsSummary(userAnswers map[string]string, correctAnswers map[string]string) {
	correctTotal := 0

	for idx, value := range correctAnswers {
		if userAnswers[idx] == value {
			correctTotal++
		}
	}

	fmt.Printf("You answered %v of %v \n", correctTotal, len(correctAnswers))
}
