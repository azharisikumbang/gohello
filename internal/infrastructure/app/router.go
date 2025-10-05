package app

import (
	"net/http"
)

type Route struct {
	method  string
	path    string
	handler http.HandlerFunc
}

func NewRoute(method string, path string, handler func(http.ResponseWriter, *http.Request)) RouteInterface {
	return &Route{
		method:  method,
		path:    path,
		handler: handler,
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

type Router struct {
	Routes []RouteInterface
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) GetRoutes() []RouteInterface {
	return r.Routes
}

func (r *Router) Get(path string, handler func(http.ResponseWriter, *http.Request)) {
	route := NewRoute("GET", path, handler)
	r.Routes = append(r.Routes, route)
}

func (r *Router) Post(path string, handler func(http.ResponseWriter, *http.Request)) {
	route := NewRoute("POST", path, handler)

	r.Routes = append(r.Routes, route)
}

func (r *Router) Put(path string, handler func(http.ResponseWriter, *http.Request)) {
	route := NewRoute("PUT", path, handler)

	r.Routes = append(r.Routes, route)
}

func (r *Router) Patch(path string, handler func(http.ResponseWriter, *http.Request)) {
	route := NewRoute("PATCH", path, handler)

	r.Routes = append(r.Routes, route)
}

func (r *Router) Delete(path string, handler func(http.ResponseWriter, *http.Request)) {
	route := NewRoute("DELETE", path, handler)

	r.Routes = append(r.Routes, route)
}
