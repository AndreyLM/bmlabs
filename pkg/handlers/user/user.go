package user

import (
	"encoding/json"
	"net/http"

	"github.com/andreylm/bmlabs/pkg/db"
	"github.com/andreylm/bmlabs/pkg/entities"
	"github.com/andreylm/bmlabs/pkg/logger"
	routing "github.com/qiangxue/fasthttp-routing"
)

type userHandler struct {
	db db.UserRepository
}

func NewUserHandler(db db.UserRepository) *userHandler {
	return &userHandler{db: db}
}

func (u *userHandler) Create(ctx *routing.Context) error {
	user := entities.User{}
	if err := json.Unmarshal(ctx.Request.Body(), &user); err != nil {
		logger.Get().Error(err)

		handlers.MakeResponse(ctx, http.StatusInternalServerError, nil)
		return nil
	}

	logger.Get().Info("Create")
	handlers.MakeResponse(ctx, http.StatusInternalServerError, nil)
	return nil
}

func (u *userHandler) Search(c *routing.Context) error {
	logger.Get().Info("Search")
	return nil
}
