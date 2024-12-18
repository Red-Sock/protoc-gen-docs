package plugin

import (
	"strings"
	"unicode"

	"github.com/Red-Sock/toolbox"
	"google.golang.org/protobuf/proto"
	plugin "google.golang.org/protobuf/types/pluginpb"

	"github.com/Red-Sock/protoc-gen-docs/internal/plugin/patterns/doc_handler_template"
	"github.com/Red-Sock/protoc-gen-docs/internal/request"
)

type Plugin struct {
	Req    *plugin.CodeGeneratorRequest
	Params request.InputParams
}

func (p *Plugin) Generate() *plugin.CodeGeneratorResponse {
	resp := &plugin.CodeGeneratorResponse{}

	resp.SupportedFeatures = proto.Uint64(
		uint64(plugin.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL))
	resp.File = append(resp.File, p.generateHandlerParameter())

	return resp
}

func (p *Plugin) generateHandlerParameter() *plugin.CodeGeneratorResponse_File {
	var swaggerUIHandler plugin.CodeGeneratorResponse_File

	swaggerUIHandler.Name = toolbox.ToPtr("/docs.swagger_ui.go")

	genReq := doc_handler_template.SwaggerUIGenReq{
		BasePath:       p.Params.BasePath,
		SwaggerWebPath: p.Params.SwaggerWebPath,
		SwaggerFolder:  p.Params.SwaggerFolderPath,
		Tittle:         p.Params.Title,
		Specs:          make([]doc_handler_template.Spec, 0, len(p.Req.SourceFileDescriptors)),
	}

	for _, protoFile := range p.Req.SourceFileDescriptors {
		protoName := protoFile.GetName()
		protoName = protoName[:len(protoName)-len(".proto")]

		li := strings.LastIndex(protoName, "/")
		if li != -1 {
			protoName = protoName[li+1:]
		}

		spec := doc_handler_template.Spec{
			Name:     snakeToPascal(protoName),
			FileName: protoName,
		}

		genReq.Specs = append(genReq.Specs, spec)
	}

	genReq.PrimarySpecName = genReq.Specs[0].Name

	li := strings.LastIndex(genReq.PrimarySpecName, "/")
	if li != -1 {
		genReq.PrimarySpecName = genReq.PrimarySpecName[li+1:]
	}

	content, err := doc_handler_template.Generate(genReq)
	if err != nil {
		panic(err.Error())
	}
	swaggerUIHandler.Content = toolbox.ToPtr(string(content))

	return &swaggerUIHandler
}

func snakeToPascal(in string) string {
	sb := strings.Builder{}
	lastWordStart := 0

	for idx, char := range []rune(in) {
		switch char {
		case '-', '_':
			if lastWordStart != -1 {
				sb.WriteRune(unicode.ToUpper(rune(in[lastWordStart])))
				sb.WriteString(in[lastWordStart+1 : idx])
				lastWordStart = -1
			}
		default:
			if lastWordStart == -1 {
				lastWordStart = idx
			}
		}
	}
	sb.WriteRune(unicode.ToUpper(rune(in[lastWordStart])))
	sb.WriteString(in[lastWordStart+1:])
	lastWordStart = -1

	return sb.String()
}
