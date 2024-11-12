package request

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParameters(t *testing.T) {
	ip := ParseInputParams("base_path=non_default_docs_path,swaggers_folder_path=/swg/,swaggers_web_path=/swc/")
	expected := InputParams{
		BasePath:          "non_default_docs_path",
		SwaggerFolderPath: "swg",
		SwaggerWebPath:    "swc",
	}
	require.Equal(t, expected, ip)
}
