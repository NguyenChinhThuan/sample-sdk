package sdk

import "github.com/labstack/echo"

// Restful is object contain function response
type Restful struct {
	Context echo.Context
}

// ResultOK is return RestResponseData  with status code  =200
func (r Restful) ResultOK(message string, data interface{}) *RestResponseData {
	return &RestResponseData{
		Code:    StatusCode.Ok,
		Message: message,
		Data:    data,
	}

}

// ResponseOK is response data with status 200
func (r Restful) ResponseOK(message string, data interface{}) error {
	var result = r.ResultOK(message, data)
	return r.Response(result)
}

// Response is response custom
func (r Restful) Response(result *RestResponseData) error {
	return r.Context.JSON(result.Code, result)
}
