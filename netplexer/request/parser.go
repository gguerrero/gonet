package request

import "io"

// Parser defines the interface for any protocol parser, expects an io.Reader and will return an
// error.
type Parser interface {
	Parse(io.Reader) error
}
