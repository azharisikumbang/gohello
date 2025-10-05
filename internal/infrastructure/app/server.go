package app

import "github.com/azharisikumbang/gohello/internal/infrastructure/server"

func NewHTTPServer() HTTPServerInterface {
	return server.NewHTTPServer()
}
