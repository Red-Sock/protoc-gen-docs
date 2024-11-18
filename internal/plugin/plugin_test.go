package plugin

import (
	_ "embed"
	"testing"

	"google.golang.org/protobuf/types/pluginpb"

	"github.com/Red-Sock/protoc-gen-docs/internal/request"
)

//go:embed input_examples/example_input.json
var genReq []byte

func TestPlugin(t *testing.T) {
	req := getRequest()
	_ = req
	println(1)
}

func getRequest() *pluginpb.CodeGeneratorRequest {
	return request.ParseJSON(genReq)
}
