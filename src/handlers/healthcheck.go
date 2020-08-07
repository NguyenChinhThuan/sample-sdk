package handlers

import (
	"github.com/NguyenChinhThuan/sample-sdk/src/sdk"
	"github.com/labstack/echo"
)

// Healthcheck is handlers response server is Active
func Healthcheck(c echo.Context) error {
	var APIHandler = &sdk.Restful{
		Context: c,
	}

	return APIHandler.Response(&sdk.RestResponseData{
		Code:    sdk.StatusCode.Ok,
		Message: "pong",
	})
}
