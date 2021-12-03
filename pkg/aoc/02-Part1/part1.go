package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func main() {
	f, err := os.Open("assets/day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	x, y := 0, 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		if len(input) != 2 {
			log.Fatalf("input is invalid: %s", scanner.Text())
		}

		command, magnitudeStr := input[0], input[1]

		magnitude, err := strconv.Atoi(magnitudeStr)
		if err != nil {
			log.Fatal(errors.Wrapf(err, "could not extract magnitude from: %s", magnitudeStr))
		}

		switch command {
		case "forward":
			x += magnitude
		case "up":
			y -= magnitude
		case "down":
			y += magnitude
		default:
			log.Fatal("invalid command")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("x: %d, y: %d\nproduct: %d\n", x, y, x*y)
}
