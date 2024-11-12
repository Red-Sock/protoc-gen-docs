package main

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/Red-Sock/protoc-gen-docs/example/pkg/docs"
)

func main() {
	p, h := docs.Swagger()
	http.Handle(p, h)
	logrus.Info("Checkout http://localhost:8080/non_default_docs_path")
	http.ListenAndServe(":8080", nil)
}
