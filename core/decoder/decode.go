package decoder

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"

	"github.com/mdstella/minesweeper/core/model"
)

//DecodeNewGameRequest - decode the newGame request to generate the model request
func DecodeNewGameRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request model.NewGameRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

//DecodePickCellRequest - decode the pickCell request to generate the model request
func DecodePickCellRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	gameId, ok := params["gameId"]
	if !ok {
		return nil, errors.New("Bad routing, game id not provided")
	}
	var request model.PickCellRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	request.GameId = gameId
	return request, nil
}
