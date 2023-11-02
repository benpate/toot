package toot

import "net/http"

// APIFunc represents an API method. It can return either a single result or a paged result.
type APIFunc[AuthToken ScopesGetter, Input any, Output any] interface {
	APIFunc_SingleResult[AuthToken, Input, Output] | APIFunc_PagedResult[AuthToken, Input, Output]
}

// APIFunc_SingleResult represents an API method that returns a single result.
type APIFunc_SingleResult[AuthToken ScopesGetter, Input any, Output any] func(AuthToken, Input) (Output, error)

// APIFunc_PagedResult represents an API method that returns a many results with pagination data.
type APIFunc_PagedResult[AuthToken ScopesGetter, Input any, Output any] func(AuthToken, Input) (Output, PageInfo, error)

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
