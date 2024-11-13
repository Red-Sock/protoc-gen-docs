build-n-gen: build generate-proto

build:
	mkdir -p example/bin
	go build -o example/bin/protoc-gen-docs main.go

generate-proto:
	cd example && make gen