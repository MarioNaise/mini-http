package main

import (
	"fmt"
)

var (
	resOk         string = "HTTP/1.1 200 OK\r\n\r\n"
	resCreated    string = "HTTP/1.1 201 Created\r\n\r\n"
	resNotFound   string = "HTTP/1.1 404 Not Found\r\n\r\n"
	resNotAllowed string = "HTTP/1.1 405 Method Not Allowed\r\n\r\n"
	resError      string = "HTTP/1.1 500 Internal Server Error\r\n\r\n"
)

func parseTextRes(str string) string {
	return fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s\r\n", len(str), str)
}
