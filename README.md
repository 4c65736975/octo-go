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
</br>

### Middlewares

<p align="right">&#x2191 <a href="#top">back to top</a></p>

## License

Distributed under the GPL-3.0 license. See [LICENSE](https://github.com/4c65736975/octo-go/blob/main/LICENSE) for more information.

<p align="right">&#x2191 <a href="#top">back to top</a></p>

## Acknowledgments

* [Choose an Open Source License](https://choosealicense.com)
* [Best README Template](https://github.com/othneildrew/Best-README-Template)
* [Octopus icons created by iconmama - Flaticon](https://www.flaticon.com/free-icons/octopus)

<p align="right">&#x2191 <a href="#top">back to top</a></p>
