package typing

type Response struct {
	Status     string      `json:"status,omitempty"`
	StatusCode int         `json:"status_code,omitempty"`
	ErrorCode  int         `json:"error_code,omitempty"`
	ErrorMsg   string      `json:"error_msg,omitempty"`
	Metadata   interface{} `json:"metadata,omitempty"`
}

func (res *Response) SetResponse(msg string, data interface{}) {
	if msg != "" {
		res.ErrorMsg = msg
		res.ErrorCode = 400
		return
	}
	res.Status = "Success"
	res.StatusCode = 200
	res.Metadata = data
}
