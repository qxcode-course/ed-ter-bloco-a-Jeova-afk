package main

import (
	"bufio"
	"fmt"
	"os"
)

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	if len(board) == 0 {
		return
	} 

	rows := len(board)
	cols := len(board[0])

	visited := make([][]bool, rows)

	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	var dfs func(int, int)
	
	dfs = func(r int, c int) {
		if r < 0 || r >= rows || c < 0|| c >= cols {
			return 
		}
	
		if visited[r][c] || board[r][c] != 'O' {
			return
		}
	
		visited[r][c] = true
	
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	
	}

	for i := 0; i < rows; i++ {
		dfs(i, 0)
		dfs(i, cols-1)
	}

	for j := 0; j < cols; j++ {
		dfs(0, j)
		dfs(rows-1, j)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O' && !visited[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}


// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan() 
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
