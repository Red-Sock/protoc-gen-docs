package main

import (
	"fmt"

	"github.com/Red-Sock/protoc-gen-docs/internal/request"
)

func main() {
	req := request.ReadFromStd()
	fmt.Println(req.String())
}
