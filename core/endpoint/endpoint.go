package endpoint

import (
	"golang.org/x/net/context"

	"github.com/go-kit/kit/endpoint"
	"github.com/mdstella/minesweeper/core/model"
	"github.com/mdstella/minesweeper/core/service"
)

//MakeNewGameEndpoint - endpoint to invoke for starting a new game.
func MakeNewGameEndpoint(svc service.MinesweeperService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		gameDefinition, err := svc.NewGame()
		return model.NewGameResponse{
			GameId: gameDefinition.GameId,
			Board:  gameDefinition.Board,
			Err:    err,
		}, nil
	}
}
