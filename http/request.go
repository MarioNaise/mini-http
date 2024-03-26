package http

import "fmt"

func NewRequest() Request {
	return Request{Method: Get, Path: "/", Headers: Headers{}, Body: ""}
}

func NewBodyRequest(body string) Request {
	req := Request{Method: Get, Path: "/", Headers: Headers{}, Body: ""}
	req.SetBody(body)
	return req
}

func (req *Request) SetBody(body string) {
	req.Body = body
	if body != "" {
		req.Headers.Set("Content-Length", fmt.Sprintf("%d", len(body)))
	}
}

func (req *Request) ToString() string {
	return fmt.Sprintf("%s %s %s", req.Method, req.Path, Version+req.Headers.ToString()+req.Body)
}
