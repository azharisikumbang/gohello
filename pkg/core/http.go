package core

import (
	"net/http"
)

type HTTPServer struct{}

func NewHTTPServer() *HTTPServer {
	return &HTTPServer{}
}

func (s *HTTPServer) GetInstance() *http.ServeMux {
	return http.NewServeMux()
}
