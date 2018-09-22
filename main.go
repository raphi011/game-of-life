package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Board [][]int

const (
	TICK    = 500 * time.Millisecond
	ROWS    = 10
	COLUMNS = 10
)

func main() {
	b := newBoard(ROWS, COLUMNS)
	b.seed()
	b.run(15)
}

func newBoard(rows, columns int) Board {
	board := make([][]int, columns)

	for column := range board {
		board[column] = make([]int, rows)
	}

	return board
}

func (b Board) run(steps int) {
	clearScreen()
	b.print()

	for i := 0; i < steps; i++ {
		time.Sleep(TICK)

		clearScreen()

		b = b.step()
		b.print()
	}
}

func (b Board) dimensions() (rows, columns int) {
	rows = len(b)
	columns = len(b[0])

	return
}

func (b Board) neighbors(x, y int) int {
	rows, columns := b.dimensions()
	count := 0

	left := mod(x-1, columns)
	right := mod(x+1, columns)
	top := mod(y-1, rows)
	bottom := mod(y+1, rows)

	count += b[top][left]
	count += b[top][x]
	count += b[top][right]

	count += b[y][left]
	count += b[y][right]

	count += b[bottom][left]
	count += b[bottom][x]
	count += b[bottom][right]

	return count
}

func (b Board) seed() {
	for y := range b {
		for x := range b[y] {
			b[y][x] = int(math.Round(rand.Float64()))
		}
	}
}

func (b Board) step() Board {
	copy := b.copy()

	for y := range copy {
		for x, live := range copy[y] {
			neighbors := b.neighbors(x, y)

			copy[y][x] = aliveOrDead(neighbors, live)
		}
	}

	return copy
}

func aliveOrDead(neighbors, live int) int {
	if live > 0 {

		if neighbors < 2 {
			return 0
		} else if neighbors < 4 {
			return 1
		} else {
			return 0
		}

	} else if neighbors == 3 {
		return 1
	} else {
		return 0
	}
}

func (b Board) copy() Board {
	rows, columns := b.dimensions()

	copy := newBoard(rows, columns)

	for row := range b {
		for column, filled := range b[row] {
			copy[row][column] = filled
		}
	}

	return copy
}

func (b Board) print() {
	for row := range b {
		for _, filled := range b[row] {
			if filled > 0 {
				fmt.Print("x")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	fmt.Println()
}

func mod(a, b int) int {
	result := a % b

	if result < 0 {
		return result + b
	}

	return result
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
