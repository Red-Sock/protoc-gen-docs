package plugin

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/Red-Sock/protoc-gen-docs/internal/request"
)

//go:embed input_examples/example_input.json
var genReq []byte

func TestPlugin(t *testing.T) {
	t.Run("SourceRelative", func(t *testing.T) {
		req := &pluginpb.CodeGeneratorRequest{
			Parameter: proto.String("paths=source_relative"),
			SourceFileDescriptors: []*descriptorpb.FileDescriptorProto{
				{
					Name: proto.String("api/grpc/hello_world_api.proto"),
					Options: &descriptorpb.FileOptions{
						GoPackage: proto.String("github.com/Red-Sock/protoc-gen-docs/example/pkg/docs;docs"),
					},
				},
			},
		}
		p := Plugin{
			Req:    req,
			Params: request.ParseInputParams(req.GetParameter()),
		}
		resp := p.Generate()
		require.NotNil(t, resp)
		require.Len(t, resp.File, 1)
		require.Equal(t, "api/grpc/docs/docs.swagger_ui.go", resp.File[0].GetName())
	})

	t.Run("Default", func(t *testing.T) {
		req := &pluginpb.CodeGeneratorRequest{
			Parameter: proto.String(""),
			SourceFileDescriptors: []*descriptorpb.FileDescriptorProto{
				{
					Name: proto.String("api/grpc/hello_world_api.proto"),
					Options: &descriptorpb.FileOptions{
						GoPackage: proto.String("github.com/Red-Sock/protoc-gen-docs/example/pkg/docs;docs"),
					},
				},
			},
		}
		p := Plugin{
			Req:    req,
			Params: request.ParseInputParams(req.GetParameter()),
		}
		resp := p.Generate()
		require.NotNil(t, resp)
		require.Len(t, resp.File, 1)
		require.Equal(t, "github.com/Red-Sock/protoc-gen-docs/example/pkg/docs/docs/docs.swagger_ui.go", resp.File[0].GetName())
	})
}

func getRequest() *pluginpb.CodeGeneratorRequest {
	return &pluginpb.CodeGeneratorRequest{}
}
