package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	f, err := os.Open("assets/day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var counter [][]int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		binaryStr := scanner.Text()

		if counter == nil {
			for range []rune(binaryStr) {
				counter = append(counter, make([]int, 2))
			}
		}

		for i, s := range []rune(binaryStr) {
			switch s {
			case '0':
				counter[i][0] += 1
			case '1':
				counter[i][1] += 1
			default:
				log.Fatalf("invalid binary value: %v", s)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	g, e := make([]rune, len(counter)), make([]rune, len(counter))
	for i, counts := range counter {
		if counts[0] > counts[1] {
			g[i] = '0'
			e[i] = '1'
		} else {
			g[i] = '1'
			e[i] = '0'
		}
	}

	gamma := binaryToInt(g)
	epsilon := binaryToInt(e)

	fmt.Printf("gamma: %d, epsilon: %d\npower: %d\n", gamma, epsilon, gamma*epsilon)
}

func binaryToInt(binary []rune) int {
	ret := 0

	for i, s := range binary {
		if s == '1' {
			ret += int(math.Pow(2, float64(len(binary)-1-i)))
		}
	}

	return ret
}
