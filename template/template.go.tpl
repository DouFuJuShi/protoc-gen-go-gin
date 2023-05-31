type {{$.Interface}} interface {
{{range .Methods}}
    {{.Name}}(context.Context{{if .Request}}, *{{.Request}}{{end}}) {{if .Reply}}(*{{.Reply}}, error){{else}}error{{end}}
{{- end}}
}

type default{{$.Name}}Resp struct {}

func (resp default{{$.Name}}Resp) Response(ctx *gin.Context, httpStatus int, code int, message string, data interface{}) {
	body := &struct {
                     Code    int         `json:"code"`
                     Message string      `json:"msg"`
                     Data    interface{} `json:"data"`
                 }{
		Code:    code,
		Message: message,
		Data:    data,
	}
	ctx.PureJSON(httpStatus, body)
}

func (resp default{{$.Name}}Resp) Error(ctx *gin.Context, err error) {
	httpStatus := 500
	code := -1
	message := "未知错误"

	type iResponse interface {
		HttpStatus() int
		Code() int
		Message() string
	}

	var iResp iResponse
	if err != nil {
		if errors.As(err, &iResp) {
			httpStatus = iResp.HttpStatus()
			code = iResp.Code()
			message = iResp.Message()
		} else {
			message += ";" + err.Error()
		}
	} else {
		message += "; err is nil"
	}
    _ = ctx.Error(err)
	resp.Response(ctx, httpStatus, code, message, nil)
}

// ParamsError 参数错误
func (resp default{{$.Name}}Resp) ParamsError (ctx *gin.Context, err error) {
	_ = ctx.Error(err)
	resp.Response(ctx, 400, 40001, err.Error(), nil)
}

func (resp default{{$.Name}}Resp) Success(ctx *gin.Context, data interface{}) {
	resp.Response(ctx, http.StatusOK, 10000, "OK", nil)
}

type {{$.Name}}Controller struct {
	service {{$.Interface}}
	router gin.IRouter
	resp   interface {
		Error(ctx *gin.Context, err error)
		ParamsError (ctx *gin.Context, err error)
		Success(ctx *gin.Context, data interface{})
	}
}

{{- range .Methods}}
func (c *{{$.Name}}Controller) {{.Name}}(ctx *gin.Context) {
{{- if .Request}}
     var in {{.Request}}
    {{if .ShouldBindUri }}
    if err := ctx.ShouldBindUri(&in); err != nil {
        c.resp.ParamsError(ctx, err)
        return
    }
    {{end}}
    {{if eq .HttpMethod "GET" "DELETE" }}
        if err := ctx.ShouldBindQuery(&in); err != nil {
            c.resp.ParamsError(ctx, err)
            return
        }
    {{else if eq .HttpMethod "POST" "PUT" }}
        if err := ctx.ShouldBindJSON(&in); err != nil {
            c.resp.ParamsError(ctx, err)
            return
        }
    {{else}}
        if err := ctx.ShouldBind(&in); err != nil {
            c.resp.ParamsError(ctx, err)
            return
        }
    {{end}}
{{end}}
    md := metadata.New(nil)
    for k, v := range ctx.Request.Header {
        md.Set(k, v...)
    }
    newCtx := metadata.NewIncomingContext(ctx, md)

{{if .Reply}}
    out, err := c.service.{{.Name}}(newCtx{{if .Request}}, &in{{end}})
    if err != nil {
        c.resp.ParamsError(ctx, err)
        return
    }
    c.resp.Success(ctx, out)
{{- else}}
    err := c.service.{{.Name}}(newCtx{{if .Request}}, &in{{end}})
    if err != nil {
        c.resp.ParamsError(ctx, err)
        return
    }
    c.resp.Success(ctx, nil)
{{- end}}
}
{{end}}

func (c *{{$.Name}}Controller) RegisterService() {
{{- range .Methods}}
		c.router.Handle("{{.HttpMethod}}", "{{.Path}}", c.{{.Name}})
{{- end}}
}

func Register{{$.Interface}}(router gin.IRouter, srv {{$.Interface}}) {
    c := &{{$.Name}}Controller{
        service: srv,
        router: router,
        resp: default{{$.Name}}Resp{},
    }
    c.RegisterService()
}