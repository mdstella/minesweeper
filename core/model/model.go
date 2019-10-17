package model

// swagger:parameters NewGameRequest
type NewGameRequest struct{}

// New game with a clean board to play and a gameId identifying that board
// swagger:response NewGameResponse
type NewGameResponse struct {
	// The gameId
	// in: body
	GameId string `json:"gameId,omitempty"`
	// TODO - maybe we can add a cell struct in the future when the cell has more fields and logic,
	// now we can represent the board with a multidimension (2 dimensions) array with the cell value as a string

	// The game board
	// in: body
	Board [][]string `json:"board"`
	//swagger:ignore
	Err error `json:"error,omitempty"`
}

// Implementing error method
func (r NewGameResponse) Error() error { return r.Err }

// swagger:parameters PickCellRequest AddFlagRequest
type PickCellRequest struct {
	// The row of the cell
	// required: true
	// in: body
	Row int `json:"row"`
	// The column of the cell
	// required: true
	// in: body
	Column int `json:"column"`
	// The game id
	// required: true
	// in: path
	GameId string `json:"gameId"`
}

// Picking/flaging a cell, the users receives the updated board and statuses flags
// swagger:response PickCellResponse
type PickCellResponse struct {
	// The gameId
	// in: body
	GameId string `json:"gameId,omitempty"`
	// The ended game flag
	// in: body
	EndedGame bool `json:"endedGame"`
	// The won flag
	// in: body
	Won bool `json:"won"`
	// The game board
	// in: body
	Board [][]string `json:"board"`
	//swagger:ignore
	Err error `json:"error,omitempty"`
}

// Implementing error method
func (r PickCellResponse) Error() error { return r.Err }
