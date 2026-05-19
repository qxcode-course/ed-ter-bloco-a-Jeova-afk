package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l int
	c int
}
	
func search(grid [][]rune, p Pos, end Pos) bool {
	if p == end {
		return true
	}

	dir := []Pos{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},

	}

	for _, d := range dir {
		nl := p.l + d.l
		nc := p.c + d.c

		if nl < 0|| nl >= len(grid) || nc < 0 || nc > len(grid[0]) {
			continue
		}

		if grid[nl][nc] == ' ' {
			grid[nl][nc] = '.'

			if search(grid, Pos{nl, nc}, end) {
				return true
			}

			grid [nl][nc] = ' '
		}

	}

	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nl_nc := scanner.Text()
	var nl, nc int
	fmt.Sscanf(nl_nc, "%d %d", &nl, &nc)
	grid := make([][]rune, nl)

	// Lê a gridriz
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []rune(scanner.Text())
	}

	// Procura posições de início e endPos e conserta para _
	var startPos, endPos Pos
	for i := 0; i < nl; i++ {
		for c := 0; c < nc; c++ {
			if grid[i][c] == 'I' {
				grid[i][c] = '.'
				startPos = Pos{i, c}
			}
			if grid[i][c] == 'F' {
				grid[i][c] = ' '
				endPos = Pos{i, c}
			}
		}
	}

	search(grid, startPos, endPos)

	// Imprime o labirinto final
	for _, line := range grid {
		fmt.Println(string(line))
	}
}
