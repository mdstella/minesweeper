package service_test

import (
	"strconv"
	"testing"

	"github.com/mdstella/minesweeper/core/errors"
	"github.com/mdstella/minesweeper/core/model"
	"github.com/mdstella/minesweeper/core/service"

	"github.com/stretchr/testify/assert"
)

// Unit tests for the service layer

//Test_New_Game_OK validates that the 9X9 board that is retrieved to the client is empty and without revealed cells
func Test_New_Game_OK(t *testing.T) {
	assert := assert.New(t)

	srv := service.NewMinesweeperService()
	gameDef, err := srv.NewGame()
	assert.Nil(err)
	// response assertions, gameId not empty and 9X9 for the board size
	assert.NotEmpty(gameDef.GameId)
	assert.False(gameDef.EndedGame)
	assert.False(gameDef.Won)
	assert.Len(gameDef.Board, 9)
	for _, row := range gameDef.Board {
		assert.Len(row, 9)
		for _, cell := range row {
			assert.Empty(cell)
		}
	}
}

//Test_Flag_Cell_Ok flags a cell selected by the user
func Test_Flag_Cell_OK(t *testing.T) {
	assert := assert.New(t)

	srv := service.NewMinesweeperService()
	gameDef, err := srv.NewGame()
	assert.Nil(err)
	// response assertions, gameId not empty and 9X9 for the board size
	assert.NotEmpty(gameDef.GameId)
	assert.False(gameDef.EndedGame)
	assert.False(gameDef.Won)
	assert.Len(gameDef.Board, 9)
	for _, row := range gameDef.Board {
		assert.Len(row, 9)
		for _, cell := range row {
			assert.Empty(cell)
		}
	}

	gd2, err := srv.AddFlag(gameDef.GameId, 0, 0)
	assert.Nil(err)
	// response assertions, gameId not empty and 9X9 for the board size
	assert.NotEmpty(gd2.GameId)
	assert.Equal(gameDef.GameId, gd2.GameId)
	assert.Len(gd2.Board, 9)
	for i, row := range gd2.Board {
		assert.Len(row, 9)
		for j, cell := range row {
			if i == 0 && j == 0 {
				// Cell 0,0 should have the flag
				assert.NotEmpty(cell)
				assert.Equal("?", gd2.Board[0][0])
			} else {
				// rest of the cells should be empty
				assert.Empty(cell)
			}
		}
	}

}

//Test_Pick_Cell_OK flags a cell selected by the user
func Test_Pick_Cell_OK(t *testing.T) {
	assert := assert.New(t)

	srv := service.NewMinesweeperService()
	gameDef, err := srv.NewGame()
	assert.Nil(err)
	// response assertions, gameId not empty and 9X9 for the board size
	assert.NotEmpty(gameDef.GameId)
	assert.False(gameDef.EndedGame)
	assert.False(gameDef.Won)
	assert.Len(gameDef.Board, 9)
	for _, row := range gameDef.Board {
		assert.Len(row, 9)
		for _, cell := range row {
			assert.Empty(cell)
		}
	}

	gd2, err := srv.PickCell(gameDef.GameId, 0, 0)
	assert.Nil(err)
	assert.NotEmpty(gd2.GameId)
	// response assertions, gameId not empty and 9X9 for the board size
	assert.Equal(gameDef.GameId, gd2.GameId)
	assert.Len(gd2.Board, 9)
	for _, row := range gd2.Board {
		assert.Len(row, 9)
	}
	// cell 0,0 shouldn't be flaged
	assert.NotEmpty(gd2.Board[0][0])
	assert.NotEqual("?", gd2.Board[0][0])
	if gd2.EndedGame {
		// if game ended cell 0,0 should have a mine
		assert.Equal("*", gd2.Board[0][0])
	} else {
		// if game didn't end cell 0,0 should have a value between 0 to 3
		cellVal, err := strconv.Atoi(gd2.Board[0][0])
		assert.Nil(err)
		assert.True(cellVal >= 0 && cellVal <= 3)
	}

}

func Test_Pick_Cell_Error_WrongGameId(t *testing.T) {
	assert := assert.New(t)

	srv := service.NewMinesweeperService()
	_, err := srv.PickCell("wrongGameId", 0, 0)
	assert.NotNil(err)
	coreErr := err.(*errors.CoreError)
	assert.Equal(400, coreErr.Code)
}

