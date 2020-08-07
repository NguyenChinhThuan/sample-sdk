package main

import (
	"github.com/NguyenChinhThuan/sample-sdk/src/routers"
	"github.com/labstack/echo"
)

func main() {

	// new server
	e := echo.New()

	// setup router
	routers.Setup(e)

	// run with port :1323
	e.Logger.Fatal(e.Start(":1323"))
}
