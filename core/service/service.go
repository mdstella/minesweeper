package service

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/mdstella/minesweeper/core/errors"

	"github.com/golang/groupcache/lru"
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

	// NewGame is for starting a new game creating the board
	NewGame() (model.GameDefintion, error)

	// PickCell allows the user to pick a cell on a given game/board
	PickCell(gameId string, row, column int) (model.GameDefintion, error)
}

type MinesweeperSrvImpl struct {
	// For this first stage will keep the boards on memory, in this map we will have the gameId as a key and the Board struct as value
	// This map will be acces with a mutex to be sure we are not modifying it at the same time by different requests.
	games *lru.Cache
	mutex *sync.RWMutex
}

func NewMinesweeperService() MinesweeperService {
	return &MinesweeperSrvImpl{
		// initialize the LRU in 10, to avoid storing to much information in memory, will allow 10 games at the same time
		games: lru.New(10),
		mutex: &sync.RWMutex{},
	}
}

////////NEW GAME///////////

// NewGame is for starting a new game creating the board
func (msi *MinesweeperSrvImpl) NewGame() (model.GameDefintion, error) {
	board, err := msi.generateBoard(ROWS, COLS, MINES_COUNT)
	if err != nil {
		return model.GameDefintion{}, err
	}

	gameId := shortuuid.New()
	msi.mutex.Lock()
	msi.games.Add(gameId, board)
	msi.mutex.Unlock()

	return model.GameDefintion{
		GameId: gameId,
		Board:  board.UserBoard,
	}, nil
}

// generateBoard - will generate a board with the configuration sent by parameter (rows, columns and mines)
// this method starts with lower case as it's not exposed and it can't be invoked outside the service layer.
func (msi *MinesweeperSrvImpl) generateBoard(rows, cols, mines int) (model.Board, error) {

	if rows == 0 || cols == 0 || rows*cols < mines {
		return model.Board{}, errors.NewBadParamError("Wrong configuration for the game, review the amount of rows, columns and mines")
	}

	gameBoard := make([][]string, rows)
	userBoard := make([][]string, rows)

	// first stage of board generation, empty board with
	for i := 0; i < rows; i++ {
		gameBoard[i] = make([]string, cols)
		userBoard[i] = make([]string, cols)
	}

	// Adding random mines on the board
	for i := 0; i < mines; i++ {
		row, column := msi.getRandomCell(rows, cols)
		if gameBoard[row][column] != "*" {
			gameBoard[row][column] = "*"
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
			if gameBoard[i][j] != "*" {
				gameBoard[i][j] = strconv.Itoa(msi.calculateCloseMines(gameBoard, i, j))
			}
		}
	}

	board := model.Board{
		GameBoard: gameBoard,
		UserBoard: userBoard,
	}

	fmt.Println(board)
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

/////////PICK A CELL////////////////

// PickCell allows the user to pick a cell on a given game/board
func (msi *MinesweeperSrvImpl) PickCell(gameId string, row, column int) (model.GameDefintion, error) {
	gameId = strings.TrimSpace(gameId)
	if gameId == "" {
		return model.GameDefintion{}, errors.NewBadParamError("GameId is empty")
	}

	if row < 0 || row >= ROWS {
		return model.GameDefintion{}, errors.NewBadParamError(fmt.Sprintf("Wrong row value should be between 0 and %d", ROWS))
	}

	if column < 0 || column >= COLS {
		return model.GameDefintion{}, errors.NewBadParamError(fmt.Sprintf("Wrong column value should be between 0 and %d", COLS))
	}

	msi.mutex.RLock()
	boardIntf, ok := msi.games.Get(gameId)
	if !ok {
		msi.mutex.RUnlock()
		return model.GameDefintion{}, errors.NewBadParamError(fmt.Sprintf("Error trying to obtain game by id %s. Please start a new game", gameId))
	}
	msi.mutex.RUnlock()
	board := boardIntf.(model.Board)

	cellItem := board.GameBoard[row][column]
	board.UserBoard[row][column] = cellItem

	// if we found a mine we notify the game ended and remove from the cache the game
	if cellItem == "*" {
		// remove from the cache
		msi.mutex.Lock()
		msi.games.Remove(gameId)
		msi.mutex.Unlock()

		return model.GameDefintion{
			Board:     board.UserBoard,
			EndedGame: true,
			Won:       false,
			GameId:    gameId,
		}, nil
	}

	// if it's not a mine we update the cache and retrieve the new board to the client
	// TODO --> adding flag and notify when the user wins
	msi.mutex.Lock()
	msi.games.Add(gameId, board)
	msi.mutex.Unlock()

	fmt.Println(board)

	return model.GameDefintion{
		Board:     board.UserBoard,
		EndedGame: false,
		Won:       false,
		GameId:    gameId,
	}, nil
}
