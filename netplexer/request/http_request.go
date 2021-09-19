package request

import (
	"io"
	"net/url"
)

// HttpRequest will contain all the parsed fields from the incoming request.
type HttpRequest struct {
	Protocol      string
	Method        string
	Host          string
	Path          string
	UserAgent     string
	Cookie        string
	ContentType   string
	ContentLength int
	Headers       url.Values
	Params        url.Values
	Body          string
}

// Creates a new HttpRequest that will read from the given io.Reader.
//
// That can be a net.Conn or simply a file, any implementation from io.Reader will work.
// The reader wiil be parse with the HTTP parser returning an error if it fails.
func NewHttpRequest(r io.Reader) (*HttpRequest, error) {
	req := new(HttpRequest)
	parser := NewHttpParser(r)

	err := parser.Parse(req)

	return req, err
}
