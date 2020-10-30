package chatp2p

import (
	"fmt"
	"io/ioutil"
	"os"
)

type iChat interface {
	CheckLogin()
	Login()
	Register()
	PasswordReset()
	GroupList()
	JoinGroup()
	LeaveGroup()
	CreateGroup()
	DeleteGroup()
	ChatInGroup()
	Bot()
}

type Chat struct {
	conf *Config
}

func NewChat(conf *Config) *Chat {
	return &Chat{conf: conf}
}

func (c *Chat) Login()  {

}

func (c *Chat) CheckLogin()  {

}

func (c *Chat) GroupList()  {

}

func (c *Chat) GroupJoin(groupName string)  {

}

func checkUsername()  {
	f := "chat.txt"
	fp, err := os.OpenFile(f, os.O_RDWR|os.O_CREATE, 0777)
	all, err := ioutil.ReadAll(fp)
	if err!=nil {
		panic(err.Error())
	}
	if string(all) == "" {
		fmt.Println("请给自己去一个帅气的名字: ")
		fmt.Scanln(&u)
		fp.Write([]byte(u))
		fp.Close()
	} else {
		u = string(all)
	}
}