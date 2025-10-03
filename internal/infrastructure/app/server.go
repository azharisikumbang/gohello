package app

import (
	"fmt"
	"net/http"
)

func Run() {

	mux := http.NewServeMux()
	LoadRoutes(mux)

	fmt.Println("Listening on port :8000")

	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		panic(err)
	}
}
