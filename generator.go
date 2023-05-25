package main

import (
	_ "embed"
	"fmt"
	"github.com/DouFuJuShi/protoc-gen-go-gin/template"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"net/http"
)

const (
	outSuffix = "_gin.pb.go"
	ctxPKG    = protogen.GoImportPath("context")
	ginPKG    = protogen.GoImportPath("github.com/gin-gonic/gin")
	errsPKG   = protogen.GoImportPath("errors")
	metaPKG   = protogen.GoImportPath("google.golang.org/grpc/metadata")
	httpPKG   = protogen.GoImportPath("net/http")
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
		httpRule, ok := proto.GetExtension(m.Desc.Options(), annotations.E_Http).(*annotations.HttpRule)
		if !ok || httpRule == nil {
			continue
		}

		var (
			path   string
			method string
		)
		switch pattern := httpRule.Pattern.(type) {
		case *annotations.HttpRule_Get:
			path = pattern.Get
			method = http.MethodGet
		case *annotations.HttpRule_Put:
			path = pattern.Put
			method = http.MethodPut
		case *annotations.HttpRule_Post:
			path = pattern.Post
			method = http.MethodPost
		case *annotations.HttpRule_Delete:
			path = pattern.Delete
			method = http.MethodDelete
		case *annotations.HttpRule_Patch:
			path = pattern.Patch
			method = http.MethodPatch
		case *annotations.HttpRule_Custom:
			path = pattern.Custom.Path
			method = pattern.Custom.Kind
		}

		st.AddMethod(&template.MethodTemplate{
			Name:       m.GoName,
			Request:    m.Input.GoIdent.GoName,
			Reply:      m.Output.GoIdent.GoName,
			Path:       path,
			HttpMethod: method,
		})
	}

	g.generatedFile.P(st.String())
}

func (g FileGenerator) before() {
	g.generatedFile.P("// Code generated by protoc-gen-go-gin. DO NOT EDIT.\"")
	g.generatedFile.P()
	g.generatedFile.P("package ", g.ProtoFile.GoPackageName)
	g.generatedFile.P()

	// 这里是为了程序编译的时候确保这些包是正确的
	g.generatedFile.P("// ", ctxPKG.Ident(""))
	g.generatedFile.P("// ", ginPKG.Ident(""))
	g.generatedFile.P("// ", errsPKG.Ident(""))
	g.generatedFile.P("// ", httpPKG.Ident(""))
	g.generatedFile.P("// ", metaPKG.Ident(""))
	g.generatedFile.P()
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
