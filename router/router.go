package router

import (
	"net/http"

	radix "github.com/armon/go-radix"
)

type RadixRouter struct {
	rtree *radix.Tree
}

func NewRadixRouter() *RadixRouter {
	return &RadixRouter{
		rtree: radix.New(),
	}
}

func (h RadixRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	prefix, handler, found := h.rtree.LongestPrefix(r.RequestURI)
	if found {
		r.RequestURI = r.RequestURI[len(prefix):]
		handler.(http.HandlerFunc)(w, r)
	} else {
		http.NotFound(w, r)
	}
}

func (h *RadixRouter) HandleFunc(route string, handler http.HandlerFunc) {
	h.rtree.Insert(route, handler)
}
