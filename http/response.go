package http

import (
	"fmt"
)

func NewResponse(s Status) Response {
	return Response{Status: s, Headers: Headers{}, Body: ""}
}

func NewBodyResponse(body string) Response {
	res := Response{Status: Ok, Headers: Headers{}, Body: ""}
	res.SetBody(body)
	return res
}

func (res *Response) SetStatus(status Status) {
	res.Status = status
}

func (res *Response) SetBody(body string) {
	res.Body = body
	if body != "" {
		res.Headers.Set("Content-Length", fmt.Sprintf("%d", len(body)))
		if res.Headers.Get("Content-Type") == "" {
			res.Headers.Set("Content-Type", "text/plain")
		}
	} else {
		res.Headers = Headers{}
	}
}

func (res Response) ToString() string {
	response := fmt.Sprintf("%s %s", Prefix, res.Status)
	response += res.Headers.ToString()
	return response + CRLF + CRLF + res.Body
}
