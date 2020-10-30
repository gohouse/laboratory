package chatp2p

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"sync"
)

var (
	rds *Redis
	once sync.Once
)
var (
	keyGroup     = "h:group"   // group
	keyUser      = "h:user"    // user
	keyGroupUser = "s:%s"      // set keyGroupUser user1 user2
	keyGroupMsg  = "z:%s"      // zadd keyGroupMsg msg1 msg2
	keyVersion   = "h:version" // hset version 1 | hset url xxx
	keySecret    = "s:key:%s"  // setex a:key:xxxx 1 [ttl]
	keyMsgId = "h:msgid"	// 各个群的消息自增ID hinc keyMsgId [group name]
)

type Redis struct {
	*redis.Client
}

func DB() *Redis {
	InitRedis(nil)
	// 验证是不是本人啊
	// 验证登录
	key, err := GetUserKey()
	if err!=nil {
		panic(err.Error())
	}
	if rds.Exists(fmt.Sprintf(keySecret, key)).Val() == 0 {
		panic("请登录")
	}
	return rds
}

func InitRedis(opt *redis.Options) {
	once.Do(func() {
		if opt==nil || opt.Addr == "" || opt.Password == "" {
			opt = &redis.Options{
				Addr:     "redis-17876.c1.ap-southeast-1-1.ec2.cloud.redislabs.com:17876",
				Password: "DOiWgwDu7RaWdb5HIqh4TkmJbjoXQnt886",
			}
		}
		rds = &Redis{Client: redis.NewClient(opt)}
	})
}

func (r *Redis) Register(u *User) error {
	// 放入数据库
	b, err := r.HSetNX(keyUser, Md5(u.Username), Serialize(*u)).Result()
	if err!=nil {
		return err
	}
	if !b {
		return errors.New("已存在,请登录")
	}
	return nil
}

func (r *Redis) Login(u *User) error {
	pass := Md5(u.Password)
	val, err := r.HGet(keyUser, Md5(u.Username)).Bytes()
	if err != nil {
		return err
	}
	UnSerialize(val, u)
	if u.Password != pass {
		return errors.New("wrong with email or password")
	}
	return nil
}

func (r *Redis) GroupCreate(g *Group) error {
	result, err := r.HSetNX(keyGroup, Md5(g.Name), Serialize(*g)).Result()
	if err!=nil {
		return err
	}
	if !result {
		return errors.New("create group fail")
	}
	return nil
}
func (r *Redis) GroupExists(g *Group) bool {
	return r.HExists(keyGroup, g.Name).Val()
}

func (r *Redis) GroupJoin(g *Group, u *User) error {
	if !r.GroupExists(g) {
		return errors.New("group not exists")
	}
	key := fmt.Sprintf(keyGroupUser, Md5(g.Name))
	if !r.SIsMember(key, Md5(u.Username)).Val() {
		aff := r.SAdd(key, Md5(u.Username)).Val()
		if aff == 0 {
			return errors.New("join group fail")
		}
	}
	return nil
}

func (r *Redis) GroupLeave(g *Group, u *User)  {
	key := fmt.Sprintf(keyGroupUser, Md5(g.Name))
	r.SRem(key, Md5(u.Username)).Val()
}

func (r *Redis) MsgSave(m *Message, g *Group)  {
	key := fmt.Sprintf(keyGroupMsg, Md5(g.Name))
	r.ZAdd(key, redis.Z{
		Score:  r.HIncrByFloat(keyMsgId, Md5(g.Name), 1).Val(),
		Member: Serialize(*m),
	})
}

func (r *Redis) MsgList(g *Group, limit int64) (msg []Message) {
	key := fmt.Sprintf(keyGroupUser, Md5(g.Name))
	s := r.ZRevRange(key, 0, limit).String()
	UnSerialize([]byte(s), &msg)
	return
}
