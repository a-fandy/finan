package web

type GeneralResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewSuccessResponse(data interface{}) interface{} {
	return GeneralResponse{200, "Success", data}
}

func NewErrorResponse(message string, data interface{}) interface{} {
	return GeneralResponse{errorCode(message), message, data}
}

func errorCode(message string) (code int) {
	switch message {
	case "Bad Request":
		code = 400
	case "Not Found":
		code = 404
	case "Unauthorized":
		code = 401
	case "Forbidden":
		code = 403
	default:
		code = 500
	}
	return code
}
