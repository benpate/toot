package toot

import "net/http"

// APIFunc represents an API method.  Inputs and outputs may vary
// but the essential function signature remains the same
type APIFunc[Input any, Output any] func(*http.Request, Input) (Output, error)
