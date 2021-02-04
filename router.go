package main

import (
	"net/http"
)

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *Router) FindHandler(method string, path string) (http.HandlerFunc, bool, bool) {
	_, exist := r.rules[path]
	handler, methodExist := r.rules[path][method]
	return handler, methodExist, exist
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, methodExist, exist := r.FindHandler(request.Method, request.URL.Path)

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if !exist || !methodExist {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	handler(w, request)
}
