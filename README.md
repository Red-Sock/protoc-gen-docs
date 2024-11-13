# Swagger documentation self-hosting

# Installing
```shell
  go install github.com/Red-Sock/protoc-gen-docs@v0.0.1
```

# Using
You can find copy&paste example at example folder

## Basic usage

```makefile
.gen-server-grpc:
	protoc \
    	-I=./api \
    	-I $(GOPATH)/bin \
    	--openapiv2_out=./pkg/docs/swaggers \
     	--docs_out=./pkg/docs \  
    	./api/grpc/*.proto
```
This script will generate following folder structure

```
.
└── docs
    ├── docs.swagger_ui.go
    └── swagger
        ├── gictionary.swagger.json
        └── hello_world_api.swagger.json
```

### docs_out
Value **MUST** be parent folder for your `openapiv2_out`'s value

if `docs_out` is set to `./pkg/docs` 
then your `openapiv2_out` **MUST** be set to something like `./pkg/docs/swaggers`.
Otherwise, protoc-gen-docs wouldn't be able to find swaggers 

By default all `swagger.json` files hosted at `/docs/swaggers`
e.g. 
```
http://localhost/docs/swaggers/hello_world_v1.swagger.json
```


# Customization
Following script can be used to generate docs handle along with swaggers for your service
```makefile
.gen-server-grpc:
	protoc \
    	-I=./api \
    	-I $(GOPATH)/bin \
    	--openapiv2_out=./pkg/docs/swagger \
     	--docs_out=./pkg/docs \  
    	--docs_opt base_path=non_default_docs_path \
    	--docs_opt swaggers_folder_path=/swagger/ \
    	--docs_opt swaggers_web_path=/swagger/ \
    	--docs_opt tittle="My custom swagger 🤡" \
    	./api/grpc/*.proto
```

## base_path
Can be used to override web path at which swagger web ui will be available

**By default**, is set to `/docs`   

## swaggers_folder_path
Can be used to override folder path to swagger files, BUT required to be folder 
BELOW `docs_out` due to golang's embeding rules

**By default**, is set to `swaggers`

## swaggers_web_path
Can be used to override web path at which swaggers json files will be hosted

This value will be appended to `base_path` 

**By default**, is set to `swaggers`

## tittle
Can be used to override web UI title

**By default**, is set to `Swagger`