package request

import (
	"bufio"
	"errors"
	"io"
	"log"
	"net/url"
	"strconv"
	"strings"
)

type HttpParser struct {
	buff *bufio.Reader
}

func NewHttpParser(r io.Reader) *HttpParser {
	return &HttpParser{
		buff: bufio.NewReader(r),
	}
}

// With a basic HttpParser object as a reciever, Parse will parse from the io.Reader and return
// HttpRequest as a result.
func (parser *HttpParser) Parse(req *HttpRequest) error {
	parser.parseRequestLine(req)
	parser.parseRequestHeaders(req)
	parser.parseRequestContent(req)
	parser.parseParams(req)

	return nil
}

func (parser *HttpParser) parseRequestLine(req *HttpRequest) error {
	ln, _, err := parser.buff.ReadLine()
	if err != nil {
		return err
	}
	log.Println(string(ln))

	tokens := strings.Fields(string(ln))
	if len(tokens) < 3 {
		return errors.New("unexpected token size on RequestLine")
	}

	req.Method = tokens[0]
	req.Path = tokens[1]
	req.Protocol = tokens[2]

	return nil
}

func (parser *HttpParser) parseRequestHeaders(req *HttpRequest) error {
	req.Headers = make(url.Values)

	for {
		ln, _, err := parser.buff.ReadLine()
		if err != nil {
			return err
		}
		log.Println(string(ln))

		if string(ln) == "" {
			break
		}

		tokens := strings.Split(string(ln), ": ")
		if len(tokens) < 2 {
			return errors.New("unexpected token size on Header")
		}

		if err = assignHeader(req, tokens...); err != nil {
			return err
		}
	}

	return nil
}

func (parser *HttpParser) parseRequestContent(req *HttpRequest) error {
	if contentLen := req.ContentLength; contentLen != 0 {
		contentBuff := make([]byte, contentLen)
		_, err := io.ReadFull(parser.buff, contentBuff)
		if err != nil {
			return err
		}
		log.Println(string(contentBuff))

		req.Body = string(contentBuff)
	}

	return nil
}

func (parser *HttpParser) parseParams(req *HttpRequest) error {
	switch req.Method {
	case "GET":
		url, err := url.ParseRequestURI(req.Path)
		if err != nil {
			return err
		}
		req.Params = url.Query()

	case "POST":
		if req.ContentType == "application/x-www-form-urlencoded" {
			req.Params, _ = url.ParseQuery(req.Body)
		}
	}

	return nil
}

func assignHeader(req *HttpRequest, tokens ...string) error {
	switch tokens[0] {
	case "Host":
		req.Host = tokens[1]
	case "User-Agent":
		req.UserAgent = tokens[1]
	case "Cookie":
		req.Cookie = tokens[1]
	case "Content-Type":
		req.ContentType = tokens[1]
	case "Content-Length":
		contentLength, err := strconv.Atoi(tokens[1])
		if err != nil {
			return err
		}
		req.ContentLength = contentLength
	default:
		req.Headers[tokens[0]] = strings.Split(tokens[1], ",")
	}

	return nil
}
