<div id="top"></div>
<br/>
<div align="center">
  <a href="https://github.com/4c65736975/octo-go">
    <img src="https://github.com/4c65736975/octo-go/assets/107006334/9e36e10c-c35c-4fef-b079-6b15ef622c41" alt="Logo" width="128" height="128">
  </a>
  <h3>Octo Go</h3>
  <p>
    Simple router for Go
    <br />
    <br />
    <a href="https://github.com/4c65736975/octo-go/issues">Report Bug</a>
    ·
    <a href="https://github.com/4c65736975/octo-go/issues">Request Feature</a>
    ·
    <a href="https://github.com/4c65736975/octo-go/blob/main/CHANGELOG.md">Changelog</a>
  </p>
</div>
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li>
          <a href="#prerequisites">Prerequisites</a>
        </li>
        <li>
          <a href="#installation">Installation</a>
        </li>
      </ul>
    </li>
    <li>
      <a href="#usage">Usage</a>
      <ul>
        <li>
          <a href="#routes">Routes</a>
        </li>
        <li>
          <a href="#group-routes">Group Routes</a>
        </li>
        <li>
          <a href="#middlewares">Middlewares</a>
          <ul>
            <li>
              <a href="#global-middlewares">Global Middlewares</a>
            </li>
            <li>
              <a href="#route-middlewares">Route Middlewares</a>
            </li>
            <li>
              <a href="#group-middlewares">Group Middlewares</a>
            </li>
          </ul>
        </li>
        <li>
          <a href="#parameters">Parameters</a>
        </li>
        <li>
          <a href="#query">Query</a>
        </li>
      </ul>
    </li>
    <li>
      <a href="#license">License</a>
    </li>
    <li>
      <a href="#acknowledgments">Acknowledgments</a>
    </li>
  </ol>
</details>

## About the project

Octo Go is a lightweight, high-performance router for creating robust HTTP endpoints in Go. It started its history as a router built entirely from scratch, but with the release of Go 1.22 Octo Go turned into a wrapper around the original mux, adding convenient creation of routes, middleware support and group routes. Octo Go is the perfect tool for simple creating powerful APIs in Go, staying as close to the original mux as possible, without unnecessary complexity.

<p align="right">&#x2191 <a href="#top">back to top</a></p>

## Getting started

```go
func main() {
  mux := router.NewRouter()

  mux.GET("/user", userGetHandler)
  mux.POST("/user", userPostHandler)

  http.ListenAndServe("localhost:3000", mux)

  // Registering route
  // METHOD: GET, POST, PUT, PATCH, DELETE
  // PATH: Route path e.g "/", "/users" or with parameters "/users/{id}", "/users/{id}/{group}"
  // HANDLER: func(w http.ResponseWriter, r *http.Request)
  // MIDDLEWARES: func(http.ResponseWriter, *http.Request, func()) you can add as many as you want

  mux.METHOD(PATH, HANDLER, MIDDLEWARES)

  // Registering global middleware
  // MIDDLEWARE: func(http.ResponseWriter, *http.Request, func())

  mux.Use(MIDDLEWARE)

  // Registering group of routes
  // BASE_PATH: default path for this group e.g "/products", all methods registered inside starts with this path

  mux.Group(BASE_PATH, func(mux *router.Router) {
    mux.METHOD(PATH, HANDLER, MIDDLEWARES)
  }, MIDDLEWARES)
}
```

### Prerequisites

