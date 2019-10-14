package service

import (
	"fmt"
	"strings"

	"github.com/mdstella/minesweeper/core/errors"
	"github.com/mdstella/minesweeper/core/model"
)

// This file will have common functions for the services

// validateInputs checks if the inputs sent by the client are valid
func (msi *MinesweeperSrvImpl) validateInputs(gameId string, row, column int) error {
	gameId = strings.TrimSpace(gameId)
	if gameId == "" {
		return errors.NewBadParamError("GameId is empty")
	}

	if row < 0 || row >= ROWS {
		return errors.NewBadParamError(fmt.Sprintf("Wrong row value should be between 0 and %d", ROWS))
	}

	if column < 0 || column >= COLS {
		return errors.NewBadParamError(fmt.Sprintf("Wrong column value should be between 0 and %d", COLS))
	}
	return nil
}

// getBoardFromLru retrieve the board from the cache by gameId
func (msi *MinesweeperSrvImpl) getBoardFromLru(gameId string) (model.Board, error) {
	msi.mutex.RLock()
	boardIntf, ok := msi.games.Get(gameId)
	if !ok {
		msi.mutex.RUnlock()
		return model.Board{}, errors.NewBadParamError(fmt.Sprintf("Error trying to obtain game by id %s. Please start a new game", gameId))
	}
	msi.mutex.RUnlock()
	return boardIntf.(model.Board), nil
}
