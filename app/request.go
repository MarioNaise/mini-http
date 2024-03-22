package main

import (
	"strings"
)

type request struct {
	method  string
	path    string
	headers map[string]string
	body    string
}

func parseRequest(req string) request {
	reqSlice := strings.Split(req, "\r\n\r\n")
	head, body := reqSlice[0], reqSlice[1]
	headerSlice := strings.Split(head, "\r\n")[1:]
	headers := map[string]string{}
	for _, header := range headerSlice {
		headerSplit := strings.Split(header, ": ")
		headers[strings.ToLower(headerSplit[0])] = headerSplit[1]
	}

	return request{
		method:  strings.Fields(head)[0],
		path:    strings.Fields(head)[1],
		headers: headers,
		body:    body,
	}
}
