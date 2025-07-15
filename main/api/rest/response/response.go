package response

type Response struct {
	Status  string `json:"status"`
	Content []byte `json:"content"`
	Error   string `json:"error,omitempty"`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func Error(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}

func OK(content []byte) Response {
	return Response{
		Status:  StatusOK,
		Content: content,
	}

}
func OKNoContent() Response {
	return Response{
		Status: StatusOK,
	}
}
