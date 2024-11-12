package main

import (
	"github.com/Red-Sock/protoc-gen-docs/internal/plugin"
	"github.com/Red-Sock/protoc-gen-docs/internal/request"
	"github.com/Red-Sock/protoc-gen-docs/internal/response"
)

func main() {
	req := request.ReadFromStd()

	pl := plugin.Plugin{
		Req:    req,
		Params: request.ParseInputParams(req.GetParameter()),
	}

	resp := pl.Generate()
	response.WriteResponseToStd(resp)
}
