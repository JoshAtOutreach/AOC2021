package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

func main() {
	f, err := os.Open("assets/day1/input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	winSize := 3
	previous := []int{}
	current := []int{}

	increases := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		d, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(errors.Wrapf(err, "Could not convert: %s", scanner.Text()))
		}
		if len(previous) != winSize {
			current = append(previous, d)
		} else {
			current = append(previous[1:], d)
			curSum := sum(current)
			prevSum := sum(previous)
			if curSum > prevSum {
				increases += 1
			}
		}

		previous = current
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("increases: %d\n", increases)
}

func sum(vals []int) int {
	sum := 0
	for _, val := range vals {
		sum += val
	}
	return sum
}
