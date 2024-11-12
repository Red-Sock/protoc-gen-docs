gen: .pre-gen-server-grpc .deps-grpc .gen-server-grpc

.pre-gen-server-grpc:
	rm -rf pkg/
	rm -rf pkg/docs
	rm -rf pkg/docs/swaggers
	rm -rf pkg/web

	mkdir -p pkg/
	mkdir -p pkg/docs
	mkdir -p pkg/docs/swaggers
	mkdir -p pkg/web

.deps-grpc:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest

	rm -rf api/google
	rm -rf api/validate
	ln -sf $(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis/google api/google
	ln -sf $(GOPATH)/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v1.0.2/validate api/validate


.gen-server-grpc:
	protoc \
    	-I=./api/grpc \
    	-I=./api \
    	-I $(GOPATH)/bin \
    	--plugin=protoc-gen-docs=bin/protoc-gen-docs \
    	--openapiv2_out=./pkg/docs/swaggers \
    	--docs_out=./pkg/docs \
    	--docs_opt base_path=/swagger/* \
    	./api/grpc/*.proto