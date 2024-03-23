package http

const (
	CRLF   string = "\r\n"
	Prefix string = "HTTP/1.1"

	Get  string = "GET"
	Post string = "POST"

	Ok         Status = "200 OK"
	Created    Status = "201 Created"
	NotFound   Status = "404 Not Found"
	NotAllowed Status = "405 Method Not Allowed"
	Error      Status = "500 Internal Server Error"
)
