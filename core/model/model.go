package model

type NewGameRequest struct{}

type NewGameResponse struct {
	GameId string `json:"gameId,omitempty"`
	// TODO - maybe we can add a cell struct in the future when the cell has more fields and logic,
	// now we can represent the board with a multidimension (2 dimensions) array with the cell value as a string
	Board [][]string `json:"board"`
	Err   error      `json:"error,omitempty"`
}

// Implementing error method
func (r NewGameResponse) Error() error { return r.Err }
