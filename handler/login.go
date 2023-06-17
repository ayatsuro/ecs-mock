package handler

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
	"strings"
)

const (
	AuthHeader = "X-Sds-Auth-Token"
	MockToken  = "blah"
)

func Login(ctx *gin.Context) {
	var basic string
	for k, v := range ctx.Request.Header {
		if k == "Authorization" {
			basic = v[0]
		}
	}
	if basic == "" {
		ctx.AbortWithStatus(403)
		return
	}
	basic = strings.TrimPrefix(basic, "Basic ")
	auth, err := base64.StdEncoding.DecodeString(basic)
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}
	username, password, found := strings.Cut(string(auth), ":")
	if !found {
		ctx.AbortWithStatus(403)
		return
	}
	if username != "hello" || password != "world" {
		ctx.AbortWithStatus(403)
		return
	}
	slog.Info("logged in")
	ctx.Header(AuthHeader, MockToken)
}
