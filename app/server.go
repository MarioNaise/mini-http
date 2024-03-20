package main

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"
)

type request struct {
	method  string
	path    string
	headers map[string]string
	body    string
}

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	req := getReq(conn)
	res := getRes(req)
	conn.Write([]byte(res))

	defer conn.Close()
	defer l.Close()
}

func getReq(conn net.Conn) request {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
		os.Exit(1)
	}
	req := string(buf)
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

func getRes(req request) string {
	pathRegex := regexp.MustCompile("[A-z|-]+")
	path := pathRegex.FindString(req.path)
	resNotFound := "HTTP/1.1 404 Not Found\r\n\r\n"
	resOk := "HTTP/1.1 200 OK\r\n\r\n"
	resText := "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\n"
	if req.path == "/" {
		return resOk
	}
	if path == "echo" {
		echo := strings.TrimPrefix(req.path, "/echo/")
		return fmt.Sprintf("%sContent-Length: %d\r\n\r\n%s\r\n", resText, len(echo), echo)
	}
	header, ok := req.headers[strings.ToLower(path)]
	if ok {
		return fmt.Sprintf("%sContent-Length: %d\r\n\r\n%s\r\n", resText, len(header), header)
	}
	return resNotFound
}
