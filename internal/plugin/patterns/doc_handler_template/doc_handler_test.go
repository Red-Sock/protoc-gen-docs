package doc_handler_template

import (
	"go/format"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerate_OutputIsGofmtClean(t *testing.T) {
	req := SwaggerUIGenReq{
		BasePath:        "docs",
		SwaggerWebPath:  "/docs",
		SwaggerFolder:   "swaggers",
		Tittle:          "Swagger",
		PrimarySpecName: "Hello",
		Specs:           []Spec{{Name: "Hello", FileName: "hello"}},
	}

	got, err := Generate(req)
	require.NoError(t, err)

	want, err := format.Source(got)
	require.NoError(t, err)
	require.Equal(t, string(want), string(got))
}
