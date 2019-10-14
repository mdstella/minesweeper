package main_test

import (
	"strconv"
	"testing"

	"github.com/mdstella/minesweeper/core/model"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"

	"github.com/mdstella/minesweeper/core/endpoint"
	"github.com/mdstella/minesweeper/core/service"
)

// NOTE: on tests I think is better to use hardcoded strings, instead of the project constants like "*" or "?" (for the mine and flag)
// becaus if the flag change by any reason the test should fail and we will be notified about that when running tests

//Test_New_Game validates that the 9X9 board that is retrieved to the client is empty and without revealed cells
func Test_New_Game(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	ep := endpoint.MakeNewGameEndpoint(service.NewMinesweeperService())
	resp, err := ep(ctx, model.NewGameRequest{})
	assert.Nil(err)
	assert.NotNil(resp)
	rs := resp.(model.NewGameResponse)

	// response assertions, gameId not empty and 9X9 for the board size
	assert.NotEmpty(rs.GameId)
	assert.Nil(rs.Err)
	assert.Len(rs.Board, 9)
	for _, row := range rs.Board {
		assert.Len(row, 9)
		for _, cell := range row {
			assert.Empty(cell)
		}
	}
}

//Test_Flag_Cell flags a cell selected by the user
func Test_Flag_Cell(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	srv := service.NewMinesweeperService()

	ep := endpoint.MakeNewGameEndpoint(srv)
	resp, err := ep(ctx, model.NewGameRequest{})
	assert.Nil(err)
	assert.NotNil(resp)
	rs := resp.(model.NewGameResponse)
	assert.NotEmpty(rs.GameId)

	ep = endpoint.MakeAddFlagEndpoint(srv)
	resp, err = ep(ctx, model.PickCellRequest{
		Column: 0,
		Row:    0,
		GameId: rs.GameId,
	})

	rs2 := resp.(model.PickCellResponse)
	// response assertions, gameId not empty and 9X9 for the board size
	assert.NotEmpty(rs2.GameId)
	assert.Equal(rs.GameId, rs2.GameId)
	assert.Nil(rs2.Err)
	assert.Len(rs2.Board, 9)
	for i, row := range rs2.Board {
		assert.Len(row, 9)
		for j, cell := range row {
			if i == 0 && j == 0 {
				// Cell 0,0 should have the flag
				assert.NotEmpty(cell)
				assert.Equal("?", rs2.Board[0][0])
			} else {
				// rest of the cells should be empty
				assert.Empty(cell)
			}
		}
	}

}

//Test_Pick_Cell flags a cell selected by the user
func Test_Pick_Cell(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	srv := service.NewMinesweeperService()

	ep := endpoint.MakeNewGameEndpoint(srv)
	resp, err := ep(ctx, model.NewGameRequest{})
	assert.Nil(err)
	assert.NotNil(resp)
	rs := resp.(model.NewGameResponse)
	assert.NotEmpty(rs.GameId)

	ep = endpoint.MakePickCellEndpoint(srv)
	resp, err = ep(ctx, model.PickCellRequest{
		Column: 0,
		Row:    0,
		GameId: rs.GameId,
	})

	rs2 := resp.(model.PickCellResponse)
	assert.NotEmpty(rs2.GameId)
	// response assertions, gameId not empty and 9X9 for the board size
	assert.Equal(rs.GameId, rs2.GameId)
	assert.Nil(rs2.Err)
	assert.Len(rs2.Board, 9)
	for _, row := range rs2.Board {
		assert.Len(row, 9)
	}
	// cell 0,0 shouldn't be flaged
	assert.NotEmpty(rs2.Board[0][0])
	assert.NotEqual("?", rs2.Board[0][0])
	if rs2.EndedGame {
		// if game ended cell 0,0 should have a mine
		assert.Equal("*", rs2.Board[0][0])
	} else {
		// if game didn't end cell 0,0 should have a value between 0 to 3
		cellVal, err := strconv.Atoi(rs2.Board[0][0])
		assert.Nil(err)
		assert.True(cellVal >= 0 && cellVal <= 3)
	}

}
