package http

import (
	"net/http"
)

type Router struct {
	Hd Handler
}

func NewRouter(handler Handler) *Router {
	return &Router{Hd: handler}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	pathMain := req.URL.Path
	if pathMain == "/tasks" {
		switch req.Method {
		case http.MethodGet:
			if query := req.URL.Query(); len(query) > 0 {
				r.Hd.FindId(w, req)
			} else {
				r.Hd.GetTasks(w, req)
			}
		case http.MethodPost:
			if query := req.URL.Query(); query.Has("title") {
				r.Hd.SaveTask(w, req)
			}
		case http.MethodPut:
			if query := req.URL.Query(); query.Has("id") {
				r.Hd.UpdateTask(w, req)
			}
		case http.MethodDelete:
			if query := req.URL.Query(); query.Has("id") {
				r.Hd.DeleteTask(w, req)

			}
		default:
			r.Hd.ServeHTTP(w, req)
		}
	}
}
