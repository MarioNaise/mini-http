package http

type (
	Buffer []byte

	Status string

	Headers map[string]string

	Request struct {
		Method  string
		Path    string
		Headers Headers
		Body    string
	}

	Response struct {
		Status  Status
		Headers Headers
		Body    string
	}
)
