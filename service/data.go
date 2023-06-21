package service

import (
	"ecs-mock/model"
	"math/rand"
	"time"
)

var (
	data    map[string]*model.Namespace
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func InitData() {
	rand.Seed(time.Now().UnixNano())
	data = make(map[string]*model.Namespace)
	data["ci12345-native-user"] = &model.Namespace{
		Name:        "ci12345-native-user",
		NativeUsers: []model.NativeUser{{Userid: "nativeUser1", Name: "nativeUser1"}},
	}
	data["ci45678-native-user-iam-user-1key"] = &model.Namespace{
		Name:        "ci45678-native-user-iam-user-1key",
		NativeUsers: []model.NativeUser{{Userid: "nativeUser2", Name: "nativeUser2"}},
		IamUsers:    []*model.IamUser{{UserName: "iamUser1", AccessKeys: []model.AccessKey{{AccessKeyId: "accessKeyId1", SecretAccessKey: "secretAccessKey1"}}}},
	}
	data["ci898640-native-user-iam-user-2keys"] = &model.Namespace{
		Name:        "ci898640-native-user-iam-user-2keys",
		NativeUsers: []model.NativeUser{{Userid: "nativeUser3", Name: "nativeUser3"}},
		IamUsers:    []*model.IamUser{{UserName: "iamUser2", AccessKeys: []model.AccessKey{{AccessKeyId: "accessKeyId2", SecretAccessKey: "secretAccessKey2"}, {AccessKeyId: "accessKeyId3", SecretAccessKey: "secretAccessKey3"}}}},
	}
}

func ListNs() []model.Namespace {
	var ns []model.Namespace
	for _, n := range data {
		ns = append(ns, *n)
	}
	return ns
}

func GetNs(name string) (*model.Namespace, bool) {
	ns, ok := data[name]
	return ns, ok
}

func ListNativeUsers(namespace string) ([]model.NativeUser, bool) {
	ns, ok := data[namespace]
	if !ok {
		return nil, false
	}
	return ns.NativeUsers, true
}

func GetNativeUser(uid string) (model.NativeUser, bool) {
	for _, ns := range data {
		for _, u := range ns.NativeUsers {
			if u.Userid == uid {
				return u, true
			}
		}
	}
	return model.NativeUser{}, false
}

func ListIamUsers(namespace string) []model.IamUser {
	ns := data[namespace]
	var users []model.IamUser
	for _, u := range ns.IamUsers {
		users = append(users, *u)
	}
	return users
}

func GetIamUser(namespace, username string) (*model.IamUser, int) {
	ns := data[namespace]
	for _, u := range ns.IamUsers {
		if u.UserName == username {
			return u, 200
		}
	}
	return nil, 404
}

func CreateIamUser(namespace, username string) (model.IamUser, int) {
	ns := data[namespace]
	found := false
	for _, u := range ns.IamUsers {
		if u.UserName == username {
			found = true
			break
		}
	}
	if found {
		return model.IamUser{}, 409
	}
	user := model.IamUser{UserName: username}
	ns.IamUsers = append(ns.IamUsers, &user)
	return user, 200
}

func DeleteIamUser(namespace, username string) int {
	ns := data[namespace]
	var keep []*model.IamUser
	for _, u := range ns.IamUsers {
		if u.UserName != username {
			keep = append(keep, u)
		}
	}
	if len(keep) == len(ns.IamUsers) {
		return 404
	}
	ns.IamUsers = keep
	return 200
}

func CreateAccessKey(namespace, username string) (model.AccessKey, int) {
	key := model.AccessKey{}
	ns := data[namespace]
	var user *model.IamUser
	for _, u := range ns.IamUsers {
		if u.UserName == username {
			user = u
			break
		}
	}
	if user == nil {
		return key, 404
	}
	if len(user.AccessKeys) == 2 {
		return key, 409
	}
	key.UserName = username
	key.AccessKeyId = RandString(8)
	key.SecretAccessKey = RandString(16)
	key.CreateDate = time.Now().Format(time.RFC3339)
	user.AccessKeys = append(user.AccessKeys, key)
	return key, 200
}

func DeleteAccessKey(namespace, username, keyId string) int {
	ns := data[namespace]
	var user *model.IamUser
	for _, u := range ns.IamUsers {
		if u.UserName == username {
			user = u
			break
		}
	}
	if user == nil {
		return 404
	}
	code := 404
	var keep []model.AccessKey
	for _, k := range user.AccessKeys {
		if k.AccessKeyId == keyId {
			code = 200
			continue
		}
		keep = append(keep, k)
	}
	user.AccessKeys = keep
	return code

}

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func ListAccessKeys(namespace, username string) ([]model.AccessKey, int) {
	ns := data[namespace]
	var user *model.IamUser
	for _, u := range ns.IamUsers {
		if u.UserName == username {
			user = u
			break
		}
	}
	if user == nil {
		return nil, 404
	}
	return user.AccessKeys, 200
}
