package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var fileName = flag.String("file_name", "problems.csv", "choose quiz file")
var timer = flag.Int("timer", 2, "set the time for the quiz")

func main() {
	flag.Parse()
	csvFile, err := os.Open(*fileName)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer csvFile.Close()
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Printf("time: %v\tpress Enter to start...\n", *timer)
	consoleReader.ReadString('\n')
	doneChan := make(chan bool)
	go func() {
		time.Sleep(time.Duration(*timer) * time.Second)
		doneChan <- true
	}()

	var correctAnswersCount int

	go func() {
		for _, v := range csvLines {
			fmt.Printf("%v", v[0])
			answer, _ := consoleReader.ReadString('\n')

			if strings.TrimRight(answer, "\n") == v[1] {
				correctAnswersCount++
			}
		}

		doneChan <- true
	}()

	<-doneChan
	fmt.Printf("\nquestions: %v, correct answers: %v\n", len(csvLines), correctAnswersCount)
}
