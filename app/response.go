package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func getResponse(req request) string {
	pathRegex := regexp.MustCompile("[A-z-]+")
	path := pathRegex.FindString(req.path)

	if req.path == "/" {
		return resOk
	}

	if path == "echo" {
		echo := strings.TrimPrefix(req.path, "/echo/")
		return parseTextRes(strings.TrimSuffix(echo, "/"))
	}

	header, ok := req.headers[strings.ToLower(path)]
	if ok {
		return parseTextRes(header)
	}

	if path == "files" {
		return fileRouteHandler(req)
	}

	return resNotFound
}

func fileRouteHandler(req request) string {
	reqFileName := strings.TrimPrefix(req.path, "/files/")
	if DIRNAME != "" {
		DIRNAME += "/"
	}

	if req.method == "GET" {
		file, err := os.ReadFile(DIRNAME + reqFileName)
		if err != nil {
			fmt.Println("Error reading file:", err.Error())
			return resNotFound
		}
		return fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: application/octet-stream\r\nContent-Length: %d\r\n\r\n%s\r\n", len(string(file)), string(file))
	}

	if req.method == "POST" {
		content := []byte(req.body)
		for i := range content {
			if content[i] == 0 {
				content = content[:i]
				break
			}
		}
		err := os.WriteFile(DIRNAME+reqFileName, content, 0644)
		if err != nil {
			fmt.Println("Error writing file:", err.Error())
			return resError
		}
		return resCreated
	}

	return resNotAllowed
}
