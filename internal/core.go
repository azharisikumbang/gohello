package core

import (
	"database/sql"
	"net/http"
)

type ApplicationInterface interface {
	GetHTTPServer() HTTPServerInterface
	UseHTTPServer(HTTPServerInterface)
	GetRouter() RouterInterface
	UseRouter(RouterInterface)
	GetDatabase() DatabaseInterface
	UseDatabase(DatabaseInterface)
	GetLogger()
	GetConfig() Config
	AddFeature(FeatureInterface)
	Run()
}

type Config struct {
	App AppConfig
	DB  DBConfig
}

type AppConfig struct {
	Port string
	Key  string
}

type DBConfig struct {
	Host     string
	Username string
	Password string
	Name     string
	Port     string
	Driver   string
}

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
	Boot(a ApplicationInterface)
}

type DatabaseInterface interface {
	GetInstance() *sql.DB
}

type HTTPServerInterface interface {
	GetInstance() *http.ServeMux
	UseInstance(*http.ServeMux)
}
