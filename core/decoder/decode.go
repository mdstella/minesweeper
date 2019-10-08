package decoder

import (
	"encoding/json"
	"net/http"

	"golang.org/x/net/context"

	"github.com/mdstella/minesweeper/core/model"
)

//DecodeSkeletonRequest - decode the skeleton request to generate the model request
func DecodeSkeletonRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request model.SkeletonRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
