package request

import (
	"strings"
)

type InputParams struct {
	BasePath          string
	SwaggerFolderPath string
	SwaggerWebPath    string
	Title             string
}

const (
	optPrefix = "docs_opts"

	basePathOption          = "base_path"
	swaggerFolderPathOption = "swaggers_folder_path"
	swaggerWebPathOption    = "swaggers_web_path"
	tittleOption            = "tittle"
)

func ParseInputParams(params string) InputParams {
	ip := InputParams{
		BasePath:          "docs",
		SwaggerFolderPath: "swaggers",
		SwaggerWebPath:    "swaggers",
		Title:             "Swagger",
	}
	for _, param := range strings.Split(params, ",") {
		paramName := param[:strings.Index(param, "=")]
		paramValue := param[len(paramName)+1:]

		switch paramName {
		case basePathOption:
			ip.BasePath = strings.TrimRight(strings.TrimLeft(paramValue, "/"), "/")
		case swaggerFolderPathOption:
			ip.SwaggerFolderPath = strings.TrimRight(strings.TrimLeft(paramValue, "/"), "/")
		case swaggerWebPathOption:
			ip.SwaggerWebPath = strings.TrimRight(strings.TrimLeft(paramValue, "/"), "/")
		case tittleOption:
			ip.Title = paramValue
		}
	}

	return ip
}
