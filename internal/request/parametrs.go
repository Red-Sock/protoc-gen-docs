package request

import (
	"strings"
)

type InputParams struct {
	BasePath string
}

const optPrefix = "--docs_opts"

func ParseInputParams(params string) InputParams {
	ip := InputParams{
		BasePath: "docs",
	}
	for _, arg := range strings.Split(params, "\n") {
		if !strings.HasPrefix(arg, optPrefix) {
			continue
		}

		param := arg[strings.Index(arg, " "):]
		paramName := param[:strings.Index(param, "=")]

		switch paramName {
		case "base_path":
			ip.BasePath = strings.TrimRight(strings.TrimLeft(param[len(paramName)+2:], "/"), "/")
		}
	}

	return ip
}
