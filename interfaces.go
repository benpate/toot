package toot

import "net/http"

// APIFunc represents an API method.
type APIFunc[AuthToken ScopesGetter, Input any, Output any] func(AuthToken, Input) (Output, error)

// APIFuncWithHeader represents an API method that includes custom header results.
type APIFuncWithHeader[AuthToken ScopesGetter, Input any, Output any] func(AuthToken, Input) (Output, http.Header, error)

// Authorizer represents a function that can authorize an HTTP request.
// If the Bearer token is valid, then it returns an AuthToken object (which is also a ScopesGetter).
// If the request is not authorized, then it returns an error.
type Authorizer[AuthToken ScopesGetter] func(*http.Request) (AuthToken, error)

// ScopesGetter represents an object that can return a list of Scopes that it includes.
// It is used as the base for AuthTokens so that we can verify allowed scopes separately from the
// custom Authorizer code.
type ScopesGetter interface {
	Scopes() []string
}
