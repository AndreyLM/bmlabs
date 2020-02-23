package router

import (
	"github.com/andreylm/bmlabs/pkg/db"
	userHandler "github.com/andreylm/bmlabs/pkg/handlers/user"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

type Router struct {
	router *routing.Router
	storer db.Storer
}

func NewRouter(storer db.Storer) *Router {
	return &Router{
		router: routing.New(),
		storer: storer,
	}
}

func (r *Router) Init() {
	uHandler := userHandler.NewUserHandler(r.storer.User())
	users := r.router.Group("/users")
	users.Post("", uHandler.Create)
	users.Post("/search", uHandler.Search)
}

func (r *Router) HandleRequest(ctx *fasthttp.RequestCtx) {
	r.router.HandleRequest(ctx)
}
