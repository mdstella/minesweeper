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
	return model.NewGameRequest{}, nil
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

//DecodeAddFlagRequest - decode the addFlag request to generate the model request
func DecodeAddFlagRequest(_ context.Context, r *http.Request) (interface{}, error) {
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
