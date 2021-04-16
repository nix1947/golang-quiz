package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	fName := flag.String("f", "default.csv", "path of csv file")
	flag.Parse()

	fObj, err := os.Open(*fName)
	if err != nil {
		exit(fmt.Sprintf("error occured in open %s", *fName))
	}
	csvR := csv.NewReader(fObj)
	cLines, err := csvR.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("error in reading all line from %s file", *fName))
	}
	problems := parseProblem(cLines)
	fmt.Println(problems)

}

func parseProblem(lines [][]string) []problem {
	r := make([]problem, len(lines))
	for i := 0; i < len(lines); i++ {
		r[i] = problem{q: lines[i][0], a: lines[i][1]}
	}
	return r
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)

}
