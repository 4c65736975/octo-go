// @author: 4c65736975, All Rights Reserved
// @version: 1.0.0.0, 14|02|2024
// @filename: router.go

package router

import (
	"fmt"
	"net/http"
)

type Router struct {
  mux *http.ServeMux
  groupPrefix string // I don't like this solution, but it works
  groupMiddlewares []Middleware
  globalMiddlewares []Middleware
}

// NewRouter allocates and returns a new [Router].
func NewRouter() *Router {
  return &Router{
    mux: http.NewServeMux(),
  }
}

type Middleware func(http.ResponseWriter, *http.Request, func())

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
  r.mux.ServeHTTP(w, req)
}

// Addes specified global middleware
func (r *Router) Use(middleware Middleware) {
  r.globalMiddlewares = append(r.globalMiddlewares, middleware)
}

// Creates group of endpoints with specified base path and middlewares
func (r *Router) Group(path string, route func(router *Router), middlewares ...Middleware) {
  r.groupPrefix = path
  r.groupMiddlewares = middlewares
  route(r)
  r.groupPrefix = ""
  r.groupMiddlewares = nil
}

func (r *Router) registerRoute(path string, method string, handler http.HandlerFunc, middlewares ...Middleware) {
  finalHandler := handler

  if len(middlewares) != 0 || len(r.globalMiddlewares) != 0 || len(r.groupMiddlewares) != 0 {
    allMiddlewares := append(r.globalMiddlewares, middlewares...)
    allMiddlewares = append(allMiddlewares, r.groupMiddlewares...)

    for i := len(allMiddlewares) - 1; i >= 0; i-- {
      finalHandler = useMiddleware(allMiddlewares[i], finalHandler)
    }
  }

  r.mux.HandleFunc(r.formatPattern(path, method), finalHandler)
}

func useMiddleware(middleware Middleware, handler http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, req *http.Request) {
    middleware(w, req, func() {
      handler.ServeHTTP(w, req)
    })
  }
}

func (r *Router) formatPattern(path string, method string) string {
  if r.groupPrefix != "" {
    return fmt.Sprintf("%s %s%s", method, r.groupPrefix, path)
  }

  return fmt.Sprintf("%s %s", method, path)
}

// Creates GET endpoint with specified path, handler and middlewares
func (r *Router) GET(path string, handler http.HandlerFunc, middlewares ...Middleware) {
  r.registerRoute(path, http.MethodGet, handler, middlewares...)
}

// Creates PUT endpoint with specified path, handler and middlewares
func (r *Router) PUT(path string, handler http.HandlerFunc, middlewares ...Middleware) {
  r.registerRoute(path, http.MethodPut, handler, middlewares...)
}

// Creates POST endpoint with specified path, handler and middlewares
func (r *Router) POST(path string, handler http.HandlerFunc, middlewares ...Middleware) {
  r.registerRoute(path, http.MethodPost, handler, middlewares...)
}

// Creates PATCH endpoint with specified path, handler and middlewares
func (r *Router) PATCH(path string, handler http.HandlerFunc, middlewares ...Middleware) {
  r.registerRoute(path, http.MethodPatch, handler, middlewares...)
}

// Creates DELETE endpoint with specified path, handler and middlewares
func (r *Router) DELETE(path string, handler http.HandlerFunc, middlewares ...Middleware) {
  r.registerRoute(path, http.MethodDelete, handler, middlewares...)
}