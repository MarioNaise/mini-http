package http

import (
	"fmt"
	"net"
	"os"
	"strings"
)

const defaultHandler = "default"

func NewServer() Server {
	return Server{routes: make(routeMap), Host: "", Port: 4221}
}

func (server *Server) Listen(callback func(), port uint16) {
	server.Port = port
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.Host, server.Port))
	if err != nil {
		fmt.Printf("Failed to bind to port %d\n", server.Port)
		os.Exit(1)
	}
	callback()

	defer closeListener(&l)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			return
		}
		go handleConnection(conn, &server.routes)
	}
}

func handleConnection(conn net.Conn, routes *routeMap) {
	defer closeConnection(&conn)
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
	res := NewResponse(NotFound)
	handleRoutes(req, &res, routes)
	writeConnection(res.ToString(), &conn)
}

func handleRoutes(req Request, res *Response, routes *routeMap) {
	for _, handler := range (*routes)[defaultHandler] {
		if strings.HasPrefix(req.Path, handler.path) {
			*res = handler.callback(&req)
		}
	}
	for _, handler := range (*routes)[req.Method] {
		if strings.HasPrefix(req.Path, handler.path) {
			*res = handler.callback(&req)
		}
	}
}

func writeConnection(response string, conn *net.Conn) {
	_, err := (*conn).Write([]byte(response))
	if err != nil {
		return
	}
}

func closeConnection(conn *net.Conn) {
	err := (*conn).Close()
	if err != nil {
		return
	}
}

func closeListener(l *net.Listener) {
	err := (*l).Close()
	if err != nil {
		return
	}
}

func (server *Server) Handle(path string, callback callback) {
	server.routes[defaultHandler] = append(server.routes[defaultHandler], handler{callback: callback, path: path})
}

func (server *Server) Get(path string, callback callback) {
	server.routes[Get] = append(server.routes[Get], handler{callback: callback, path: path})
}

func (server *Server) Post(path string, callback callback) {
	server.routes[Post] = append(server.routes[Post], handler{callback: callback, path: path})
}
