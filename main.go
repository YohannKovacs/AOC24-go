package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/yohannKovacs/aoc24/solution"
)

func main() {
	f, err := os.Open("./inputs/D1.txt")
	if err != nil {
		log.Fatalf("Error while file open : %v", err)
	}
	var buf []byte
	buf, err = io.ReadAll(f)
	if err != nil {
		log.Fatalf("Error while reading from file : %v", err)
	}

	p1, err1, p2, err2 := solution.SolutionDay1(strings.Trim(string(buf), "\n"))
	if err1 != nil && err2 != nil {
		log.Fatalf("Error while finding solution of part 1 and part 2: %v %v", err1, err2)
	}
	if err1 != nil {
		log.Fatalf("Error while finding solution of part 1: %v", err2)
	}
	if err2 != nil {
		log.Fatalf("Error while finding solution of part 2: %v", err2)
	}

	fmt.Printf("Answer\np1 : %v\np2 : %v\n", p1, p2)
}
