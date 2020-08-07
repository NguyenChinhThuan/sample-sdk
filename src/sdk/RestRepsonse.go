package sdk

// RestResponseData is data return from API
type RestResponseData struct {
	Code    int         `json:"code" validate:"required"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Total   int         `json:"total,omitempty"`
}

// EnumStatus const
type EnumStatus struct {
	Ok        int
	Error     int
	NotFound  int
	Invalid   int
	Forbidden int
}

// StatusCode const
var StatusCode = &EnumStatus{
	Ok:        200,
	Error:     500,
	NotFound:  404,
	Invalid:   400,
	Forbidden: 403,
}
