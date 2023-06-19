package service

import (
	"bufio"
	"ecs-mock/model"
	"github.com/gookit/slog"
	"os"
)

var token string

func GetToken() string {
	return token
}

func WriteAdminPwd(username string, user model.VdcUser) {
	if err := os.WriteFile(username+".txt", []byte(user.Password), 0644); err != nil {
		slog.Error(err)
		slog.Exit(1)
	}
}

func Login(username, pwd string) (string, bool) {
	stored := readAdminPwd(username)
	if pwd == stored {
		token = RandString(10)
		return token, true
	}
	return "", false
}

func readAdminPwd(username string) string {
	f, err := os.Open(username + ".txt")
	defer f.Close()
	if err != nil {
		slog.Error(err)
		slog.Exit(1)
	}
	scan := bufio.NewScanner(f)
	scan.Scan()
	pwd := scan.Text()
	return pwd
}
