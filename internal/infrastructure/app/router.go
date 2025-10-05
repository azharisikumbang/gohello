package app

import (
	"net/http"
)

type Route struct {
	method      string
	path        string
	handler     http.HandlerFunc
	middlewares []MiddlewareInterface
}

func NewRoute(method string, path string, handler func(http.ResponseWriter, *http.Request), ms []MiddlewareInterface) RouteInterface {
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

func (r *Route) GetMiddlewares() []MiddlewareInterface {
	return r.middlewares
}

type Router struct {
	Routes []RouteInterface
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) GetRoutes() []RouteInterface {
	return r.Routes
}

func (r *Router) Get(
	path string,
	handler func(http.ResponseWriter, *http.Request),
	ms []MiddlewareInterface) {
	route := NewRoute("GET", path, handler, ms)
	r.Routes = append(r.Routes, route)
}

func (r *Router) Post(path string, handler func(http.ResponseWriter, *http.Request), ms []MiddlewareInterface) {
	route := NewRoute("POST", path, handler, ms)
	r.Routes = append(r.Routes, route)
}

func (r *Router) Put(path string, handler func(http.ResponseWriter, *http.Request), ms []MiddlewareInterface) {
	route := NewRoute("PUT", path, handler, ms)
	r.Routes = append(r.Routes, route)
}

func (r *Router) Patch(path string, handler func(http.ResponseWriter, *http.Request), ms []MiddlewareInterface) {
	route := NewRoute("PATCH", path, handler, ms)
	r.Routes = append(r.Routes, route)
}

func (r *Router) Delete(path string, handler func(http.ResponseWriter, *http.Request), ms []MiddlewareInterface) {
	route := NewRoute("DELETE", path, handler, ms)
	r.Routes = append(r.Routes, route)
}
