package main

import (
	"net/http"

	todov1connect "github.com/raveesh-me/todo/gen/connect/protos/todo/v1/todov1connect"
	"github.com/raveesh-me/todo/server"
	"github.com/rs/zerolog/log"
)

func main() {
	mux := http.NewServeMux()
	todoServer := server.NewTodoServer()

	path, handler := todov1connect.NewTodoServiceHandler(todoServer)
	mux.Handle(path, handler)

	log.Info().Msg("Starting server on :8080")
	http.ListenAndServe(":8080", mux)
}
