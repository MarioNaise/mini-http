package main

import (
	"flag"
	"fmt"
	"mini-http/http"
	"os"
	"strings"
)

var (
	DIRNAME string
	PORT    uint
)

func init() {
	flag.StringVar(&DIRNAME, "directory", "", "")
	flag.UintVar(&PORT, "port", 4221, "")
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
		fmt.Println("Server listening on port", uint16(PORT))
	}, uint16(PORT))
}

func handleRoute(req *http.Request) http.Response {
	res := http.NewResponse(http.NotFound)
	if req.Path == "/" {
		res = http.NewResponse(http.Ok)
	}
	for k, v := range req.Headers {
		if strings.ToLower(k) == strings.TrimPrefix(req.Path, "/") {
			res = http.NewBodyResponse(v)
		}
	}
	return res
}

func handleFiles(req *http.Request) http.Response {
	reqFileName := strings.TrimPrefix(req.Path, "/files/")
	if DIRNAME != "" {
		DIRNAME += "/"
	}

	if req.Method == "GET" {
		file, err := os.ReadFile(DIRNAME + reqFileName)
		if err != nil {
			fmt.Println("Error reading file:", err.Error())
			return http.NewResponse(http.NotFound)
		}
		res := http.NewBodyResponse(string(file))
		res.Headers.Set("Content-Type", "application/octet-stream")
		return res
	}

	if req.Method == "POST" {
		content := []byte(req.Body)
		for i := range content {
			if content[i] == 0 {
				content = content[:i]
				break
			}
		}
		err := os.WriteFile(DIRNAME+reqFileName, content, 0644)
		if err != nil {
			fmt.Println("Error writing file:", err.Error())
			return http.NewResponse(http.Error)
		}
		return http.NewResponse(http.Created)
	}

	return http.NewResponse(http.NotFound)
}
