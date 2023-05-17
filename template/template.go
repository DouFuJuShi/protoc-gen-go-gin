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

type MethodTemplate struct {
	Name         string
	Num          int
	Request      string
	Reply        string
	Path         string
	HttpMethod   string
	Body         string
	ResponseBody string
}

func (m MethodTemplate) ShouldBindUri() bool {
	params := strings.Split(m.Path, "/")
	for _, p := range params {
		if len(p) > 0 && (p[0] == ':' || p[0] == '*') {
			return true
		}
	}
	return false
}

type ServiceTemplate struct {
	Name      string
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
