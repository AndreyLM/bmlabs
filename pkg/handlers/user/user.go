package user

import (
	"github.com/andreylm/bmlabs/pkg/logger"
	routing "github.com/qiangxue/fasthttp-routing"
)

func Create(c *routing.Context) error {
	logger.Get().Info("Create")
	return nil
}

func Search(c *routing.Context) error {
	logger.Get().Info("Search")
	return nil
}
