package model

type Namespaces struct {
	Namespace []Namespace `json:"namespace"`
}

type Namespace struct {
	Name string `json:"name"`
}

type NativeUsers struct {
	Users []NativeUser `json:"blobuser"`
}

type NativeUser struct {
	Userid string `json:"userid,omitempty"`
	Name   string `json:"name,omitempty"`
}

type IamUsers struct {
	ListUsersResult Users `json:"ListUsersResult"`
}

type Users struct {
	Users []IamUser `json:"Users"`
}

type IamUser struct {
	UserName string `json:"UserName"`
}

type AccessKeys struct {
	ListAccessKeysResult ListAccessKeysResult `json:"ListAccessKeysResult"`
}

type ListAccessKeysResult struct {
	AccessKeyMetadata []AccessKey `json:"AccessKeyMetadata"`
}

type AccessKey struct {
	AccessKeyId string `json:"AccessKeyId"`
	UserName    string `json:"UserName"`
}
