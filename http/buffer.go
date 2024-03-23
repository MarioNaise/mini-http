package http

import (
	"strings"
)

func (req Buffer) ToRequest() Request {
	reqSlice := strings.Split(string(req), CRLF+CRLF)
	headerSlice := strings.Split(reqSlice[0], CRLF)[1:]
	headers := Headers{}
	for _, header := range headerSlice {
		headerSplit := strings.Split(header, ": ")
		headers[strings.ToLower(headerSplit[0])] = headerSplit[1]
	}

	return Request{
		Method:  strings.Fields(reqSlice[0])[0],
		Path:    strings.Fields(reqSlice[0])[1],
		Headers: headers,
		Body:    reqSlice[1],
	}
}
