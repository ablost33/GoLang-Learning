package algorithms

import "fmt"

var moves = []rune{'X', 'O', ' '}

func BuildGames(position int, board [][]rune, results *[][][]rune) {
	row, col := position/3, position%3
	// board is full
	if position == len(board)*len(board) {
		if isWinningBoard(board) {
			boardCopy := make([][]rune, len(board))
			for i := range boardCopy {
				boardCopy[i] = make([]rune, len(board[0]))
				copy(boardCopy[i], board[i])
			}
			*results = append(*results, boardCopy)
		}
		return
	}

	for _, move := range moves {
		board[row][col] = move
		BuildGames(position+1, board, results)
	}
}

func isWinningBoard(board [][]rune) bool {
	for i := 0; i < len(board); i++ {
		for j := 1; j < len(board); j++ {
			if board[i][j] != board[i][0] || board[i][j] == ' ' {
				return false
			}
			if board[j][i] != board[0][i] || board[j][i] == ' ' {
				return false
			}
		}
	}
	for i := 1; i < len(board); i++ {
		if board[i][i] != board[0][0] || board[i][i] == ' ' {
			return false
		}
		if board[0][len(board)-1] != board[i][len(board)-1-i] || board[i][len(board)-1-i] == ' ' {
			return false
		}
	}
	return true
}

func main() {
	size := 3 // Example for a variable board size

	// Initialize the board
	board := make([][]rune, size)
	for i := range board {
		board[i] = make([]rune, size)
	}

	// Container for all winning boards
	var results [][][]rune

	// Generate all board configurations and check for wins
	BuildGames(0, board, &results)

	// Print all winning configurations
	fmt.Printf("Found %d winning configurations:\n", len(results))
	for _, board := range results {
		for _, row := range board {
			for _, cell := range row {
				fmt.Printf("%c ", cell)
			}
			fmt.Println()
		}
		fmt.Println("----------")
	}
}
