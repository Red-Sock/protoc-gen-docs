package response

import (
	"io"
	"os"

	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

func WriteResponseToStd(resp *plugin.CodeGeneratorResponse) {
	WriteResponse(os.Stdout, resp)
}

func WriteResponse(writer io.Writer, resp *plugin.CodeGeneratorResponse) {
	cont, err := proto.Marshal(resp)
	if err != nil {
		panic(err.Error())
	}

	_, err = writer.Write(cont)
	if err != nil {
		panic(err.Error())
	}
}
