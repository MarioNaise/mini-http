package http

import "fmt"

func (headers Headers) Get(key string) string {
	val, ok := headers[key]
	if !ok {
		return ""
	}
	return val
}

func (headers Headers) Set(key, value string) {
	headers[key] = value
}

func (headers Headers) Del(key string) {
	delete(headers, key)
}

func (headers Headers) ToString() string {
	var s string
	for k, v := range headers {
		s += fmt.Sprintf("%s%s: %s", CRLF, k, v)
	}
	return s + CRLF + CRLF
}
