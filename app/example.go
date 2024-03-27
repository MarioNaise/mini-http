package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/MarioNaise/mini-http/http"
)

var (
	DIRNAME string
	ADDR    string
)

func init() {
	flag.StringVar(&DIRNAME, "directory", "", "")
	flag.StringVar(&ADDR, "addr", ":4221", "")
	flag.Parse()
}

func main() {
	server := http.NewServer()
	server.Handle("/", handleRoute)
	server.Handle("/files", handleFiles)
	server.Get("/echo", func(req *http.Request) http.Response {
		echo := strings.TrimPrefix(req.Path, "/echo/")
		return http.NewBodyResponse(strings.TrimSuffix(echo, "/"))
	})

	server.Listen(func() {
		fmt.Println("Server listening on", ADDR)
	}, ADDR)
}

func handleRoute(req *http.Request) http.Response {
	res := func(resp http.Response) http.Response {
		if req.Method != http.Get {
			return http.NewResponse(http.MethodNotAllowed)
		}
		return resp
	}
	if req.Path == "/" {
		return res(http.NewResponse(http.Ok))
	}
	for k, v := range req.Headers {
		if strings.ToLower(k) == strings.TrimPrefix(req.Path, "/") {
			return res(http.NewBodyResponse(v))
		}
	}

	return http.NewResponse(http.NotFound)
}

func handleFiles(req *http.Request) http.Response {
	// santitizing inputs would be a good idea here
	reqFileName := strings.TrimPrefix(req.Path, "/files/")
	if DIRNAME != "" {
		DIRNAME += "/"
	}

	if req.Method == http.Get {
		file, err := os.ReadFile(DIRNAME + reqFileName)
		if err != nil {
			fmt.Println("Error reading file:", err.Error())
			return http.NewResponse(http.NotFound)
		}
		res := http.NewBodyResponse(string(file))
		res.Headers.Set("Content-Type", "application/octet-stream")
		return res
	}

	if req.Method == http.Post {
		content := []byte(req.Body)
		content = content[:bytes.IndexByte(content, 0)]
		err := os.WriteFile(DIRNAME+reqFileName, content, 0644)
		if err != nil {
			fmt.Println("Error writing file:", err.Error())
			return http.NewResponse(http.Error)
		}
		return http.NewResponse(http.Created)
	}

	return http.NewResponse(http.NotFound)
}
