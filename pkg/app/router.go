package app

import (
	"net/http"

	core "github.com/azharisikumbang/gohello/internal"
)

type Route struct {
	method      string
	path        string
	handler     http.HandlerFunc
	middlewares []core.MiddlewareInterface
}

func NewRoute(method string, path string, handler func(http.ResponseWriter, *http.Request), ms []core.MiddlewareInterface) core.RouteInterface {
	return &Route{
		method:      method,
		path:        path,
		handler:     handler,
		middlewares: ms,
	}
}

func (r *Route) GetMethod() string {
	return r.method
}

func (r *Route) GetPath() string {
	return r.path
}

func (r *Route) GetHandler() http.HandlerFunc {
	return r.handler
}

func (r *Route) GetMiddlewares() []core.MiddlewareInterface {
	return r.middlewares
}

type Router struct {
	Routes []core.RouteInterface
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) GetRoutes() []core.RouteInterface {
	return r.Routes
}

func (r *Router) Get(path string, handler func(http.ResponseWriter, *http.Request), ms []core.MiddlewareInterface) {
	route := NewRoute("GET", path, handler, ms)
	r.Routes = append(r.Routes, route)
}

func (r *Router) Post(path string, handler func(http.ResponseWriter, *http.Request), ms []core.MiddlewareInterface) {
	route := NewRoute("POST", path, handler, ms)
	r.Routes = append(r.Routes, route)
}

func (r *Router) Put(path string, handler func(http.ResponseWriter, *http.Request), ms []core.MiddlewareInterface) {
	route := NewRoute("PUT", path, handler, ms)
	r.Routes = append(r.Routes, route)
}

func (r *Router) Patch(path string, handler func(http.ResponseWriter, *http.Request), ms []core.MiddlewareInterface) {
	route := NewRoute("PATCH", path, handler, ms)
	r.Routes = append(r.Routes, route)
}

func (r *Router) Delete(path string, handler func(http.ResponseWriter, *http.Request), ms []core.MiddlewareInterface) {
	route := NewRoute("DELETE", path, handler, ms)
	r.Routes = append(r.Routes, route)
}
