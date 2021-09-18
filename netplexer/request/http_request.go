package request

import (
	"io"
	"net/url"
)

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

func NewHttpRequest(r io.Reader) (*HttpRequest, error) {
	req := new(HttpRequest)
	parser := NewHttpParser(r)

	err := parser.Parse(req)

	return req, err
}
