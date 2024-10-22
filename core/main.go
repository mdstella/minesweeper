// Package classification Minesweeper API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API for Minesweeper game
//
//     Schemes: http, https
//     Host: localhost
//     BasePath: /v1
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/mdstella/minesweeper/core/decoder"
	"github.com/mdstella/minesweeper/core/endpoint"
	"github.com/mdstella/minesweeper/core/errors"
	"github.com/mdstella/minesweeper/core/service"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/rs/cors"
	"golang.org/x/net/context"
)

//go:generate swagger generate spec -o swagger-ui/swagger.json
func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	port := os.Getenv("PORT")

	if port == "" {
		logger.Log("$PORT must be set")
		port = ":8000"
	}

	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	srv := service.NewMinesweeperService()
	r := mux.NewRouter()

	// We generate the handler functions that will be executed once each endpoint is invoked. It will use a go-kit endpoint
	// (the one used by the framework selected to create the server), a decoder that will decode the message sent to the endpoint
	// into the model, and a common response encoder
	newGameHandler := httptransport.NewServer(endpoint.MakeNewGameEndpoint(srv), decoder.DecodeNewGameRequest, EncodeResponse)
	// Define the method, path and the handler in the router to be able to dispatch requests to it.

	// swagger:route POST /game minesweeper NewGameRequest
	//
	// Generates a new Minesweeper game
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       default: NewGameResponse
	//       200: NewGameResponse
	r.Methods("POST").Path("/minesweeper/v1/game").Handler(newGameHandler)

	// Adding routing for picking a cell
	// swagger:route POST /game/:gameId minesweeper PickCellRequest
	//
	// Picks and reveal a cell
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       default: PickCellResponse
	//       200: PickCellResponse
	pickCellHandler := httptransport.NewServer(endpoint.MakePickCellEndpoint(srv), decoder.DecodePickCellRequest, EncodeResponse)
	r.Methods("POST").Path("/minesweeper/v1/game/{gameId}").Handler(pickCellHandler)

	// Adding routing for adding a flag
	// swagger:route POST /flag/:gameId minesweeper AddFlagRequest
	//
	// Add a flag to a cell
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       default: PickCellResponse
	//       200: PickCellResponse
	addFlagHandler := httptransport.NewServer(endpoint.MakeAddFlagEndpoint(srv), decoder.DecodeAddFlagRequest, EncodeResponse)
	r.Methods("POST").Path("/minesweeper/v1/flag/{gameId}").Handler(addFlagHandler)

	// adding swagger endpoint to have the API doc available
	swaggerUrl := "/swagger-ui/"
	r.PathPrefix(swaggerUrl).Handler(http.StripPrefix(swaggerUrl, http.FileServer(http.Dir("./swagger-ui/"))))

	handler := cors.Default().Handler(r)

	logger.Log("msg", "HTTP", "addr", port)
	logger.Log("err", http.ListenAndServe(port, handler))

}

// errorer is implemented by all concrete response types that may contain errors.
type errorer interface {
	Error() error
}

// EncondeResponse will encode the response message that will be sent to the client
func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if e, ok := response.(errorer); ok && e.Error() != nil {
		EncodeHttpError(ctx, e.Error(), w)
		return nil
	}

	return json.NewEncoder(w).Encode(response)
}

// HttpCodeFrom retrieve the http status code, if it's not a core error will be consider an internal server error
func HttpCodeFrom(err error) int {
	if e, ok := err.(errors.CoreError); ok {
		return e.Code
	} else if e, ok := err.(*errors.CoreError); ok {
		return e.Code
	}
	return http.StatusInternalServerError
}

// EncodeHttpError encode the error response
func EncodeHttpError(ctx context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}

	httpStatus := HttpCodeFrom(err)
	w.WriteHeader(httpStatus)

	// gokit adds the prefix "Encode:" when encoding a msg and "Decode: " when error in the decoding phase
	message := strings.TrimPrefix(strings.TrimPrefix(strings.TrimPrefix(err.Error(), "Encode: "), "Decode: "), "Do: ")
	var errorResponse errors.ErrorResponse
	if nerr := errors.GetCoreError(err); nerr != nil {
		errorResponse = errors.ErrorResponse{
			Error: errors.CoreError{Message: nerr.Message, Code: nerr.Code, Type: nerr.Type}}
	} else {
		errorResponse = errors.ErrorResponse{Error: errors.CoreError{Message: message}}
	}

	json.NewEncoder(w).Encode(errorResponse)
}