[Go 1.22](https://go.dev/dl/)

### Installation

```sh
go get -u github.com/4c65736975/octo-go
```

<p align="right">&#x2191 <a href="#top">back to top</a></p>

## Usage

### Routes

```go
mux.GET("/", getHandler)
mux.PUT("/", putHandler)
mux.POST("/", postHandler)
mux.PATCH("/", patchHandler)
mux.DELETE("/", deleteHandler)
```

### Group Routes

```go
mux.Group("/users", func(mux *router.Router) {
  mux.GET("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "List all users")
  })
  mux.GET("/{id}", func(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    fmt.Fprintf(w, "Get user with ID %s\n", id)
  })
  mux.POST("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Create a new user")
  })
  mux.PUT("/{id}", func(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    fmt.Fprintf(w, "Update user with ID %s\n", id)
  })
  mux.DELETE("/{id}", func(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    fmt.Fprintf(w, "Delete user with ID %s\n", id)
  })
})

mux.Group("/products", func(mux *router.Router) {
  mux.GET("/", getProductsHandler)
  mux.GET("/{id}", getProductHandler)
  mux.POST("/", createProductHandler)
  mux.PUT("/{id}", updateProductHandler)
  mux.DELETE("/{id}", deleteProductHandler)
})
```

### Middlewares

```go
func exampleMiddleware(w http.ResponseWriter, req *http.Request, next func()) {
  fmt.Fprintln(w, "Example middleware")
  next()
}
```
If the next() function is not executed in the middleware, the request ends with this middleware, if it is executed, the next middleware or final handler is executed.

#### Global Middlewares
```go
mux.Use(loggingMiddleware)
mux.Use(ipMiddleware)
```
Global middleware works on every registered route after adding global middleware, all routes added earlier are ignored
```go
mux.GET("/products", productsHandler) // not using loggingMiddleware
mux.Use(loggingMiddleware)
mux.GET("/products/{id}", productHandler) // using loggingMiddleware
mux.DELETE("/products/{id}", deleteProductHandler) // using loggingMiddleware
```
#### Route Middlewares
```go
mux.GET("/profile", profileHandler, authMiddleware)
mux.GET("/profile/settings", settingsHandler, authMiddleware, settingsMiddleware, usageMiddleware)
```
#### Group Middlewares
```go
mux.Group("/users", func(mux *router.Router) {
  mux.GET("/", usersHandler)
  mux.DELETE("/{id}", deleteUserHandler, authMiddleware)
}, telemetryMiddleware)
```

Middlewares are executed in the order in which they are registered, unless the next() function is called before the middleware code as in the case below

```go
mux.GET("/", routeHandler, localMiddleware, localMiddleware2, localMiddleware3, localMiddleware4)

func localMiddleware(w http.ResponseWriter, req *http.Request, next func()) {
  fmt.Println("Local middleware 1")
  next()
}

func localMiddleware2(w http.ResponseWriter, req *http.Request, next func()) {
  next()
  fmt.Println("Local middleware 2")
}

func localMiddleware3(w http.ResponseWriter, req *http.Request, next func()) {
  fmt.Println("Local middleware 3")
  next()
}

func localMiddleware4(w http.ResponseWriter, req *http.Request, next func()) {
  fmt.Println("Local middleware 4")
  next()
}

Console Output:
Local middleware 1
Local middleware 3
Local middleware 4
Local middleware 2
```

### Parameters

With Go 1.22 we can register dynamic routes using the built-in mux. We can get the parameter in the handler or middleware as shown below.

```go
mux.GET("/users/{id}/products/{category}", func(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Handler")
  fmt.Println(req.PathValue("id"))
  fmt.Println(req.PathValue("category"))
}, testMiddleware)

func testMiddleware(w http.ResponseWriter, req *http.Request, next func()) {
  fmt.Println("Middleware")
  fmt.Println(req.PathValue("id"))
  fmt.Println(req.PathValue("category"))
  next()
}

Requested route: /users/1/products/human

Console Output:
Middleware
1
human
Handler
1
human
```

### Query

We can access the query value as with the default mux and similarly to parameters, in the handler and middleware.

```go
mux.GET("/users/{id}/products/{category}", func(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Handler")
  fmt.Println(req.URL.Query())
  fmt.Println(req.URL.Query().Get("maxHeight"))
}, testMiddleware)

func testMiddleware(w http.ResponseWriter, req *http.Request, next func()) {
  fmt.Println("Middleware")
  fmt.Println(req.URL.Query())
  fmt.Println(req.URL.Query().Get("maxHeight"))
  next()
}

Requested route: /users/1/products/human?sort=asc&maxHeight=180

Console Output:
Middleware
map[maxHeight:[180] sort:[asc]]
180
Handler
map[maxHeight:[180] sort:[asc]]
180
```

<p align="right">&#x2191 <a href="#top">back to top</a></p>

## License

Distributed under the GPL-3.0 license. See [LICENSE](https://github.com/4c65736975/octo-go/blob/main/LICENSE) for more information.

<p align="right">&#x2191 <a href="#top">back to top</a></p>

## Acknowledgments

* [Choose an Open Source License](https://choosealicense.com)
* [Best README Template](https://github.com/othneildrew/Best-README-Template)
* [Octopus icons created by iconmama - Flaticon](https://www.flaticon.com/free-icons/octopus)

<p align="right">&#x2191 <a href="#top">back to top</a></p>
