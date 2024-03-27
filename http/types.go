package http

type (
	buffer []byte

	Headers map[string]string

	Request struct {
		Method  string
		Path    string
		Headers Headers
		Body    string
	}

	Response struct {
		Status  string
		Headers Headers
		Body    string
	}

	callbackFuncHandler func(req *Request) Response
	handler             struct {
		callback callbackFuncHandler
		path     string
	}
	routeMap map[string][]handler
	Server   struct {
		routes routeMap
		Addr   string
	}
)
