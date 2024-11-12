package doc_handler_template

import (
	"bytes"
	_ "embed"
	"text/template"

	errors "github.com/Red-Sock/trace-errors"
)

//go:embed doc_handler.go.pattern
var byteTemplate string
var swaggerHandleTemplate = template.Must(template.New("swagger_handle").Parse(byteTemplate))

type SwaggerUIGenReq struct {
	BasePath       string
	SwaggerWebPath string
	SwaggerFolder  string

	Tittle          string
	PrimarySpecName string
	Specs           []Spec
}

type Spec struct {
	Name     string
	FileName string
}

func Generate(req SwaggerUIGenReq) ([]byte, error) {
	buf := bytes.NewBuffer(nil)

	err := swaggerHandleTemplate.Execute(buf, req)
	if err != nil {
		return nil, errors.Wrap(err, "error generating")
	}

	return buf.Bytes(), nil
}
