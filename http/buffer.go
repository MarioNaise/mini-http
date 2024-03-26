package http

import (
	"errors"
	"strings"
)

func (buf Buffer) ToRequest() (Request, error) {
	reqSlice := strings.Split(string(buf), CRLF+CRLF)
	if len(reqSlice) < 2 || len(strings.Fields(reqSlice[0])) < 3 {
		return Request{}, errors.New("invalid request")
	}
	headerSlice := strings.Split(reqSlice[0], CRLF)[1:]
	headers := Headers{}
	for _, header := range headerSlice {
		headerSplit := strings.Split(header, ": ")
		if len(headerSplit) < 2 {
			continue
		}
		headers.Set(headerSplit[0], headerSplit[1])
	}

	req := Request{
		Method:  strings.Fields(reqSlice[0])[0],
		Path:    strings.Fields(reqSlice[0])[1],
		Headers: headers,
		Body:    reqSlice[1],
	}
	if !isValidMethod(req.Method) {
		return Request{}, errors.New("invalid method")
	}
	if !isValidPath(req.Path) {
		return Request{}, errors.New("invalid path")
	}
	if strings.Fields(reqSlice[0])[2] != Version {
		return Request{}, errors.New("invalid http version")
	}
	return req, nil
}

func isValidMethod(method string) bool {
	return method == Get || method == Post
}

func isValidPath(path string) bool {
	return len(path) > 0 && path[0] == '/'
}
