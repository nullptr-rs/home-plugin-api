package router

import "github.com/Yggdrasil80/home-plugin-api/router/context"

type Router interface {
	AddRoute(method string, path string, handler func(ctx *context.Context) error)
	AddFile(path string, file string)
}
