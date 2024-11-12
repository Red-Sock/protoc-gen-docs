package docs

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"path"

	swaggerui "github.com/Red-Sock/go-swagger-ui"
)

//go:embed all:swaggers
var swaggers embed.FS

const (
	BasePath    = "/docs/"
	swaggerPath = BasePath + "swaggers/"
)

func Swagger() (p string, handler http.HandlerFunc) {
	mux := http.NewServeMux()

	mux.Handle(BasePath, swaggerui.NewHandler(
		swaggerui.WithBasePath(BasePath),
		swaggerui.WithHTMLTitle("Service Docs"),
		swaggerui.WithSpecURLs("Gictionary",
			[]swaggerui.SpecURL{
				{
					Name: "Gictionary",
					URL:  path.Join(swaggerPath, "gictionary.swagger.json"),
				},
				{
					Name: "HelloWorldApi",
					URL:  path.Join(swaggerPath, "hello_world_api.swagger.json"),
				},
			}),
		swaggerui.WithShowExtensions(true),
	))

	stripped, err := fs.Sub(swaggers, "swaggers")
	if err != nil {
		log.Fatal(err)
	}

	ffs := http.StripPrefix(swaggerPath, http.FileServer(http.FS(stripped)))
	mux.Handle(swaggerPath, ffs)

	return BasePath, mux.ServeHTTP
}
