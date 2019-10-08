package decoder

import (
	"encoding/json"
	"net/http"

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
