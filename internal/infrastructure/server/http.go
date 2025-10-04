package server

import "net/http"

func CreateHTTPServer() *http.ServeMux {
	return http.NewServeMux()
}
