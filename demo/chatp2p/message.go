package chatp2p

import (
	"encoding/json"
	"github.com/go-redis/redis"
)

type MsgType int

const (
	MTmsg MsgType = iota
	MTbot
)

var chatChannel string

type Message struct {
	Username string
	Content  string
	Type     MsgType
}

func (m *Message) SaveMsg() {
	// 获取自增id作为score
	id := Redis().Incr(key).Val()
	// 入库zset表
	str := buildMsg(user, msg)
	if str == "" {
		return
	}
	Redis().ZAdd(key, redis.Z{
		Score:  float64(id),
		Member: str,
	}).Val()
	// 发布频道
	Redis().Publish("", str)
}

func LoadMsg(key string, limit int64) (res []Message) {
	strs := Redis().ZRevRange(key, 0, limit).Val()
	for _, item := range strs {
		var msg Message
		if err := json.Unmarshal([]byte(item), &msg); err == nil {
			res = append(res, msg)
		}
	}
	return
}

func buildMsg(user, msg string) string {
	marshal, err := json.Marshal(Message{
		Username: user,
		Content:  msg,
	})
	if err != nil {
		return ""
	}
	return string(marshal)
}
