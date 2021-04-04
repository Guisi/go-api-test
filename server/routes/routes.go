package routes

import (
	v1 "go-api-test/server/api/v1"
	"net/http"
)

var routes = []HttpRoute{
	//V1 APIs
	{Path: "/api/v1/systemStatus", Handler: v1.GetSystemStatus},
	{Path: "/api/v1/posts", Handler: v1.Posts},
}

func GetServerRouter() http.Handler {
	sm := http.NewServeMux()
	for _, route := range routes {
		sm.Handle(route.Path, route.Handler)
	}
	return sm
}
