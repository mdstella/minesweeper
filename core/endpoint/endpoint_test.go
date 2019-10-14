package endpoint_test

import (
	er "errors"
	"testing"

	"golang.org/x/net/context"

	"github.com/mdstella/minesweeper/core/endpoint"

	"github.com/mdstella/minesweeper/core/model"
	"github.com/mdstella/minesweeper/core/service/mocks"

	"github.com/stretchr/testify/assert"
)

// Unit tests for the endpoint layer, will mock service layer as it unit test

func getEmptyBoard() [][]string {
	userBoard := make([][]string, 9)

	// first stage of board generation, empty board with
	for i := 0; i < 9; i++ {
		userBoard[i] = make([]string, 9)
	}
	return userBoard
}

func Test_Make_New_Game_OK(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	srv_m := &mocks.MinesweeperService{}
	gd := model.GameDefintion{
		GameId: "gameId",
		Board:  getEmptyBoard(),
	}
	// Mocks response when NewGame service is invoked
	srv_m.On("NewGame").Return(gd, nil)

	ep := endpoint.MakeNewGameEndpoint(srv_m)
	//validates results
	resp, err := ep(ctx, model.NewGameRequest{})
	assert.Nil(err)
	rs := resp.(model.NewGameResponse)
	assert.Nil(rs.Err)
	assert.Equal("gameId", rs.GameId)
	assert.Len(rs.Board, 9)
	for _, row := range rs.Board {
		assert.Len(row, 9)
		for _, cell := range row {
			assert.Empty(cell)
		}
	}

	// assert expectations over mock, this means that we are validating the
	// mocks defined on the test over srv_m were effectively invoked
	srv_m.AssertExpectations(t)
}

func Test_Make_New_Game_Error(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	srv_m := &mocks.MinesweeperService{}

	// Mocks response when NewGame service is invoked
	srv_m.On("NewGame").Return(model.GameDefintion{}, er.New("SOME FAKE ERROR"))

	ep := endpoint.MakeNewGameEndpoint(srv_m)
	//validates results
	resp, err := ep(ctx, model.NewGameRequest{})
	assert.Nil(err)
	rs := resp.(model.NewGameResponse)
	assert.NotNil(rs.Err)

	// assert expectations over mock, this means that we are validating the
	// mocks defined on the test over srv_m were effectively invoked
	srv_m.AssertExpectations(t)
}

func Test_Make_PickCell_OK(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	srv_m := &mocks.MinesweeperService{}
	board := getEmptyBoard()
	board[0][0] = "*"
	gd := model.GameDefintion{
		GameId:    "gameId",
		Board:     board,
		Won:       false,
		EndedGame: true,
	}
	// Mocks response when PickCell service is invoked
	srv_m.On("PickCell", "gameId", 0, 0).Return(gd, nil)

	ep := endpoint.MakePickCellEndpoint(srv_m)
	//validates results
	resp, err := ep(ctx, model.PickCellRequest{
		Column: 0,
		Row:    0,
		GameId: "gameId",
	})
	assert.Nil(err)
	rs := resp.(model.PickCellResponse)
	assert.Nil(rs.Err)
	assert.Equal("gameId", rs.GameId)
	assert.Len(rs.Board, 9)
	for i, row := range rs.Board {
		assert.Len(row, 9)
		for j, cell := range row {
			if i == 0 && j == 0 {
				assert.Equal("*", cell)
			} else {
				assert.Empty(cell)
			}
		}
	}

	// assert expectations over mock, this means that we are validating the
	// mocks defined on the test over srv_m were effectively invoked
	srv_m.AssertExpectations(t)
}

func Test_Make_Pick_Cell_Error(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	srv_m := &mocks.MinesweeperService{}

	// Mocks response when PickCell service is invoked
	srv_m.On("PickCell", "gameId", 0, 0).Return(model.GameDefintion{}, er.New("SOME FAKE ERROR"))

	ep := endpoint.MakePickCellEndpoint(srv_m)
	//validates results
	resp, err := ep(ctx, model.PickCellRequest{
		Column: 0,
		Row:    0,
		GameId: "gameId",
	})
	assert.Nil(err)
	rs := resp.(model.PickCellResponse)
	assert.NotNil(rs.Err)

	// assert expectations over mock, this means that we are validating the
	// mocks defined on the test over srv_m were effectively invoked
	srv_m.AssertExpectations(t)
}

func Test_Make_AddFlag_OK(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	srv_m := &mocks.MinesweeperService{}
	board := getEmptyBoard()
	board[0][0] = "?"
	gd := model.GameDefintion{
		GameId:    "gameId",
		Board:     board,
		Won:       false,
		EndedGame: true,
	}
	// Mocks response when AddFlag service is invoked
	srv_m.On("AddFlag", "gameId", 0, 0).Return(gd, nil)

	ep := endpoint.MakeAddFlagEndpoint(srv_m)
	//validates results
	resp, err := ep(ctx, model.PickCellRequest{
		Column: 0,
		Row:    0,
		GameId: "gameId",
	})
	assert.Nil(err)
	rs := resp.(model.PickCellResponse)
	assert.Nil(rs.Err)
	assert.Equal("gameId", rs.GameId)
	assert.Len(rs.Board, 9)
	for i, row := range rs.Board {
		assert.Len(row, 9)
		for j, cell := range row {
			if i == 0 && j == 0 {
				assert.Equal("?", cell)
			} else {
				assert.Empty(cell)
			}
		}
	}

	// assert expectations over mock, this means that we are validating the
	// mocks defined on the test over srv_m were effectively invoked
	srv_m.AssertExpectations(t)
}

func Test_Make_AddFlag_Error(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	srv_m := &mocks.MinesweeperService{}

	// Mocks response when AddFlag service is invoked
	srv_m.On("AddFlag", "gameId", 0, 0).Return(model.GameDefintion{}, er.New("SOME FAKE ERROR"))

	ep := endpoint.MakeAddFlagEndpoint(srv_m)
	//validates results
	resp, err := ep(ctx, model.PickCellRequest{
		Column: 0,
		Row:    0,
		GameId: "gameId",
	})
	assert.Nil(err)
	rs := resp.(model.PickCellResponse)
	assert.NotNil(rs.Err)

	// assert expectations over mock, this means that we are validating the
	// mocks defined on the test over srv_m were effectively invoked
	srv_m.AssertExpectations(t)
}
