package handlers

type response struct {
	Status string 	`json:"status"`
	Data	any 	`json:"data,omitempty"`
}

func NewResp(status string, data any) *response {
	return &response{
		Status: status,
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