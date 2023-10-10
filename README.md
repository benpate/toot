# Hannibal / toot

[![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://pkg.go.dev/github.com/benpate/toot)
[![Version](https://img.shields.io/github/v/release/benpate/toot?include_prereleases&style=flat-square&color=brightgreen)](https://github.com/benpate/toot/releases)
[![Build Status](https://img.shields.io/github/actions/workflow/status/benpate/toot/go.yml?style=flat-square)](https://github.com/benpate/toot/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/benpate/toot?style=flat-square)](https://goreportcard.com/report/github.com/benpate/toot)
[![Codecov](https://img.shields.io/codecov/c/github/benpate/toot.svg?style=flat-square)](https://codecov.io/gh/benpate/toot)

## Mastodon API Server for Go

## Why would I do this?

Most Fediverse client apps do not (yet) support the official ActivityPub Client-to-Server (C2S) API.  But most use the Mastodon REST API.  Using `toot`, you can open your Go server application to a wide array of client apps.