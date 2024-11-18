package request

import (
	"encoding/json"
	"io"
	"os"

	errors "github.com/Red-Sock/trace-errors"
	"google.golang.org/protobuf/proto"
	plugin "google.golang.org/protobuf/types/pluginpb"
)

func ReadFromStd() *plugin.CodeGeneratorRequest {
	return Read(os.Stdin)
}

func Read(reader io.Reader) *plugin.CodeGeneratorRequest {
	content, err := io.ReadAll(reader)
	if err != nil {
		panic(errors.Wrap(err, "error reading request"))
	}

	return Parse(content)
}

func Parse(content []byte) *plugin.CodeGeneratorRequest {
	req := &plugin.CodeGeneratorRequest{}
	err := proto.Unmarshal(content, req)
	if err != nil {
		panic(err)
	}

	return req
}

func ParseJSON(content []byte) *plugin.CodeGeneratorRequest {
	req := &plugin.CodeGeneratorRequest{}
	err := json.Unmarshal(content, req)
	if err != nil {
		panic(err)
	}

	return req
}
