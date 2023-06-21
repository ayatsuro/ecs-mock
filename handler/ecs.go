package handler

import (
	"ecs-mock/model"
	"ecs-mock/service"
	"github.com/gin-gonic/gin"
	"strings"
)

func ListNs(ctx *gin.Context) {
	items := service.ListNs()
	output := model.Namespaces{Namespace: items}
	ctx.JSON(200, output)
}

func ListNativeUsers(ctx *gin.Context) {
	name := ctx.Param("item")
	name = strings.TrimSuffix(name, ".json")
	users, ok := service.ListNativeUsers(name)
	if !ok {
		ctx.AbortWithStatus(404)
		return
	}
	output := model.NativeUsers{Users: users}
	ctx.JSON(200, output)
}

func GetNativeUser(ctx *gin.Context) {
	uid := ctx.Param("item")
	user, ok := service.GetNativeUser(uid)
	if !ok {
		ctx.AbortWithStatus(404)
		return
	}
	ctx.JSON(200, user)
}

func IAMAction(ctx *gin.Context) {
	action := ctx.Query("Action")
	ns := ctx.Request.Header.Get("x-emc-namespace")
	if ns == "" {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "missing namespace header"})
		return
	}
	if action == "ListUsers" {
		users, ok := service.ListIamUsers(ns)
		if !ok {
			ctx.AbortWithStatus(404)
			return
		}
		output := model.ListIamUsers{ListUsersResult: model.Users{Users: users}}
		ctx.JSON(200, output)
		return
	}
	if action == "CreateUser" {
		username := ctx.Query("UserName")
		user, code := service.CreateIamUser(ns, username)
		if code != 200 {
			ctx.AbortWithStatus(code)
			return
		}
		ctx.JSON(200, user)
		return
	}
	if action == "CreateAccessKey" {
		username := ctx.Query("UserName")
		key, code := service.CreateAccessKey(ns, username)
		if code != 200 {
			ctx.AbortWithStatus(code)
			return
		}
		output := model.CreateAccessKey{
			CreateAccessKeyResult: model.CreateAccessKeyResult{
				AccessKey: key,
			},
		}
		ctx.JSON(200, output)
		return
	}
	if action == "DeleteAccessKey" {
		username := ctx.Query("UserName")
		keyId := ctx.Query("AccessKeyId")
		if code := service.DeleteAccessKey(ns, username, keyId); code != 200 {
			ctx.AbortWithStatus(code)
			return
		}
		ctx.Status(200)
	}
}

func UpdateVdcUser(ctx *gin.Context) {
	username := ctx.Param("item")
	username = strings.TrimSuffix(username, ".json")
	var vdcUser model.VdcUser
	if err := ctx.BindJSON(&vdcUser); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	service.WriteAdminPwd(username, vdcUser)
}
