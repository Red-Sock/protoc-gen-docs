package request

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParameters(t *testing.T) {
	t.Run("Default", func(t *testing.T) {
		ip := ParseInputParams("base_path=non_default_docs_path,swaggers_folder_path=/swg/,swaggers_web_path=/swc/")
		expected := InputParams{
			BasePath:          "non_default_docs_path",
			SwaggerFolderPath: "swg",
			SwaggerWebPath:    "swc",
			Title:             "Swagger",
			SourceRelative:    false,
		}
		require.Equal(t, expected, ip)
	})

	t.Run("SourceRelative", func(t *testing.T) {
		ip := ParseInputParams("paths=source_relative")
		require.True(t, ip.SourceRelative)
	})
}
