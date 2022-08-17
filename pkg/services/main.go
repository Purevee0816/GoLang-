package services

// Response : resposne type
type Response struct {
	Code        int                 `json:"code" mapstructure:"code"`
	Status      string              `json:"status" mapstructure:"status"`
	Message     string              `json:"message" mapstructure:"message"`
	MessageCode string              `json:"message_code" mapstructure:"message_code"`
	Result      interface{}         `json:"result" mapstructure:"result"`
	ValidError  []map[string]string `json:"valid_error" mapstructure:"valid_error"`
}

// Success :
func (r *Response) Success() *Response {
	return r.SuccessWithParams(map[string]interface{}{})
}

// SuccessWithParams :
func (r *Response) SuccessWithParams(res interface{}) *Response {
	return &Response{
		Code:    200,
		Status:  "success",
		Message: "Success",
		Result:  res,
	}
}

// BadRequest valid:
func (r *Response) BadRequestValid(v []map[string]string) *Response {
	return &Response{
		Code:       400,
		Status:     "error",
		Message:    "Error.",
		ValidError: v,
	}
}

// BadRequest :
func (r *Response) BadRequestWithMessageAndCode(message string) *Response {
	return &Response{
		Code:    400,
		Status:  "error",
		Message: message,
	}
}

// Unauthorized  :
func (r *Response) Unauthorized() *Response {
	return &Response{
		Code:    401,
		Status:  "error",
		Message: "Unauthorized error.",
	}
}

// NotFound :
func (r *Response) NotFound() *Response {
	return &Response{
		Code:    404,
		Status:  "error",
		Message: "Not found page.",
	}
}

// ErrorWithMessage :
func (r *Response) ErrorWithMessage(message string) *Response {
	return &Response{
		Code:    500,
		Status:  "error",
		Message: message,
	}
}
