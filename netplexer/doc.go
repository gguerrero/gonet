/*
Pacakge nexplexer is a TCP server that handles HTTP requests and returns information about them.

The package purpose is simply to learn about handling TCP connection in GoLang and
understand the basics of the HTTP protocol as well by creating it's own request/response handler.

Running the server

To run the server on your machine (inside the project folder):
	$ go run netplexer.go
	2021/09/19 19:37:31 Netplexer listening at 0.0.0.0:8000

Now you can request to http://localhost:8000 from you browser, any path, query string or form you
want to. The request will be be parsed and the response will contain the details of it.

Example POST form

Open the package template `form.html` on your browser to POST a form to the NetPLexer
*/
package netplexer
