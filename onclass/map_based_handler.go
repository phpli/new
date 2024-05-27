package main

import "net/http"

type Routable interface {
	Route(method string, pattern string,
		handlerFunc func(ctx *Context))
}

type Handler interface {
	ServeHTTP(c *Context)
	Routable
}

func (h *HandlerBasedOnMap) Route(
	method string,
	pattern string,
	handlerFunc func(ctx *Context)) {

	key := h.key(method, pattern)
	h.handlers[key] = handlerFunc

	////TODO implement me
	//http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
	//	ctx := NewContext(w, r)
	//	handlerFunc(ctx)
	//})
	//panic("implement me")
}

type HandlerBasedOnMap struct {
	//key 应该是method+url =》string
	handlers map[string]func(ctx *Context)
}

func (h *HandlerBasedOnMap) ServeHTTP(c *Context) {
	//TODO implement me
	//panic("implement me")
	key := h.key(c.R.Method, c.R.URL.Path)
	if handler, ok := h.handlers[key]; ok {
		handler(NewContext(c.W, c.R))
	} else {
		c.W.WriteHeader(http.StatusNotFound)
		c.W.Write([]byte("Not Found"))
	}
}

func (h *HandlerBasedOnMap) key(method string, pattern string) string {
	return method + "#" + pattern
}

var _ Handler = &HandlerBasedOnMap{}

func NewHandlerBasedOnMap() Handler {
	return &HandlerBasedOnMap{
		handlers: make(map[string]func(ctx *Context)),
	}
}
