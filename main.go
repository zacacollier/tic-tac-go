package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	grid [][]string
	turn string
	win  bool
}

type Win string

const (
	PORT         = 8080
	Row      Win = "Row"
	Column   Win = "Column"
	Diagonal Win = "Diagonal"
	None     Win = "None"
)

func MakeBoard() Game {
	b := Game{
		grid: [][]string{
			{"_", "_", "_"},
			{"_", "_", "_"},
			{"_", "_", "_"},
		},
		turn: "X",
	}
	return b
}

func (b *Game) printBoard() {
	for _, row := range b.grid {
		fmt.Println(row)
	}
	fmt.Printf("***** Current turn: %v *****\n", b.turn)
}

func (b *Game) MakeMove(x, y int) {
	b.grid[x][y] = b.turn
}

func (b *Game) toggleTurn() {
	if b.turn == "X" {
		b.turn = "O"
	} else {
		b.turn = "X"
	}
}

func CompareCells(slice []string, comparator string) bool {
	var fst, snd, third string
	UnpackSlice(slice, &fst, &snd, &third)
	return fst == snd && snd == third && fst == comparator && snd == comparator && third == comparator
}

func UnpackSlice(s []string, vars ...*string) {
	for i, str := range s {
		*vars[i] = str
	}
}

func (b *Game) CheckForWin() (Win, int, bool) {
	columns := make([][]string, len(b.grid))

	// check for horizontal wins
	for i, row := range b.grid {
		if CompareCells(row, b.turn) {
			fmt.Println(i, true)
			b.win = true
			return Row, i, true
		}
		for j := range row {
			columns[j] = append(columns[j], row[j])
		}
	}

	// check for vertical wins
	for i, col := range columns {
		if CompareCells(col, b.turn) {
			fmt.Println(i, true)
			b.win = true
			return Column, i, true
		}
	}

	// check for diagonal wins
	diags := make([][]string, 2)
	diags[0] = append(diags[0], b.grid[0][0], b.grid[1][1], b.grid[2][2])
	diags[1] = append(diags[1], b.grid[2][0], b.grid[1][1], b.grid[0][2])
	for i, diag := range diags {
		if CompareCells(diag, b.turn) {
			b.win = true
			if diags[0][2] == b.turn {
				fmt.Println(2, true)
				return Diagonal, 2, true
			}
			fmt.Println(i, true)
			return Diagonal, i, true
		}
	}

	return None, -1, false
}

func GetInput(s *bufio.Scanner) (error, int, int) {
	fmt.Print("[Row, Column]: ")
	var row, column string
	re := regexp.MustCompile("\\s")
	for s.Scan() {
		fmt.Println(s.Text())
		moveArgs := strings.Split(
			re.ReplaceAllString(s.Text(), ""),
			",",
		)
		UnpackSlice(moveArgs, &row, &column)
		rowi, err := strconv.ParseInt(row, 10, 0)
		if err != nil {
			return err, -1, -1
		}
		coli, err := strconv.ParseInt(column, 10, 0)
		if err != nil {
			return err, -1, -1
		}
		return nil, int(rowi), int(coli)
	}
	return nil, 0, 0
}

func (b *Game) printAndToggle() {
	b.toggleTurn()
	b.printBoard()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	b := MakeBoard()
	for b.win == false {
		b.printAndToggle()
		err, row, column := GetInput(scanner)
		b.MakeMove(row, column)
		b.CheckForWin()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if b.win {
			b.printBoard()
			fmt.Printf("Player %v wins!\n", b.turn)
			os.Exit(0)
		} else {
			b.printAndToggle()
			err, row, column = GetInput(scanner)
			b.MakeMove(row, column)
			b.CheckForWin()
		}
	}
}
