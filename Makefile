build:
	go install
	protoc -I ./examples/api \
		--go_out ./examples/api --go_opt=paths=source_relative \
		--go-gin_out ./examples/api --go-gin_opt=paths=source_relative ./examples/api/hello/v1/hello.proto
	protoc-go-inject-tag -input=./examples/api/hello/v1/hello.pb.go