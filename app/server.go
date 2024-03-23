package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/codecrafters-io/http-server-starter-go/http"
)

var (
	DIRNAME string
	PORT    int
)

func init() {
	flag.StringVar(&DIRNAME, "directory", "", "")
	flag.IntVar(&PORT, "port", 4221, "")
	flag.Parse()
}

func main() {
	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", PORT))
	if err != nil {
		fmt.Printf("Failed to bind to port %d\n", PORT)
		os.Exit(1)
	}
	fmt.Printf("Server listening on %d\n", PORT)

	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			return
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	var buf http.Buffer = make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
		return
	}
	res := handleRequest(buf.ToRequest())
	conn.Write([]byte(res))
}
