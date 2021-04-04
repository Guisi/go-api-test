package server

import (
	"go-api-test/inject"
	"go-api-test/service/post"
)

func initDependencyInjection() {
	// Required to allow multiple services to start at the same time in the dev env
	inject.LockInjector()
	defer inject.ReleaseInjector()

	// Services
	inject.RegisterSingleton("postService", post.NewPostService())

	// DAOs
	//inject.RegisterSingleton("chartsSourceDao", persist.NewChartsSourceDao())
}
