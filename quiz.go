package main

import (
	"os"
	"flag"
	"encoding/csv"
	"fmt"
	"strings"
)

func main(){
	csvfilename := flag.String("csv","problems.csv","CSV file which contains question and answer separated by comma")
	flag.Parse()

	file,err := os.Open(*csvfilename)
	check(err)
	defer file.Close()

	r := csv.NewReader(file)
	lines,err := r.ReadAll()
	check(err)

	problems := parseLines(lines)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d : %s\n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n",&answer)
		if answer == p.a{
			fmt.Printf("Correct\n")
			correct++
		}
	}

	fmt.Printf("You answered %d correct answers out of %d problems",correct, len(problems))
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