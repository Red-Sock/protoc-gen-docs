package response

import (
	"io"
	"os"

	"google.golang.org/protobuf/proto"
	plugin "google.golang.org/protobuf/types/pluginpb"
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
