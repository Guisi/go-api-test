package routes

import "net/http"

type HttpRoute struct {
	Path    string
	Handler http.HandlerFunc
}
