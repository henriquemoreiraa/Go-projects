package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"

	"os"
	"time"
)

func main() {
	csvFile := flag.String("level", "easy.csv", "words level")
	timeLimit := flag.Float64("time", 15, "time limit")
	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	words := parseFile(lines)

	time := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	secInMin := *timeLimit / 60

	var correct float64
	var keysPressed float64
	var wpm int
	var accuracy int

	for _, word := range words {
		answerChan := make(chan string)
		fmt.Printf(word + " ")

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChan <- answer
		}()

		select {
		case <-time.C:
			if keysPressed != 0 && correct != 0 {
				wpm = int((keysPressed / 5) / secInMin)
				accuracy = int((correct / keysPressed) * 100)
			}

			fmt.Printf("\nYou type %d WPM\n", wpm)
			fmt.Printf("You got %g words right out of %d\n", correct, len(words))
			fmt.Printf("Your Accuracy is: %d%\n", accuracy)
			return
		case answer := <-answerChan:
			if len(answer) <= len(word) {
				keysPressed += float64(len(answer))
			}
			if answer == word {
				correct++
			}
		}
	}
	wpm = int((keysPressed / 5) / secInMin)
	accuracy = int((correct / keysPressed) * 100)

	fmt.Printf("Your WPM is: %d\n", wpm)
	fmt.Printf("You got %g words right out of %d\n", correct, len(words))
	fmt.Printf("Your Accuracy is: %d\n", accuracy)
}

func parseFile(l [][]string) []string {
	w := make([]string, len(l))

	for i, line := range l {
		w[i] = line[0]
	}

	return w
}
