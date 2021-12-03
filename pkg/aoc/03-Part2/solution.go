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

	report := [][]rune{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		report = append(report, []rune(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	gList, eList := report, report

	log.Println("gamma list")
	for i := 0; len(gList) > 1; i++ {
		log.Printf("initial list: \n%v\n", gList)
		r := findCommonRuneAtIndex(i, gList, false)
		log.Printf("most common value at index %d is: %v\n", i, r)
		gList = filterByRuneAtIndex(i, gList, r)
		log.Printf("filtered list: \n%v\n", gList)
	}

	log.Println("\nepsilon list")
	for i := 0; len(eList) > 1; i++ {
		log.Printf("initial list: \n%v\n", eList)
		r := findCommonRuneAtIndex(i, eList, true)
		log.Printf("least common value at index %d is: %v\n", i, r)
		eList = filterByRuneAtIndex(i, eList, r)
		log.Printf("filtered list: \n%v\n", eList)
	}

	gamma := binaryToInt(gList[0])
	epsilon := binaryToInt(eList[0])

	fmt.Printf("gamma: %d, epsilon: %d\npower: %d\n", gamma, epsilon, gamma*epsilon)
}

func findCommonRuneAtIndex(i int, list [][]rune, least bool) rune {
	count := make([]int, 2)
	for _, binary := range list {
		if binary[i] == '0' {
			count[0]++
		} else {
			count[1]++
		}
	}
	if least {
		if count[0] < count[1] || count[0] == count[1] {
			return '0'
		}
		return '1'
	}
	if count[0] > count[1] {
		return '0'
	}
	return '1'
}

func filterByRuneAtIndex(i int, list [][]rune, r rune) [][]rune {
	var ret [][]rune
	for _, binary := range list {
		if binary[i] == r {
			ret = append(ret, binary)
		}
	}
	return ret
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
