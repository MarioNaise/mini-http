package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/http"
)

func handleRequest(req http.Request) string {
	pathRegex := regexp.MustCompile("[A-z-]+")
	path := pathRegex.FindString(req.Path)

	if req.Path == "/" {
		return http.NewResponse(http.Ok).ToString()
	}

	if path == "echo" {
		echo := strings.TrimPrefix(req.Path, "/echo/")
		return http.NewBodyResponse(strings.TrimSuffix(echo, "/")).ToString()
	}

	headerVal, ok := req.Headers[strings.ToLower(path)]
	if ok {
		return http.NewBodyResponse(headerVal).ToString()
	}

	if path == "files" {
		return fileRouteHandler(req)
	}

	return http.NewResponse(http.NotFound).ToString()
}

func fileRouteHandler(req http.Request) string {
	reqFileName := strings.TrimPrefix(req.Path, "/files/")
	if DIRNAME != "" {
		DIRNAME += "/"
	}

	if req.Method == "GET" {
		file, err := os.ReadFile(DIRNAME + reqFileName)
		if err != nil {
			fmt.Println("Error reading file:", err.Error())
			return http.NewResponse(http.NotFound).ToString()
		}
		res := http.NewBodyResponse(string(file))
		res.Headers.Set("Content-Type", "application/octet-stream")
		return res.ToString()
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
			return http.NewResponse(http.Error).ToString()
		}
		return http.NewResponse(http.Created).ToString()
	}

	return http.NewResponse(http.NotFound).ToString()
}
