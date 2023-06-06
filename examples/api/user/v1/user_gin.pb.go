// Code generated by protoc-gen-go-gin. DO NOT EDIT."

package v1

import (
	context "context"
	errors "errors"
	gin "github.com/gin-gonic/gin"
	metadata "google.golang.org/grpc/metadata"
	http "net/http"
)

// context.
// gin.
// errors.
// http.
// metadata.

type UserHTTPServer interface {
	GetInfo(context.Context, *UserRequest) (*UserReply, error)
	GetInfo2(context.Context, *UserRequest) (*UserReply, error)
	GetInfo3(context.Context, *UserRequest) (*UserReply, error)
}

type defaultUserResp struct{}

func (resp defaultUserResp) Response(ctx *gin.Context, httpStatus int, code int, message string, data interface{}) {
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

func (resp defaultUserResp) Error(ctx *gin.Context, err error) {
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
func (resp defaultUserResp) ParamsError(ctx *gin.Context, err error) {
	_ = ctx.Error(err)
	resp.Response(ctx, 400, 40001, err.Error(), nil)
}

func (resp defaultUserResp) Success(ctx *gin.Context, data interface{}) {
	resp.Response(ctx, http.StatusOK, 10000, "OK", data)
}

type UserController struct {
	service UserHTTPServer
	router  gin.IRouter
	resp    interface {
		Error(ctx *gin.Context, err error)
		ParamsError(ctx *gin.Context, err error)
		Success(ctx *gin.Context, data interface{})
	}
}

func (c *UserController) GetInfo(ctx *gin.Context) {
	var in UserRequest

	if err := ctx.ShouldBindQuery(&in); err != nil {
		c.resp.ParamsError(ctx, err)
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)

	out, err := c.service.GetInfo(newCtx, &in)
	if err != nil {
		c.resp.Error(ctx, err)
		return
	}
	c.resp.Success(ctx, out)
}

func (c *UserController) GetInfo2(ctx *gin.Context) {
	var in UserRequest

	if err := ctx.ShouldBindUri(&in); err != nil {
		c.resp.ParamsError(ctx, err)
		return
	}

	if err := ctx.ShouldBindQuery(&in); err != nil {
		c.resp.ParamsError(ctx, err)
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)

	out, err := c.service.GetInfo2(newCtx, &in)
	if err != nil {
		c.resp.Error(ctx, err)
		return
	}
	c.resp.Success(ctx, out)
}

func (c *UserController) GetInfo3(ctx *gin.Context) {
	var in UserRequest

	if err := ctx.ShouldBindUri(&in); err != nil {
		c.resp.ParamsError(ctx, err)
		return
	}

	if err := ctx.ShouldBindJSON(&in); err != nil {
		c.resp.ParamsError(ctx, err)
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)

	out, err := c.service.GetInfo3(newCtx, &in)
	if err != nil {
		c.resp.Error(ctx, err)
		return
	}
	c.resp.Success(ctx, out)
}

func (c *UserController) RegisterService() {
	c.router.Handle("GET", "/v1/user/info", c.GetInfo)
	c.router.Handle("GET", "/v1/user/info/:id/*action", c.GetInfo2)
	c.router.Handle("POST", "/v1/user/info/:id", c.GetInfo3)
}

func RegisterUserHTTPServer(router gin.IRouter, srv UserHTTPServer) {
	c := &UserController{
		service: srv,
		router:  router,
		resp:    defaultUserResp{},
	}
	c.RegisterService()
}
