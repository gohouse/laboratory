// Demo code for the Flex primitive.
package chatcli

import (
	"encoding/json"
	"fmt"
	"github.com/gen2brain/beeep"
	"github.com/gen2brain/dlgs"
	"github.com/go-redis/redis"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"runtime"
	"time"
)

var (
	rds *redis.Client
	u   *User
	ch  = "ch"
	//chatKey    = "z:gasii:chat"	// z:chat:gasii
	chatKey = "z:chat:gasii" // z:chat:gasii
	//idKey      = "a:gasii:id"
	idKey      = "a:id:gasii"
	keyVersion = "h:version" // hset version no xxx; hset version url xxx
)

const version = 11

func init() {
	rds = redis.NewClient(&redis.Options{
		Addr:     "redis-17876.c1.ap-southeast-1-1.ec2.cloud.redislabs.com:17876",
		Password: "DOiWgwDu7RaWdb5HIqh4TkmJbjoXQnt886",
	})
}

func Run() {
	// 检查版本
	checkVersion()
	// 验证登录
	checkLogin()
	var g = NewGroup("brag")
	var c = NewChat(g, u)
	// 获取历史消息
	getMsgHistory(c)
	// 不断获取最新消息,放入通道
	go getMsg(c)
	go saveMsg(c)
	ui(c)
}

func checkVersion() {
	i, _ := rds.HGet(keyVersion, "no").Int()
	ddurl := rds.HGet(keyVersion, "url").Val()
	if i > version {
		//解析url
		uri, err := url.ParseRequestURI(ddurl)
		if err != nil {
			panic("网址错误")
		}
		fileName := path.Base(uri.Path)
		log.Println("升级最新版本,下载地址: ", ddurl)
		question, _ := dlgs.Question("发现新版本", "是否更新到最新版本", false)
		if question {
			log.Println("最新版本,请稍候...")
			resp, err := http.Get(ddurl)
			if err != nil {
				log.Println("更新版本失败,请尝试使用浏览器下载,下载地址为: ", ddurl)
				return
			}
			defer resp.Body.Close()
			all, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println("更新版本失败,请尝试使用浏览器下载,下载地址为: ", ddurl)
				return
			}
			ioutil.WriteFile(fileName, all, 0777)
			log.Println("下载完毕,请打开新版本: ", fileName)
			os.Exit(1)
		}
	}
}

func getMsgHistory(c *Chat) {
	val := rds.ZRevRange(chatKey, 0, 100).Val()
	for _, v := range val {
		var m Message
		json.Unmarshal([]byte(v), &m)
		c.Messages = append(c.Messages, m)
	}
	if len(c.Messages) > 0 {
		var tmpMsg []Message
		for i := len(c.Messages) - 1; i >= 0; i-- {
			tmpMsg = append(tmpMsg, c.Messages[i])
		}
		c.Messages = tmpMsg[:]
	}
}
func getMsg(c *Chat) {
	subscribe := rds.Subscribe(ch)
	for {
		message, err := subscribe.ReceiveMessage()
		if err != nil {
			panic(err.Error())
		}

		var m Message
		json.Unmarshal([]byte(message.Payload), &m)
		c.RcvMsgs <- &m
	}
}
func saveMsg(c *Chat) {
	for {
		select {
		case msg := <-c.SendMsgs:
			marshal, _ := json.Marshal(msg)

			rds.ZAdd(chatKey, redis.Z{
				Score:  float64(rds.Incr(idKey).Val()),
				Member: string(marshal),
			})

			rds.Publish(ch, string(marshal))
			if len(c.Messages) > 100 {
				c.Messages = c.Messages[len(c.Messages)-70:]
			}
			c.Messages = append(c.Messages, *msg)
		}
	}
}

func checkLogin() {
	f := "chat.txt"
	fp, err := os.OpenFile(f, os.O_RDWR|os.O_CREATE, 0777)
	all, err := ioutil.ReadAll(fp)
	if err != nil {
		panic(err.Error())
	}
	u = &User{}
	if string(all) == "" {
		fmt.Println("请给自己去一个帅气的名字: ")
		fmt.Scanln(&u.Username)
		fp.Write([]byte(u.Username))
		fp.Close()
	} else {
		u = NewUser(string(all))
	}
}

type Message struct {
	Msg  string
	Time time.Time
	User *User
	Type string
}

func NewMessage(msg string, user *User) *Message {
	return &Message{Msg: msg, User: user, Time: time.Now()}
}

type User struct {
	Username string
}

func NewUser(username string) *User {
	return &User{Username: username}
}

type IGroup interface {
	Userlist() []User
	History() []Message
	ShowHistory(w io.Writer)
	UserJoin(user *User)
}
type Group struct {
	Groupname string
	Users     []User
	Messages  []Message
}

func NewGroup(groupname string) *Group {
	return &Group{Groupname: groupname}
}
func (g *Group) Userlist() []User {
	return g.Users
}
func (g *Group) UserJoin(user *User) {
	g.Users = append(g.Users, *user)
}
func (g *Group) UserLeave(user *User) {
	var tmp []User
	for k, v := range g.Users {
		if v.Username == user.Username {
			tmp = append(tmp, g.Users[:k]...)
			tmp = append(tmp, g.Users[k+1:]...)
			g.Users = tmp[:]
			break
		}
	}
}
func (g *Group) ShowHistory(w io.Writer) {
	for _,v := range g.Messages{
		showMsg(w, &v)
	}
}
func (g *Group) History() []Message {
	return g.Messages
}

type IMessage interface {
	SendMsgString(msg string)
	SendMsg(msg *Message)
	ReceiveMsg(w io.Writer)
}
type IChat interface {
	IGroup
	IMessage
	GetUser() *User
	GetGroup() *Group
}
type Chat struct {
	*Group
	*User
	RcvMsgs  chan *Message
	SendMsgs chan *Message
	Msgs     chan *Message
}

func NewChat(group *Group, user *User) *Chat {
	return &Chat{Group: group, User: user, RcvMsgs: make(chan *Message, 10), SendMsgs: make(chan *Message, 10)}
}

func (c *Chat) GetUser() *User {
	return c.User
}

func (c *Chat) GetGroup() *Group {
	return c.Group
}

func (c *Chat) SendMsgString(msg string) {
	if msg == "" {
		return
	}
	c.SendMsg(NewMessage(msg, c.GetUser()))
}

func (c *Chat) SendMsg(msg *Message) {
	if msg.Msg == "" {
		return
	}
	c.SendMsgs <- msg
}

func (c *Chat) ReceiveMsg(w io.Writer) {
	for {
		select {
		case msg := <-c.RcvMsgs:
			//fmt.Fprintf(w, "%s[%s]: %s\n", msg.User.Username, msg.Time.Format("01-02 15:04"), msg.Msg)
			showMsg(w, msg)
			if msg.User.Username != u.Username {
				if runtime.GOOS == "darwin" {
					beeep.Notify(msg.User.Username, "have a rest for 5 minutes.", "")
				} else {
					beeep.Notify(msg.User.Username, msg.Msg, "")
				}
			}
		}
	}
}

func showMsg(w io.Writer, msg *Message)  {
	fmt.Fprintf(w, "%s[%s]: %s\n", msg.User.Username, msg.Time.Format("01-02 15:04"), msg.Msg)
}
