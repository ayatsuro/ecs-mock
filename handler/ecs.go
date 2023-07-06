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

func GetNs(ctx *gin.Context) {
	name := ctx.Param("item")
	name = strings.TrimSuffix(name, ".json")
	item, ok := service.GetNs(name)
	if !ok {
		ctx.AbortWithStatusJSON(404, gin.H{"error": "namespace not found"})
		return
	}
	ctx.JSON(200, item)
}

func IAMAction(ctx *gin.Context) {
	action := ctx.Query("Action")
	ns := ctx.Request.Header.Get("x-emc-namespace")
	if ns == "" {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "missing namespace header"})
		return
	}
	_, ok := service.GetNs(ns)
	if !ok {
		ctx.AbortWithStatusJSON(404, gin.H{"error": "namespace not found"})
		return
	}
	if action == "ListUsers" {
		users := service.ListIamUsers(ns)
		output := model.ListIamUsers{ListUsersResult: model.Users{Users: users}}
		ctx.JSON(200, output)
	}
	if action == "GetUser" {
		user, code := service.GetIamUser(ns, ctx.Query("UserName"))
		if code != 200 {
			ctx.AbortWithStatusJSON(404, gin.H{"error": "iam user not found"})
			return
		}
		ctx.JSON(200, user)
	}
	if action == "CreateUser" {
		username := ctx.Query("UserName")
		user, code := service.CreateIamUser(ns, username)
		if code == 409 {
			ctx.AbortWithStatusJSON(409, gin.H{"error": "iam user exists already"})
			return
		}
		ctx.JSON(200, user)
	}
	if action == "DeleteUser" {
		username := ctx.Query("UserName")
		code := service.DeleteIamUser(ns, username)
		if code != 200 {
			ctx.AbortWithStatusJSON(404, gin.H{"error": "iam user not found"})
			return
		}
		ctx.Status(200)
	}
	if action == "CreateAccessKey" {
		username := ctx.Query("UserName")
		key, code := service.CreateAccessKey(ns, username)
		if code == 404 {
			ctx.AbortWithStatusJSON(404, gin.H{"error": "user not found"})
			return
		}
		if code == 409 {
			ctx.AbortWithStatusJSON(409, gin.H{"error": "max number of access keys reached"})
			return
		}
		output := model.CreateAccessKey{
			CreateAccessKeyResult: model.CreateAccessKeyResult{
				AccessKey: key,
			},
		}
		ctx.JSON(200, output)
	}
	if action == "ListAccessKeys" {
		username := ctx.Query("UserName")
		keys, code := service.ListAccessKeys(ns, username)
		if code != 200 {
			ctx.AbortWithStatusJSON(404, gin.H{"error": "user not found"})
			return
		}
		output := model.ListAccessKeys{ListAccessKeysResult: model.AccessKeyMetadata{AccessKeys: keys}}
		ctx.JSON(200, output)
	}
	if action == "DeleteAccessKey" {
		username := ctx.Query("UserName")
		keyId := ctx.Query("AccessKeyId")
		if code := service.DeleteAccessKey(ns, username, keyId); code != 200 {
			ctx.AbortWithStatusJSON(404, gin.H{"error": "user or access key not found"})
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
