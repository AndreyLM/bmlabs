package router

import (
	userHandlers "github.com/andreylm/bmlabs/pkg/handlers/user"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

type Router struct {
	router *routing.Router
}

func NewRouter() *Router {
	return &Router{router: routing.New()}
}

func (r *Router) Init() {
	users := r.router.Group("/users")
	users.Get("", userHandlers.Create)
	users.Get("/search", userHandlers.Search)
}

func (r *Router) HandleRequest(ctx *fasthttp.RequestCtx) {
	r.router.HandleRequest(ctx)
}
