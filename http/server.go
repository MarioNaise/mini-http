package http

import (
	"log"
	"net"
	"strings"
)

const defaultHandler = "default"

func NewServer() Server {
	return Server{routes: make(routeMap), Addr: ":4221"}
}

func (server *Server) Listen(callback func(), addr string) {
	server.Addr = addr
	l, err := net.Listen("tcp", server.Addr)
	if err != nil {
		log.Fatal(err.Error())
	}
	callback()

	defer closeListener(&l)
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("Error accepting connection: ", err.Error())
			continue
		}
		go handleConnection(conn, &server.routes)
	}
}

func handleConnection(conn net.Conn, routes *routeMap) {
	defer closeConnection(&conn)
	var buf buffer = make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println("Error reading from connection: ", err.Error())
		return
	}
	buf = buf[:n]
	req, err := buf.ToRequest()
	if err != nil {
		log.Println("Error parsing request: ", err.Error())
		return
	}
	res := handleRoutes(req, routes)
	writeToConnection(res.ToString(), &conn)
}

func handleRoutes(req Request, routes *routeMap) Response {
	res := NewResponse(NotFound)
	for _, handler := range (*routes)[defaultHandler] {
		if strings.HasPrefix(req.Path, handler.path) {
			res = handler.callback(&req)
		}
	}
	for _, handler := range (*routes)[req.Method] {
		if strings.HasPrefix(req.Path, handler.path) {
			res = handler.callback(&req)
		}
	}
	return res
}

func writeToConnection(response string, conn *net.Conn) {
	_, err := (*conn).Write([]byte(response))
	if err != nil {
		log.Println("Error writing to connection: ", err.Error())
		return
	}
}

func closeConnection(conn *net.Conn) {
	err := (*conn).Close()
	if err != nil {
		log.Fatal("Error closing connection: ", err.Error())
	}
}

func closeListener(l *net.Listener) {
	err := (*l).Close()
	if err != nil {
		log.Fatal("Error closing listener: ", err.Error())
	}
}

func (server *Server) Handle(path string, callback callbackFuncHandler) {
	server.routes[defaultHandler] = append(server.routes[defaultHandler], handler{callback: callback, path: path})
}

func (server *Server) Get(path string, callback callbackFuncHandler) {
	server.routes[Get] = append(server.routes[Get], handler{callback: callback, path: path})
}

func (server *Server) Post(path string, callback callbackFuncHandler) {
	server.routes[Post] = append(server.routes[Post], handler{callback: callback, path: path})
}
