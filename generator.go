package main

import (
	_ "embed"
	"fmt"
	"github.com/DouFuJuShi/protoc-gen-go-gin/template"
	"google.golang.org/protobuf/compiler/protogen"
)

const (
	outSuffix = "_gin.pb.go"
	ctxPKG    = protogen.GoImportPath("context")
	ginPKG    = protogen.GoImportPath("github.com/gin-gonic/gin")
	errsPKG   = protogen.GoImportPath("errors")
	metaPKG   = protogen.GoImportPath("google.golang.org/grpc/metadata")
)

type FileGenerator struct {
	Plugin        *protogen.Plugin
	ProtoFile     *protogen.File
	generatedFile *protogen.GeneratedFile
}

func (g FileGenerator) Exec() error {
	g.before()

	if len(g.ProtoFile.Services) == 0 {
		return nil
	}

	for _, s := range g.ProtoFile.Services {
		g.genService(s)
	}

	g.after()
	return nil
}

func (g FileGenerator) genService(s *protogen.Service) {
	st := template.ServiceTemplate{
		Name: s.GoName,
	}

	for _, m := range s.Methods {
		st.AddMethod(&template.MethodTemplate{
			Name:    m.GoName,
			Request: m.Input.GoIdent.GoName,
			Reply:   m.Output.GoIdent.GoName,
		})
	}

	// st.AddMethod()
	g.generatedFile.P(st.String())
}

func (g FileGenerator) before() {
	g.generatedFile.P("// Code generated by protoc-gen-go-gin. DO NOT EDIT.\"")
	g.generatedFile.P()
	g.generatedFile.P("package ", g.ProtoFile.GoPackageName)
	g.generatedFile.P()
	g.generatedFile.P("// ", ctxPKG.Ident(""), metaPKG.Ident(""))
	g.generatedFile.P("// ", ginPKG.Ident(""), errsPKG.Ident(""))
}

func (g FileGenerator) after() {
	// TODO
}

func NewFileGenerator(p *protogen.Plugin, f *protogen.File) FileGenerator {
	outFile := fmt.Sprintf("%s%s", f.GeneratedFilenamePrefix, outSuffix)
	gf := p.NewGeneratedFile(outFile, f.GoImportPath)

	return FileGenerator{
		Plugin:        p,
		ProtoFile:     f,
		generatedFile: gf,
	}
}
