package main

import (
	. "http"
	. "github.com/mattn/go-cgiserver/cgiserver"
	"exec"
)

func main() {
	c := CgiServer()
	c.DefaultApp = "blosxom.cgi"
	c.LangMap[".cgi"], _ = exec.LookPath("perl")
	ListenAndServe(":8080", c)
}
