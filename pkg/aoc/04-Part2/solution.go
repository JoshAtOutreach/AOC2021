package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("assets/day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	first := true
	calls := []int{}
	var board [][]int
	y := 0

	boards := []*Board{}

	for scanner.Scan() {
		switch {
		case first:
			for _, str := range strings.Split(scanner.Text(), ",") {
				val, err := strconv.Atoi(str)
				if err != nil {
					log.Fatalf("building calls: could not convert %s to an int\n", str)
				}
				calls = append(calls, val)
			}
			first = false
		case scanner.Text() != "":
			for x, str := range strings.Fields(scanner.Text()) {
				val, err := strconv.Atoi(str)
				if err != nil {
					log.Fatalf("building boards: could not convert %s to an int\n", str)
				}
				board[x][y] = val
			}
			y++
		default:
			log.Print("hit default case")
			if len(board) > 0 {
				log.Print("appending board")
				boards = append(boards, NewBoard(board))
			}
			board = make([][]int, 5)
			for x := range board {
				board[x] = make([]int, 5)
			}
			y = 0
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	boards = append(boards, NewBoard(board))
	last := &Board{}
	for rnd, call := range calls {
		log.Printf("Round %d:\n", rnd)
		filtered := []*Board{}
		for _, b := range boards {
			b.CheckAndMark(call)
			if !b.Bingo() {
				filtered = append(filtered, b)
			} else {
				last = b
			}
		}
		boards = filtered
		if len(boards) == 0 {
			log.Printf("Found the losing board. Last value called: %d board score: %d product: %d", call, last.Score(), last.Score()*call)
			os.Exit(0)
		}
	}
	log.Fatal("no losers?")
}

type Board struct {
	Squares [][]Square // 2D Array
}

func NewBoard(board [][]int) *Board {
	squares := [][]Square{}
	for _, col := range board {
		sqrs := []Square{}
		for _, val := range col {
			sqrs = append(sqrs, NewSquare(val))
		}
		squares = append(squares, sqrs)
	}
	return &Board{Squares: squares}
}

func (b *Board) Print() {
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			sqr := b.Squares[x][y]
			fmt.Printf("%2d", sqr.Value)
			if sqr.Called {
				fmt.Printf("! ")
			} else {
				fmt.Printf("  ")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func (b *Board) CheckAndMark(val int) {
	for x := range b.Squares {
		for y := range b.Squares[x] {
			if b.Squares[x][y].Value == val {
				b.Squares[x][y].Called = true
			}
		}
	}
}

func (b *Board) Bingo() bool {
	rowCounts := make([]int, 5)
	for y := range rowCounts {
		rowCounts[y] = 0
	}

	for _, col := range b.Squares {
		count := 0
		for y, sqr := range col {
			if sqr.Called {
				count++
				rowCounts[y]++
			}
		}
		if count == len(b.Squares) {
			return true
		}
	}
	for _, count := range rowCounts {
		if count == len(b.Squares) {
			return true
		}
	}
	return false
}

func (b *Board) Score() int {
	ret := 0
	for _, col := range b.Squares {
		for _, sqr := range col {
			if !sqr.Called {
				ret += sqr.Value
			}
		}
	}
	return ret
}

type Square struct {
	Value  int
	Called bool
}

func NewSquare(value int) Square {
	return Square{Value: value, Called: false}
}
