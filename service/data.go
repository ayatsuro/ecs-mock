package service

import (
	"ecs-mock/model"
	"math/rand"
	"time"
)

var (
	data    map[string]model.Namespace
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func InitData() {
	rand.Seed(time.Now().UnixNano())
	data = make(map[string]model.Namespace)
}

func ListNs() []string {
	var ns []string
	for n, _ := range data {
		ns = append(ns, n)
	}
	return ns
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

func ListIamUsers(namespace string) ([]model.IamUser, bool) {
	ns, ok := data[namespace]
	if !ok {
		return nil, false
	}
	return ns.IamUsers, true
}

func ListAccessKeys(namespace, username string) ([]model.AccessKey, bool) {
	ns, ok := data[namespace]
	if !ok {
		return nil, false
	}
	for _, u := range ns.IamUsers {
		if u.UserName == username {
			var maskedAccessKey []model.AccessKey
			for _, key := range u.AccessKeys {
				maskedAccessKey = append(maskedAccessKey, model.AccessKey{AccessKeyId: key.AccessKeyId, UserName: username, SecretAccessKey: "<masked>"})
			}
			return maskedAccessKey, true
		}
	}
	return nil, false
}

func CreateIamUser(namespace, username string) (model.IamUser, int) {
	ns, ok := data[namespace]
	if !ok {
		return model.IamUser{}, 404
	}
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
	ns.IamUsers = append(ns.IamUsers, user)
	return user, 200
}

func CreateAccessKey(namespace, username string) (model.AccessKey, int) {
	key := model.AccessKey{}
	ns, ok := data[namespace]
	if !ok {
		return key, 404
	}
	var user model.IamUser
	for i, u := range ns.IamUsers {
		if u.UserName == username {
			user = ns.IamUsers[i] // since u is local copy
			break
		}
	}
	if user.UserName == "" {
		return key, 404
	}
	if len(user.AccessKeys) == 2 {
		return key, 409
	}
	key.UserName = username
	key.AccessKeyId = randString(8)
	key.SecretAccessKey = randString(16)
	user.AccessKeys = append(user.AccessKeys, key)
	return key, 200
}

func DeleteAccessKey(namespace, username, keyId string) int {
	ns, ok := data[namespace]
	if !ok {
		return 404
	}
	var user model.IamUser
	for i, u := range ns.IamUsers {
		if u.UserName == username {
			user = ns.IamUsers[i] // since u is local copy
			break
		}
	}
	if user.UserName == "" {
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

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
