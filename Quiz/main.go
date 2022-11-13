package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "csv file")
	timeLimit := flag.Int("limit", 5, "the time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFileName)

	if err != nil {
		fmt.Printf("Fail to open the file.")
	}

	r := csv.NewReader(file)
	
	lines, err := r.ReadAll()

	if err != nil {
		fmt.Printf("Fail to parse the file.")
	}

	problems := parseFile(lines)

	time := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	count := 0

	for i, p := range problems {
		fmt.Printf("Problem %d: %s= ", i+1, p.question)
		answerChan := make(chan string)

		go func ()  {
			var answer string 
			fmt.Scanf("%s\n", &answer)
			answerChan <- answer			
		}()

		select {
			case <- time.C:
				fmt.Printf("You scored %d out of %d", count, len(problems))
				return 
			case answer := <- answerChan:
				if answer == p.answer {
					count++
				}
		}
	}
	
	fmt.Printf("You scored %d out of %d", count, len(problems))
}

func parseFile(lines [][]string) []problems {
	ret := make([]problems, len(lines)) 

	for i, line := range lines {
		ret[i] = problems{
			question: line[0],
			answer: strings.TrimSpace(line[1]),
		}
	}

	return ret
}

type problems struct {
	question string
	answer string
}