package chatp2p

import (
	"errors"
	"io/ioutil"
)

type User struct {
	Username string
	Password string
	Groups   []string
}

var userinfoFile = "chat.txt"
var userKey string

func GetUserKey() (key string, err error) {
	if userKey != "" {
		key = userKey
		return
	}
	var b []byte
	b, err = ioutil.ReadFile(userinfoFile)
	if err != nil {
		return
	}
	if len(b) == 0 {
		err = errors.New("please login first")
	}
	userKey = string(b)
	key = userKey
	return
}
