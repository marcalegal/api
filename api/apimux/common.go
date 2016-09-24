package apimux

import "fmt"

// AcceptHeader is a convenience function that provides minimal API versioning
// string.
func AcceptHeader(apiVersion int) string {
	return fmt.Sprintf("application/vnd.marcalegal+json; version=%d", apiVersion)
}
