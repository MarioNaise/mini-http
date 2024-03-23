package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/http"
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
	http.NewServer().Listen(handleRequest, uint16(PORT))
}

func handleRequest(req *http.Request) http.Response {
	pathRegex := regexp.MustCompile("[A-z-]+")
	path := pathRegex.FindString(req.Path)

	if req.Path == "/" {
		return http.NewResponse(http.Ok)
	}

	if path == "echo" {
		echo := strings.TrimPrefix(req.Path, "/echo/")
		return http.NewBodyResponse(strings.TrimSuffix(echo, "/"))
	}

	headerVal, ok := req.Headers[strings.ToLower(path)]
	if ok {
		return http.NewBodyResponse(headerVal)
	}

	if path == "files" {
		return fileRouteHandler(req)
	}

	return http.NewResponse(http.NotFound)
}

func fileRouteHandler(req *http.Request) http.Response {
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
