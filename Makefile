goinstall:
	go install

build: goinstall
	protoc -I ./examples/api \
		--go_out ./examples/api --go_opt=paths=source_relative \
		--go-gin_out ./examples/api --go-gin_opt=paths=source_relative ./examples/api/user/v1/user.proto
	protoc-go-inject-tag -input=./examples/api/user/v1/user.pb.go