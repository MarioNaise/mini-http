# Mini-Http

This Project started out as the Codecrafters challenge
[Build your own HTTP server](https://app.codecrafters.io/courses/http-server)

I turned it into a small module, that can be used to create a simple HTTP server to handle GET and POST requests.

## Useful to know

When passing a path to a handler function, all routes, starting with that given path, will be handled:

- `server.Get("/",...` will handle `"/"`, but also `"/hello"`

Additionally, the order of setting up your handler functions matters.
If you handle `"/hello"` first, and `"/"` afterwards, the second routehandler could potentially overwrite the response from your first route.

## Examples

**You can find an example server file in `app/example.go`**

- Create new server:

```
server := http.NewServer()
```

- Handle a route for all Methods:

```
server.Handle("/", func(req *http.Request) http.Response {
    if req.Method == http.Get || req.path == "/hello-world" {
        return http.NewBodyResponse("Hello, World!")
    } else {
        return http.NewResponse(http.NotFound)
    }
})
```

- Handle a GET Route:

```
server.Get("/", func(req *http.Request) http.Response {
    return http.NewBodyResponse("Hello world!")
})
```

- Handle a POST Route:

```
server.Post("/", func(req *http.Request) http.Response {
    return http.NewBodyResponse("Hello world!")
})
```

- After setting up all your routes, you can listen for connections:

```
server.Listen(func() {
		fmt.Println("Server listening on port", uint16(PORT))
	}, uint16(PORT))
```
