package http

import (
	"fmt"
	"net/http"
)

type LogMiddlware struct{}

func NewLogMiddlware() *LogMiddlware {
	return &LogMiddlware{}
}

func (m *LogMiddlware) RunMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("run middlware above")
		next.ServeHTTP(w, r)
		fmt.Println("run middlware again")
	})
}
