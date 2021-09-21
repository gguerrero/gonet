package request

import (
	"bufio"
	"strings"
	"testing"
)

var requestText = bufio.NewReader(strings.NewReader(`
GET /foo HTTP1.1
Content-Type:text/html

`))

func TestParse(t *testing.T) {
	req := new(HttpRequest)
	parser := HttpParser{requestText}

	t.Fatal("nOOOO")
	t.Log("Test parser.Parse(req)")
	err := parser.Parse(req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(req)
}
