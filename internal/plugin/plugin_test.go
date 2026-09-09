package plugin

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/Red-Sock/protoc-gen-docs/internal/request"
)

//go:embed input_examples/example_input.json
var genReq []byte

func getRequest() *pluginpb.CodeGeneratorRequest {
	return request.ParseJSON(genReq)
}

func TestPlugin_Generate(t *testing.T) {
	req := getRequest()

	pl := Plugin{
		Req:    req,
		Params: request.ParseInputParams(req.GetParameter()),
	}

	resp := pl.Generate()

	require.Len(t, resp.File, 1)
	require.Equal(t, "/docs.swagger_ui.go", resp.File[0].GetName())
	require.NotEmpty(t, resp.File[0].GetContent())
}

// Older protoc (<= 3.21.x, e.g. Debian's protobuf-compiler) never populates
// source_file_descriptors. The plugin must still generate from file_to_generate
// instead of panicking with an index-out-of-range.
func TestPlugin_Generate_NoSourceFileDescriptors(t *testing.T) {
	req := getRequest()
	req.SourceFileDescriptors = nil

	pl := Plugin{
		Req:    req,
		Params: request.ParseInputParams(req.GetParameter()),
	}

	resp := pl.Generate()

	require.Len(t, resp.File, 1)
	require.Equal(t, "/docs.swagger_ui.go", resp.File[0].GetName())
	require.NotEmpty(t, resp.File[0].GetContent())
}
