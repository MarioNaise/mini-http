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

	callback func(req *Request) Response
	handler  struct {
		callback callback
		path     string
	}
	routeMap map[string][]handler
	Server   struct {
		routes routeMap
		Host   string
		Port   uint16
	}
)
