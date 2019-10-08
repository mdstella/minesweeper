package service

import (
	"errors"
	"math/rand"
	"strconv"
	"time"

	"github.com/lithammer/shortuuid"

	"github.com/mdstella/minesweeper/core/model"
)

const (
	// For now we are just hardcoding the board size and number of mines
	ROWS        = 9
	COLS        = 9
	MINES_COUNT = 10

	// Mine representation in the board
	MINE = "*"
)

// MinesweeperService is a simple interface for the minesweeper business logic.
//go:generate mockery -name=MinesweeperService
type MinesweeperService interface {

	// Test is a initial testing method
	NewGame() (model.GameDefintion, error)
}

type MinesweeperSrvImpl struct {
	// TODO add config here
}

// Test is a initial testing method
func (msi *MinesweeperSrvImpl) NewGame() (model.GameDefintion, error) {
	board, err := msi.generateBoard(ROWS, COLS, MINES_COUNT)
	if err != nil {
		return model.GameDefintion{}, err
	}
	return model.GameDefintion{
		GameId: shortuuid.New(),
		Board:  board,
	}, nil
}

// generateBoard - will generate a board with the configuration sent by parameter (rows, columns and mines)
// this method starts with lower case as it's not exposed and it can't be invoked outside the service layer.
func (msi *MinesweeperSrvImpl) generateBoard(rows, cols, mines int) ([][]string, error) {

	if rows == 0 || cols == 0 || rows*cols < mines {
		return nil, errors.New("Wrong configuration for the game, review the amount of rows, columns and mines")
	}

	board := make([][]string, rows)

	// first stage of board generation, empty board with
	for i := 0; i < rows; i++ {
		board[i] = make([]string, cols)
	}

	// Adding random mines on the board
	for i := 0; i < mines; i++ {
		row, column := msi.getRandomCell(rows, cols)
		if board[row][column] != "*" {
			board[row][column] = "*"
		} else {
			// decrement the index to avoid obtaining duplicated random cell
			// with this will be requested again
			i--
		}
	}

	// Counting the amount of mines next to each cell
	// for completing the board
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// if the cell doesn't contains a mine we will calculate the amount of mines it has
			// in closer cells
			if board[i][j] != "*" {
				board[i][j] = strconv.Itoa(msi.calculateCloseMines(board, i, j))
			}
		}
	}

	return board, nil
}

// getRandomCell - will retrieve a random row and column representing a cell in the board
func (msi *MinesweeperSrvImpl) getRandomCell(rows, cols int) (row, col int) {
	rand.Seed(time.Now().UnixNano())
	row = rand.Intn(rows)
	col = rand.Intn(cols)
	return
}

// calculateCloseMines will calculate the amount of mines the cell has around
// for example, the cell in the middle has 6 mines around it
// "*"   "*"   "*"
// " "   "6"   "*"
// "*"   " "   "*"
func (msi *MinesweeperSrvImpl) calculateCloseMines(board [][]string, row, column int) int {
	count := 0
	// calculate upper left
	count += msi.hasMine(board, row-1, column-1)
	// calculate upper
	count += msi.hasMine(board, row-1, column)
	// calculate upper right
	count += msi.hasMine(board, row-1, column+1)
	// calculate left
	count += msi.hasMine(board, row, column-1)
	// calculate right
	count += msi.hasMine(board, row, column+1)
	// calculate lower left
	count += msi.hasMine(board, row+1, column-1)
	// calculate lower
	count += msi.hasMine(board, row+1, column)
	// calculate lower right
	count += msi.hasMine(board, row+1, column+1)
	return count
}

func (msi *MinesweeperSrvImpl) hasMine(board [][]string, row, column int) int {
	if row >= 0 && row < len(board) {
		if column >= 0 && column < len(board[row]) && board[row][column] == MINE {
			return 1
		}
	}
	return 0
}
