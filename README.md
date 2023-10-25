# toot

[![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://pkg.go.dev/github.com/benpate/toot)
[![Version](https://img.shields.io/github/v/release/benpate/toot?include_prereleases&style=flat-square&color=brightgreen)](https://github.com/benpate/toot/releases)
[![Build Status](https://img.shields.io/github/actions/workflow/status/benpate/toot/go.yml?style=flat-square)](https://github.com/benpate/toot/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/benpate/toot?style=flat-square)](https://goreportcard.com/report/github.com/benpate/toot)
[![Codecov](https://img.shields.io/codecov/c/github/benpate/toot.svg?style=flat-square)](https://codecov.io/gh/benpate/toot)

## Mastodon API Server for Go

Toot implements a REST API server that matches the [Mastodon REST API](https://docs.joinmastodon.org/api/).  It plugs into your existing application seamlessly (using [router adapter](#routers)) making your Go application accessible to any Mastodon-compatible client app.

## Why would I do this?

Most client apps use the Mastodon REST API instead of the official ActivityPub Client-to-Server (C2S) API.  However, the Mastodon API is very large, with a huge number of routes, handlers, and scopes required to support.  Toot provides a consistent catalog of all of these values, and handles lots of busywork -- matching the specific routes, authorization, parameters, and results that Mastodon requires -- so that you can write happly little API handlers like this:

```go
api.PostStatus = func(authorization model.Authorization, values txn.PostStatus) (object.Status, error) {
	// Do the thing that:
	// 1) creates a new `status`
	// 2) writes it to the database
	// 3) returns the value to the caller
}

// Once your handlers are defined, connect Toot to your router like this:
e := echo.New()
tootecho.Register(e, api)
```

## Routers

Currently, I have only written a connector to the echo Router (because that's what I'm using for my own applications).  It should, however, be a simple project to write adapters for other routers, such as [Chi](https://github.com/go-chi/chi), [Gin](https://github.com/gin-gonic/gin), [Fiber](https://github.com/gofiber/fiber), or others.  If you'd like to use Toot with a different router, please get in touch and let's work together.

## Project Status (Alpha?)

This project is still in early alpha testing, as I continue filling out the Mastodon support in [Emissary](https://emissary.dev).  Some interfaces may change in the future if I run into trouble making things work with this reference implementation.

However, at this point, feedback from other developers will be immensely helpful in making Tooth a useful tool for the Go community.  Please feel free to open issues, or submit pull requests.