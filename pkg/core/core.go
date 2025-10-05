package core

import (
	"database/sql"
	"net/http"
)

type MiddlewareInterface interface {
	RunMiddleware(next http.Handler) http.HandlerFunc
}

type RouteInterface interface {
	GetMethod() string
	GetPath() string
	GetHandler() http.HandlerFunc
	GetMiddlewares() []MiddlewareInterface
}

type RouterInterface interface {
	GetRoutes() []RouteInterface
	Get(path string, handler func(http.ResponseWriter, *http.Request), ms []MiddlewareInterface)
	Post(path string, handler func(http.ResponseWriter, *http.Request), ms []MiddlewareInterface)
	Put(path string, handler func(http.ResponseWriter, *http.Request), ms []MiddlewareInterface)
	Patch(path string, handler func(http.ResponseWriter, *http.Request), ms []MiddlewareInterface)
	Delete(path string, handler func(http.ResponseWriter, *http.Request), ms []MiddlewareInterface)
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
