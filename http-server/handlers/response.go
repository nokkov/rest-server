package handlers

type Response struct {
	SHORT_URL string `json: "short_url, omitempty"`
	STATUS    string `json: "status"`
	ERROR     string `json: "error, omitempty"`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func OK() Response {
	return Response{
		STATUS: StatusOK,
	}
}

func Error(msg string) Response {
	return Response{
		STATUS: StatusError,
		ERROR:  msg,
	}
}
