package docs

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"path"

	swaggerui "github.com/Red-Sock/go-swagger-ui"
)

//go:embed all:swagger
var swaggers embed.FS

const (
	BasePath    = "/non_default_docs_path/"
	swaggerPath = BasePath + "swagger/"
)

func Swagger() (p string, handler http.HandlerFunc) {
	mux := http.NewServeMux()

	mux.Handle(BasePath, swaggerui.NewHandler(
		swaggerui.WithBasePath(BasePath),
		swaggerui.WithHTMLTitle("My custom swagger 🤡"),
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

	stripped, err := fs.Sub(swaggers, "swagger")
	if err != nil {
		log.Fatal(err)
	}

	ffs := http.StripPrefix(swaggerPath, http.FileServer(http.FS(stripped)))
	mux.Handle(swaggerPath, ffs)

	return BasePath, mux.ServeHTTP
}
