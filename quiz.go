package main

import (
	"os"
	"flag"
	"encoding/csv"
	"fmt"
	"strings"
	"time"
)

func main(){
	csvfilename := flag.String("csv","problems.csv","CSV file which contains question and answer separated by comma")
	timeLimit := flag.Int("limit",30,"Duration of time limitation")
	flag.Parse()

	file,err := os.Open(*csvfilename)
	check(err)
	defer file.Close()

	r := csv.NewReader(file)
	lines,err := r.ReadAll()
	check(err)

	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d : %s\n", i+1, p.q)
		answerCh := make(chan string)
		go func(){
			var answer string
			fmt.Scanf("%s\n",&answer)
			answerCh <- answer
		}()
		select {
		case <- timer.C:
			fmt.Printf("You answered %d correct answers out of %d problems",correct, len(problems))
			return
		case answer := <- answerCh:
			if answer == p.a{
				fmt.Printf("Correct\n")
				correct++
			}
		}
	}
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func parseLines(lines [][]string) []problem{
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q : line[0],
			a : strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct{
	q string
	a string
}