func Test_Pick_Cell_Error_EmptyGameId(t *testing.T) {
	assert := assert.New(t)

	srv := service.NewMinesweeperService()
	_, err := srv.PickCell("", 0, 0)
	assert.NotNil(err)
	coreErr := err.(*errors.CoreError)
	assert.Equal(400, coreErr.Code)
}

func Test_Pick_Cell_Error_WrongRowLess0(t *testing.T) {
	assert := assert.New(t)

	srv := service.NewMinesweeperService()
	gameDef, err := srv.NewGame()
	assert.Nil(err)
	// response assertions, gameId not empty and 9X9 for the board size
	assert.NotEmpty(gameDef.GameId)

	_, err = srv.PickCell(gameDef.GameId, -1, 0)
	assert.NotNil(err)
	coreErr := err.(*errors.CoreError)
	assert.Equal(400, coreErr.Code)
}

func Test_Pick_Cell_Error_WrongRowGraterThanMax(t *testing.T) {
	assert := assert.New(t)

	srv := service.NewMinesweeperService()
	gameDef, err := srv.NewGame()
	assert.Nil(err)
	// response assertions, gameId not empty and 9X9 for the board size
	assert.NotEmpty(gameDef.GameId)

	_, err = srv.PickCell(gameDef.GameId, 9, 0)
	assert.NotNil(err)
	coreErr := err.(*errors.CoreError)
	assert.Equal(400, coreErr.Code)
}

func Test_Pick_Cell_Error_WrongColLess0(t *testing.T) {
	assert := assert.New(t)

	srv := service.NewMinesweeperService()
	gameDef, err := srv.NewGame()
	assert.Nil(err)
	// response assertions, gameId not empty and 9X9 for the board size
	assert.NotEmpty(gameDef.GameId)

	_, err = srv.PickCell(gameDef.GameId, 0, -1)
	assert.NotNil(err)
	coreErr := err.(*errors.CoreError)
	assert.Equal(400, coreErr.Code)
}

func Test_Pick_Cell_Error_WrongColGraterThanMax(t *testing.T) {
	assert := assert.New(t)

	srv := service.NewMinesweeperService()
	gameDef, err := srv.NewGame()
	assert.Nil(err)
	// response assertions, gameId not empty and 9X9 for the board size
	assert.NotEmpty(gameDef.GameId)

	_, err = srv.PickCell(gameDef.GameId, 0, 9)
	assert.NotNil(err)
	coreErr := err.(*errors.CoreError)
	assert.Equal(400, coreErr.Code)
}

func Test_Add_Flag_Error_WrongGameId(t *testing.T) {
	assert := assert.New(t)

	srv := service.NewMinesweeperService()
	_, err := srv.AddFlag("wrongGameId", 0, 0)
	assert.NotNil(err)
	coreErr := err.(*errors.CoreError)
	assert.Equal(400, coreErr.Code)
}

func Test_Add_Flag_Error_EmptyGameId(t *testing.T) {
	assert := assert.New(t)

	srv := service.NewMinesweeperService()
	_, err := srv.AddFlag("", 0, 0)
	assert.NotNil(err)
	coreErr := err.(*errors.CoreError)
	assert.Equal(400, coreErr.Code)
}

func Test_Add_Flag_Error_WrongRowLess0(t *testing.T) {
	assert := assert.New(t)

	srv := service.NewMinesweeperService()
	gameDef, err := srv.NewGame()
	assert.Nil(err)
	// response assertions, gameId not empty and 9X9 for the board size
	assert.NotEmpty(gameDef.GameId)

	_, err = srv.AddFlag(gameDef.GameId, -1, 0)
	assert.NotNil(err)
	coreErr := err.(*errors.CoreError)
	assert.Equal(400, coreErr.Code)
}

func Test_Add_Flag_Error_WrongRowGraterThanMax(t *testing.T) {
	assert := assert.New(t)

	srv := service.NewMinesweeperService()
	gameDef, err := srv.NewGame()
	assert.Nil(err)
	// response assertions, gameId not empty and 9X9 for the board size
	assert.NotEmpty(gameDef.GameId)

	_, err = srv.AddFlag(gameDef.GameId, 9, 0)
	assert.NotNil(err)
	coreErr := err.(*errors.CoreError)
	assert.Equal(400, coreErr.Code)
}

