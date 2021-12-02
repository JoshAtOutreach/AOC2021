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
	f, err := os.Open("assets/day1/input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	previous := -1
	increases := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(errors.Wrapf(err, "Could not convert: %s", scanner.Text()))
		}
		if previous != -1 && depth > previous {
			increases += 1
		}
		previous = depth
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("increases: %d\n", increases)
}
