package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"jcp-backend/utils"
)

var (
	conf *utils.Config
)

func init() {
	conf = utils.LoadConfig()
}

// 프로젝트 생성
func CreateProject(c echo.Context) error {
	return c.JSON(http.StatusOK, conf.Sample)
}
