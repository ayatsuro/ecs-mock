package handler

import (
	"ecs-mock/model"
	"github.com/gin-gonic/gin"
)

func ListNs(ctx *gin.Context) {
	items := model.Namespaces{Namespace: []model.Namespace{{Name: "again"}, {Name: "again2"}}}
	ctx.JSON(200, items)
}

func ListNativeUsers(ctx *gin.Context) {
	users := model.NativeUsers{Users: []model.NativeUser{{Userid: "user1"}, {Userid: "user2"}}}
	ctx.JSON(200, users)
}

func GetNativeUser(ctx *gin.Context) {
	user := model.NativeUser{}
	if ctx.Param("userId") == "user1" {
		user.Name = "john"
	} else {
		user.Name = "johnny"
	}
	ctx.JSON(200, user)
}

func IAMAction(ctx *gin.Context) {
	action := ctx.Query("Action")
	if action == "ListUsers" {
		users := model.IamUsers{ListUsersResult: model.Users{[]model.IamUser{{UserName: "aimUser1"}, {UserName: "aimUser2"}}}}
		ctx.JSON(200, users)
		return
	}
	if action == "ListAccessKeys" {
		username := ctx.Query("UserName")
		var keys model.AccessKeys
		if username == "aimUser1" {
			keys = model.AccessKeys{ListAccessKeysResult: model.ListAccessKeysResult{AccessKeyMetadata: []model.AccessKey{{AccessKeyId: "123"}}}}
		} else {
			keys = model.AccessKeys{ListAccessKeysResult: model.ListAccessKeysResult{AccessKeyMetadata: []model.AccessKey{{AccessKeyId: "123"}, {AccessKeyId: "456"}}}}
		}
		ctx.JSON(200, keys)
		return
	}
	if action == "CreateAccessKey" {
		username := ctx.Query("UserName")
		key := model.CreateAccessKey{
			CreateAccessKeyResult: model.CreateAccessKeyResult{
				AccessKey: model.AccessKey{
					UserName:        username,
					AccessKeyId:     "123",
					SecretAccessKey: "456",
				},
			},
		}
		ctx.JSON(200, key)
		return
	}
}
