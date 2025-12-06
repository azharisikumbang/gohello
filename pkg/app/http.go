package app

import (
	"net/http"
)

type HTTPServer struct {
	Mux *http.ServeMux
}

func NewHTTPServer() *HTTPServer {
	return &HTTPServer{
		Mux: http.NewServeMux(),
	}
}

func (s *HTTPServer) GetInstance() *http.ServeMux {
	return s.Mux
}

func (s *HTTPServer) UseInstance(mux *http.ServeMux) {
	s.Mux = mux
}
