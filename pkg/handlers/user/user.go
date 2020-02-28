package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/andreylm/bmlabs/pkg/db"
	"github.com/andreylm/bmlabs/pkg/entities"
	"github.com/andreylm/bmlabs/pkg/logger"
	"github.com/andreylm/bmlabs/pkg/utils"

	routing "github.com/qiangxue/fasthttp-routing"
)

// Handler ...
type Handler struct {
	db db.UserRepository
}

// NewUserHandler - create handler for user actions
func NewUserHandler(db db.UserRepository) *Handler {
	return &Handler{db: db}
}

// Create - create user action
func (u *Handler) Create(ctx *routing.Context) error {
	user := entities.User{}
	if err := json.Unmarshal(ctx.Request.Body(), &user); err != nil {
		logger.Get().Error(err)
		utils.MakeJSONResponseFromBytes(ctx, http.StatusInternalServerError, nil)
		return nil
	}

	if errors := utils.Validate(&user); len(errors) != 0 {
		utils.MakeJSONResponseFromMap(ctx, http.StatusBadRequest, errors)
		return nil
	}

	if err := u.db.Create(context.TODO(), &user); err != nil {
		logger.Get().Error(err)
		utils.MakeJSONResponseFromBytes(ctx, http.StatusInternalServerError, nil)
		return nil
	}

	utils.MakeJSONResponseFromBytes(ctx, http.StatusOK, nil)
	return nil
}

// Search - search users action
func (u *Handler) Search(ctx *routing.Context) error {
	s := SearchForm{}
	if err := json.Unmarshal(ctx.Request.Body(), &s); err != nil {
		logger.Get().Error(err)
		utils.MakeJSONResponseFromBytes(ctx, http.StatusInternalServerError, nil)
		return nil
	}

	if err := utils.ValidateSearch(&s.Search); err != nil {
		utils.MakeJSONResponseFromMap(ctx, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return nil
	}

	logger.Get().Info(s)
	return nil
}
