package main

import (
	"jcp-backend/router"
)

func main() {
	echo := router.Router()
	echo.Debug = true

	echo.Logger.Fatal(echo.Start(":1323"))
}
