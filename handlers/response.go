package handlers

type response struct {
	Success bool 	`json:"success"`
	Data	any 	`json:"data"`
}

func NewResp(success bool, data any) *response {
	return &response{
		Success: success,
		Data: data,
	}
}

type errorResp struct {
	ErrorCode			int	`json:"error"`
	Message				any
	}

func NewError( errCode int, msg any) *errorResp {
	return &errorResp{
		ErrorCode: errCode,
		Message: msg,
	}
}