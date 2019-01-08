package controllers

import (
	"net/http"

	u "github.com/gpng/go-docker-api-boilerplate/pkg/utils"
)

// HelloWorld returns a hello world http response
var HelloWorld = func(w http.ResponseWriter, r *http.Request) {
	u.Respond(w, u.Message(true, "Hello World!"))
}
