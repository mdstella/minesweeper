package model

import "fmt"

// This struct will be used internally to send information throwgh the service, dao and endpoint layer related with the game
type GameDefintion struct {
	GameId    string
	Board     [][]string
	EndedGame bool
	Won       bool
}

// This struct represent the board game, it has two different multidimensional arrays. The UserBoard will be used to retrieve to the client, will
// have the cells with it's values hidden. The GameBoard will be store in the server side and will be used to reveal the client picks
type Board struct {
	UserBoard [][]string
	GameBoard [][]string
}

// adding the stringer implementation for debuggin purpose
func (b Board) String() string {
	s := "--USER BOARD--\n"
	for i := 0; i < len(b.UserBoard); i++ {
		s += fmt.Sprint(b.UserBoard[i]) + "\n"
	}
	s += "--GAME BOARD--\n"
	for i := 0; i < len(b.GameBoard); i++ {
		s += fmt.Sprint(b.GameBoard[i]) + "\n"
	}
	return s
}
