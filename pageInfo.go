package toot

import (
	"net/http"
	"strings"
)

// PageInfo is used to add pagination headers to a response
type PageInfo struct {
	MinID  string // If present, then a <prev> link header will be created
	MaxID  string // If present, then a <next> link header will be created
	Offset string // If present, then a <next> link header will be created
}

// SetHeaders adds pagination headers to a response.
func (p PageInfo) SetHeaders(response *http.Response) {

	if response == nil {
		return
	}

	// NPE Guard
	if response.Request == nil {
		return
	}

	// NPE Guard
	if response.Request.URL == nil {
		return
	}

	uri := response.Request.URL.String()
	uri, _, _ = strings.Cut(uri, "?")

	if prev := p.GetPreviousPage(uri); prev != "" {
		response.Header.Set("Link", "<"+prev+">; rel=\"prev\"")
	}

	if next := p.GetNextPage(uri); next != "" {
		response.Header.Set("Link", "<"+next+">; rel=\"next\"")
	}
}

// GetPreviousPage returns a URL for the previou page of results.
// If there is no previous page, then an empty string is returned.
func (p PageInfo) GetPreviousPage(path string) string {

	if p.MinID == "" {
		return ""
	}

	return path + "?min_id=" + p.MinID
}

// GetNextPage returns a URL for the next page of results.
// If there is no next page, then an empty string is returned.
func (p PageInfo) GetNextPage(path string) string {

	if p.MaxID != "" {
		return path + "?max_id=" + p.MaxID
	}

	if p.Offset != "" {
		return path + "?offset=" + p.Offset
	}

	return ""
}
