package plugin

import (
	_ "embed"
	"testing"

	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"

	"github.com/Red-Sock/protoc-gen-docs/internal/request"
)

//go:embed input_examples/example_input.json
var genReq []byte

func TestPlugin(t *testing.T) {
	req := getRequest()
	_ = req
	println(1)
}

func getRequest() *plugin.CodeGeneratorRequest {
	return request.ParseJSON(genReq)
}
