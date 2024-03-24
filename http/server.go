package http

import (
	"fmt"
	"net"
	"os"
)

func NewServer() Server {
	return Server{Host: "0.0.0.0", Port: 4221}
}

func (server Server) Listen(callback func(req *Request) Response, port uint16) {
	server.Port = port
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.Host, server.Port))
	if err != nil {
		fmt.Printf("Failed to bind to port %d\n", server.Port)
		os.Exit(1)
	}
	fmt.Printf("Server listening on %d\n", server.Port)

	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			return
		}
		go handleConnection(conn, callback)
	}
}

func handleConnection(conn net.Conn, callback func(req *Request) Response) {
	defer conn.Close()
	var buf Buffer = make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading from connection: ", err.Error())
		return
	}
	req, err := buf.ToRequest()
	if err != nil {
		fmt.Println("Error parsing request: ", err.Error())
		return
	}
	res := callback(&req)
	conn.Write([]byte(res.ToString()))
}
