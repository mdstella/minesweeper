package service

import (
	"errors"

	"github.com/lithammer/shortuuid"

	"github.com/mdstella/minesweeper/core/model"
)

const (
	// For now we are just hardcoding the board size and number of mines
	ROWS  = 9
	COLS  = 9
	MINES = 10
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
	board, err := msi.generateBoard(ROWS, COLS, MINES)
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

	board := make([][]string, 0)

	// first stage of board generation, empty board with
	// every cell having value = 0
	for i := 0; i < rows; i++ {
		row := make([]string, 0)
		for j := 0; j < cols; j++ {
			row = append(row, "0")
		}
		board = append(board, row)
	}

	// TODO add mines and add counting mines for cells

	return board, nil
}
