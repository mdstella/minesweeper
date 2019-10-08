package service

import (
	"fmt"
)

// MinesweeperService is a simple interface for the minesweeper business logic.
//go:generate mockery -name=MinesweeperService
type MinesweeperService interface {

	// Test is a initial testing method
	Skeleton(string) (string, error)
}

type MinesweeperSrvImpl struct {
	// TODO add config here
}

// Test is a initial testing method
func (msi *MinesweeperSrvImpl) Skeleton(param string) (string, error) {
	return fmt.Sprintf("This is the skeleton service, parameter: %s", param), nil
}
