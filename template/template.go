package template

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"strings"
)

//go:embed template.go.tpl
var serviceTpl string

const interfaceSuffix = "HTTPServer"

type ServiceTemplate struct {
	Name string
	// FullName string
	// FilePath string
	Methods   []*MethodTemplate
	MethodSet map[string]*MethodTemplate
}

func (s *ServiceTemplate) Interface() string {
	return fmt.Sprintf("%s%s", s.Name, interfaceSuffix)
}

func (s *ServiceTemplate) AddMethod(method *MethodTemplate) {
	s.Methods = append(s.Methods, method)
}

func (s *ServiceTemplate) String() string {
	tmpl, err := template.New("http").Parse(strings.TrimSpace(serviceTpl))
	if err != nil {
		panic(err)
	}
	buffer := new(bytes.Buffer)
	if err := tmpl.Execute(buffer, s); err != nil {
		panic(err)
	}
	return buffer.String()
}

type MethodTemplate struct {
	Name         string
	Num          int
	Request      string
	Reply        string
	Path         string
	Method       string
	Body         string
	ResponseBody string
}
