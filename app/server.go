package main

import (
	"fmt"
	"strings"

	// Uncomment this block to pass the first stage
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage

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

	buf := getReq(conn)
	path := getPath(string(buf))
	res := getRes(path)
	conn.Write([]byte(res))

	defer conn.Close()
	defer l.Close()
}

func getReq(conn net.Conn) []byte {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
		os.Exit(1)
	}
	return buf
}

func getPath(req string) string {
	reqHead := strings.Split(req, "\r\n")[0]
	return strings.Fields(reqHead)[1]
}

func getRes(path string) string {
	res := "HTTP/1.1 404 Not Found\r\n\r\n"
	if path == "/" {
		res = "HTTP/1.1 200 OK\r\n\r\n"
	}
	if strings.HasPrefix(path, "/echo") {
		echo := strings.TrimPrefix(path, "/echo/")
		res = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s\r\n", len(echo), echo)
	}
	return res
}
