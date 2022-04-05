package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateProject(c echo.Context) error {
	/*
		프로젝트 생성
	*/
	example_data := []string{"hello"}
	return c.JSON(http.StatusOK, example_data)
}
