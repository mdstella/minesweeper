package endpoint

import (
	"golang.org/x/net/context"

	"github.com/go-kit/kit/endpoint"
	"github.com/mdstella/minesweeper/core/model"
	"github.com/mdstella/minesweeper/core/service"
)

//MakeSkeletonEndpoint - endpoint to invoke the skeleton service.
func MakeSkeletonEndpoint(svc service.MinesweeperService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.SkeletonRequest)
		message, err := svc.Skeleton(req.Parameter)
		return model.SkeletonResponse{
			Message: message,
			Err:     err,
		}, nil
	}
}
