package app

import (
	"database/sql"
	"net/http"
)

type RouteInterface interface {
	GetMethod() string
	GetPath() string
	GetHandler() http.HandlerFunc
}

type RouterInterface interface {
	GetRoutes() []RouteInterface
	Get(path string, handler func(http.ResponseWriter, *http.Request))
	Post(path string, handler func(http.ResponseWriter, *http.Request))
	Put(path string, handler func(http.ResponseWriter, *http.Request))
	Patch(path string, handler func(http.ResponseWriter, *http.Request))
	Delete(path string, handler func(http.ResponseWriter, *http.Request))
}

type FeatureInterface interface {
	Boot(a *Application)
}

type DatabaseInterface interface {
	GetInstance() *sql.DB
}

type HTTPServerInterface interface {
	GetInstance() *http.ServeMux
}
