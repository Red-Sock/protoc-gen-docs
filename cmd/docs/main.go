package main

import (
	"github.com/Red-Sock/protoc-gen-docs/internal/plugin"
	"github.com/Red-Sock/protoc-gen-docs/internal/request"
	"github.com/Red-Sock/protoc-gen-docs/internal/response"
)

// Output is defined by specifying docs_out parameter
// --docs_out=./pkg/docs - specifies where all code related to plugin will be generated

// plugin-specific parameters:
// 		"--docs" - common prefix for all parameters
// Base path
//		Use base_path option to specify
// 		e.g. --docs_opts base_path=/docs

func main() {
	req := request.ReadFromStd()

	pl := plugin.Plugin{
		Req:    req,
		Params: request.ParseInputParams(req.GetParameter()),
	}

	resp := pl.Generate()
	response.WriteResponseToStd(resp)
}
