package request

import (
	"bufio"
	"strings"
	"testing"
)

var requestText = bufio.NewReader(
	strings.NewReader(
		"GET /foo HTTP1.1\r\n" +
			"Host: http://example.com\r\n" +
			"Content-Type: text/html\r\n" +
			"\r\n"))

func TestParse(t *testing.T) {
	req := new(HttpRequest)
	parser := HttpParser{requestText}

	t.Log("Given an HTTP I/O to parse")

	err := parser.Parse(req)
	if err != nil {
		t.Fatal("Parse failed!", err)
	}

	t.Log("it parses the Method\t")
	t.Log(req.Method)
	if req.Method != "GET" {
		t.Fatal("Fail!")
	}

	t.Log("it parses the Host\t")
	if req.Host != "http://example.com" {
		t.Fatal("Fail!")
	}
}
