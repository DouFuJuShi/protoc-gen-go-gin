type {{$.Interface}} interface {
    {{range .Methods}}
    {{.Name}}(context.Context, *{{.Request}}) (*{{.Reply}}, error)
    {{end}}
}

func Register{{$.Interface}}(router gin.IRouter, srv {{$.Interface}}) {
    router.
}