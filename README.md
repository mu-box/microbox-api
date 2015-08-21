# Nanobox Api

Sets up a simple api with logging of requests and a single route. More routes may be added later.

## Routes

| path | description | payload | response |
| --- | --- | --- | --- |
| `/ping` | a simple ping pong route | nil | `pong` |


### Example usage

This is how to start the api.
```go
package main

import (
  "github.com/jcelliott/lumber"
  "github.com/pagodabox/na-api"
)

func main() {
  api.Name = "EXAMPLE"
  api.Logger = lumber.NewConsoleLogger("INFO")
  // this is for storing what ever is needed by the routes
  // api.User = nil
  api.Start("127.0.0.1:8080")
}

```

This is one of the routes file, there can be as many as needed
```go
package example

import (
  "github.com/pagodabox/na-api"
  "net/http"
)

func init() {
  api.Router.Get("/test", api.TraceRequest(testEndpoint))
}

func testEndpoint(res http.ResponseWriter, resq *http.Request) {
  res.WriteHeader(200)
  res.Write([]byte("This has been a sucessfull test"))
}

```