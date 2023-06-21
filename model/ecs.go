package model

type Namespaces struct {
	Namespace []Namespace `json:"namespace"`
}

type Namespace struct {
	Name        string       `json:"name"`
	NativeUsers []NativeUser `json:"-"`
	IamUsers    []*IamUser   `json:"-"`
}

type NativeUsers struct {
	Users []NativeUser `json:"blobuser"`
}

type NativeUser struct {
	Userid string `json:"userid,omitempty"`
	Name   string `json:"name,omitempty"`
}

type ListIamUsers struct {
	ListUsersResult Users `json:"ListUsersResult"`
}

type Users struct {
	Users []IamUser `json:"Users"`
}

type IamUser struct {
	UserName   string      `json:"UserName"`
	AccessKeys []AccessKey `json:"-"`
}

type AccessKey struct {
	AccessKeyId     string `json:"AccessKeyId"`
	UserName        string `json:"UserName"`
	SecretAccessKey string `json:"SecretAccessKey,omitempty"`
	CreateDate      string `json:"CreateDate"`
}

type CreateAccessKey struct {
	CreateAccessKeyResult CreateAccessKeyResult `json:"CreateAccessKeyResult"`
}

type CreateAccessKeyResult struct {
	AccessKey AccessKey `json:"AccessKey"`
}

type VdcUser struct {
	Password        string `json:"password"`
	IsSystemAdmin   string `json:"isSystemAdmin"`
	IsSystemMonitor string `json:"isSystemMonitor"`
	IsSecurityAdmin string `json:"isSecurityAdmin"`
}

type ListAccessKeys struct {
	ListAccessKeysResult AccessKeyMetadata `json:"ListAccessKeysResult"`
}

type AccessKeyMetadata struct {
	AccessKeys []AccessKey `json:"AccessKeyMetadata"`
}
