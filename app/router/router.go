package router

import (
	"net/http"

	"jcp-backend/controller"

	"github.com/labstack/echo/v4"
)

func Router() *echo.Echo {
	/*
		라우터 관리
	*/
	e := echo.New()

	e.GET("/healthy", func(c echo.Context) error {
		return c.String(http.StatusOK, "Healthy")
	})

	project_group := e.Group("/project")
	{
		project_group.GET("/create", controller.CreateProject)
	}

	return e
}
