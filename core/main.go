package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/mdstella/minesweeper/core/decoder"
	"github.com/mdstella/minesweeper/core/endpoint"
	"github.com/mdstella/minesweeper/core/service"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	srv := &service.MinesweeperSrvImpl{}

	newGameHandler := httptransport.NewServer(endpoint.MakeNewGameEndpoint(srv), decoder.DecodeNewGameRequest, EncodeResponse)

	http.Handle("/minesweeper/v1/game", newGameHandler)

	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))

}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
