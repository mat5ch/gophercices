package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
	"unicode/utf8"
)

const quizPath = "problems.csv"
const timerLimit = 30

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	var path = flag.String("csv", quizPath, "a csv file in the format of 'question,answer'")
	var limit = flag.Int("limit", timerLimit, "the time limit for the quiz in seconds")
	flag.Parse()

	data, err := os.ReadFile(*path)
	check(err)

	br := bytes.NewReader(data)
	csvReader := csv.NewReader(br)

	records, err := csvReader.ReadAll()
	check(err)

	totalScore := 0
	idx := 0

	/* Ask for user input to start the timer */
	fmt.Println("Press 'Enter' to start the quiz.")
	reader := bufio.NewReader(os.Stdin)
	/* Block further program execution until the user presses 'Enter' */
	for {
		b, err := reader.ReadBytes('\n')
		check(err)
		r, _ := utf8.DecodeRune(b)
		if int(r) == 10 {
			break
		} else {
			fmt.Println("Wrong input. Press 'Enter' to start the quiz.")
		}
	}

	/* ----- Timer handling ----- */
	timer1 := time.NewTimer(time.Duration(*limit) * time.Second)
	go func() {
		<-timer1.C
		fmt.Printf("\nTime exired.")
		fmt.Printf("\nYou scored %d out of %d.", totalScore, len(records))
		os.Exit(0)
	}()

	for {
		if idx < len(records) {
			fmt.Printf("Problem #%d: %s = ", idx+1, records[idx][0])
			input, err := reader.ReadString('\n')
			check(err)

			if strings.Trim(input, "\n") == records[idx][1] {
				totalScore++
			}
			idx++
		} else {
			break
		}

	}

	fmt.Printf("You scored %d out of %d.", totalScore, len(records))
	os.Exit(0)
}
