package main

import (
	"flag"
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

const version = "0.0.1"

func main() {
	versionFlag := flag.Bool("version", false, "print the version and exit")
	flag.Parse()
	if *versionFlag {
		fmt.Println("protoc-gen-go-gin")
		fmt.Println(fmt.Sprintf("current version: %v", version))
		return
	}

	var flags flag.FlagSet

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			_ = NewFileGenerator(gen, f).Exec()
		}
		return nil
	})
}
