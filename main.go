package main

import (
	"flag"
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
)

const version = "0.0.4"

func main() {
	versionFlag := flag.Bool("version", false, "print the version and exit")
	flag.Parse()
	if *versionFlag {
		fmt.Println("protoc-gen-go-gin")
		fmt.Printf("current version: %v\n", version)
		return
	}

	var flags flag.FlagSet

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(Handler())
}
