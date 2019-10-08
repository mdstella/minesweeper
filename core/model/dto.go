package model

// This struct will be used internally to send information throwgh the service, dao and endpoint layer related with the game
type GameDefintion struct {
	GameId string
	Board  [][]string
}
