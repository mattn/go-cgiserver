package main

import (
	. "github.com/mattn/go-cgiserver"
	. "net/http"
	"os/exec"
)

func main() {
	c := CgiServer()
	c.DefaultApp = "blosxom.cgi"
	c.LangMap[".cgi"], _ = exec.LookPath("perl")
	c.LangMap[".php"], _ = exec.LookPath("php-cgi")
	ListenAndServe(":8080", c)
}
