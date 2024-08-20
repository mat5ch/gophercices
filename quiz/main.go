package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

const quizPath = "problems.csv"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	var path = flag.String("csv", quizPath, "a csv file in the format of 'question,answer'")
	flag.Parse()

	data, err := os.ReadFile(*path)
	check(err)

	br := bytes.NewReader(data)
	csvReader := csv.NewReader(br)

	records, err := csvReader.ReadAll()
	check(err)

	totalScore := 0
	idx := 0

	for {
		if idx < len(records) {
			fmt.Printf("Problem #%d: %s = ", idx+1, records[idx][0])
			reader := bufio.NewReader(os.Stdin)
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
}
