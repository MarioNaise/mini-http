package http

const (
	CRLF    string = "\r\n"
	Version string = "HTTP/1.1"

	Get  string = "GET"
	Post string = "POST"

	Ok               string = "200 OK"
	Created          string = "201 Created"
	NotFound         string = "404 Not Found"
	MethodNotAllowed string = "405 Method Not Allowed"
	Error            string = "500 Internal Server Error"
)
