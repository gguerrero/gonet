package response

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"

	"github.com/gguerrero/gonet/netplexer/request"
)

const (
	StatusOK = 200
)

var statusText = map[int]string{
	StatusOK: "OK",
}

var tpl *template.Template

func init() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	tplPath := filepath.Join(dir, "../../netplexer/templates/report.gohtml")
	tpl = template.Must(template.ParseFiles(tplPath))
}

type httpResponseWriter struct {
	writer io.Writer
}

// Returns an httpResonseWriter where you can write the response in...
func NewhttpResponseWriter(w io.Writer) *httpResponseWriter {
	return &httpResponseWriter{
		writer: w,
	}
}

func (rw *httpResponseWriter) Write(req *request.HttpRequest) {
	rw.writeRequestLine()
	rw.writeHeaders()
	rw.writeTemplate(req)
}

func (rw *httpResponseWriter) writeRequestLine() {
	fmt.Fprintf(rw.writer, "HTTP/1.1 %d %s\r\n", StatusOK, statusText[StatusOK])
}

func (rw *httpResponseWriter) writeHeaders() {
	fmt.Fprintf(rw.writer, "Content-Type: text/html\r\n\r\n")
}

func (rw *httpResponseWriter) writeTemplate(req *request.HttpRequest) {
	tpl.Execute(rw.writer, req)
}
