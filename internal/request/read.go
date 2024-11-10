package request

import (
	"io"
	"os"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

func ReadFromStd() *plugin.CodeGeneratorRequest {
	return Read(os.Stdin)
}

func Read(reader io.Reader) *plugin.CodeGeneratorRequest {
	content, err := io.ReadAll(reader)
	if err != nil {
		panic(errors.Wrap(err, "error reading request"))
	}

	req := &plugin.CodeGeneratorRequest{}

	err = proto.Unmarshal(content, req)
	if err != nil {
		panic(err)
	}

	return req
}
