package apimux

import "net/http"

// QueryPair stores a key value to match a URL
type QueryPair struct {
	Key   string
	Value string
}

// NewQueryPair convenience function to create QueryPairs
func NewQueryPair(key, val string) *QueryPair {
	return &QueryPair{Key: key, Value: val}
}

// QueryGroup Comvenience function to favilitate the creation of
// apimux.Services that needs a query param matcher with "group_by" key.
func QueryGroup(val string) *QueryPair {
	return NewQueryPair("group_by", val)
}

// Route stores matchers for a mux.Router.
type Route struct {
	Name        string
	Method      string
	Path        string
	Query       *QueryPair
	HandlerFunc http.HandlerFunc
}

// Service is an Route slice. Is just a convenience typedef that
// improves semantics of the source code.
type Service []Route