func Test_Add_Flag_Error_WrongColLess0(t *testing.T) {
	assert := assert.New(t)

	srv := service.NewMinesweeperService()
	gameDef, err := srv.NewGame()
	assert.Nil(err)
	// response assertions, gameId not empty and 9X9 for the board size
	assert.NotEmpty(gameDef.GameId)

	_, err = srv.AddFlag(gameDef.GameId, 0, -1)
	assert.NotNil(err)
	coreErr := err.(*errors.CoreError)
	assert.Equal(400, coreErr.Code)
}

func Test_Add_Flag_Error_WrongColGraterThanMax(t *testing.T) {
	assert := assert.New(t)

	srv := service.NewMinesweeperService()
	gameDef, err := srv.NewGame()
	assert.Nil(err)
	// response assertions, gameId not empty and 9X9 for the board size
	assert.NotEmpty(gameDef.GameId)

	_, err = srv.AddFlag(gameDef.GameId, 0, 9)
	assert.NotNil(err)
	coreErr := err.(*errors.CoreError)
	assert.Equal(400, coreErr.Code)
}

func Test_Pick_Cell_GameOver_FoundMine_OK(t *testing.T) {
	assert := assert.New(t)

	srv := service.NewMinesweeperService()
	gameDef, err := srv.NewGame()
	assert.Nil(err)
	// response assertions, gameId not empty and 9X9 for the board size
	assert.NotEmpty(gameDef.GameId)
	assert.False(gameDef.EndedGame)
	assert.False(gameDef.Won)
	assert.Len(gameDef.Board, 9)
	for _, row := range gameDef.Board {
		assert.Len(row, 9)
		for _, cell := range row {
			assert.Empty(cell)
		}
	}

	gameEnded := false
	gd2 := model.GameDefintion{}
	for i := 0; i < len(gameDef.Board) && !gameEnded; i++ {
		for j := 0; j < len(gameDef.Board[i]) && !gameEnded; j++ {
			gd2, err = srv.PickCell(gameDef.GameId, i, j)
			assert.Nil(err)
			gameEnded = gd2.EndedGame
			if gameEnded {
				assert.Equal("*", gd2.Board[i][j])
			}
		}
	}
	assert.True(gameEnded)
	assert.False(gd2.Won)
}

func Test_Flag_Cell_Twice_OK(t *testing.T) {
	assert := assert.New(t)

	srv := service.NewMinesweeperService()
	gameDef, err := srv.NewGame()
	assert.Nil(err)
	// response assertions, gameId not empty and 9X9 for the board size
	assert.NotEmpty(gameDef.GameId)
	assert.False(gameDef.EndedGame)
	assert.False(gameDef.Won)
	assert.Len(gameDef.Board, 9)
	for _, row := range gameDef.Board {
		assert.Len(row, 9)
		for _, cell := range row {
			assert.Empty(cell)
		}
	}

	gd2, err := srv.AddFlag(gameDef.GameId, 0, 0)
	assert.Nil(err)
	// response assertions, gameId not empty and 9X9 for the board size
	assert.NotEmpty(gd2.GameId)
	assert.Equal(gameDef.GameId, gd2.GameId)
	assert.Len(gd2.Board, 9)
	for i, row := range gd2.Board {
		assert.Len(row, 9)
		for j, cell := range row {
			if i == 0 && j == 0 {
				// Cell 0,0 should have the flag
				assert.NotEmpty(cell)
				assert.Equal("?", gd2.Board[0][0])
			} else {
				// rest of the cells should be empty
				assert.Empty(cell)
			}
		}
	}

	gd2, err = srv.AddFlag(gameDef.GameId, 0, 0)
	assert.Nil(err)
	// response assertions, gameId not empty and 9X9 for the board size
	assert.NotEmpty(gd2.GameId)
	assert.Equal(gameDef.GameId, gd2.GameId)
	assert.Len(gd2.Board, 9)
	for _, row := range gd2.Board {
		assert.Len(row, 9)
		for _, cell := range row {
			// now all cells are empty as the flag was removed
			assert.Empty(cell)
		}
	}

}
