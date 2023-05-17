type {{$.Interface}} interface {
{{range .Methods}}
    {{.Name}}(context.Context, *{{.Request}}) (*{{.Reply}}, error)
{{- end}}
}

func Register{{$.Interface}}(router gin.IRouter, srv {{$.Interface}}) {
{{- range .Methods}}
    router.{{.HttpMethod}}("{{.Path}}", func(ctx *gin.Context) {
        var in {{.Request}}
    {{if .ShouldBindUri }}
        if err := ctx.ShouldBindUri(&in); err != nil {
            return
        }
    {{end}}
    {{if eq .HttpMethod "GET" "DELETE" }}
        if err := ctx.ShouldBindQuery(&in); err != nil {
            return
        }
    {{else if eq .HttpMethod "POST" "PUT" }}
        if err := ctx.ShouldBindJSON(&in); err != nil {
            return
        }
    {{else}}
        if err := ctx.ShouldBind(&in); err != nil {
            return
        }
    {{end}}

    md := metadata.New(nil)
    for k, v := range ctx.Request.Header {
        md.Set(k, v...)
    }
    newCtx := metadata.NewIncomingContext(ctx, md)

    
    out, err := srv.{{.Name}}(newCtx, &in)
    if err != nil {
        return
    }

    ctx.PureJSON(http.StatusOK, out)
    })
{{end}}
}