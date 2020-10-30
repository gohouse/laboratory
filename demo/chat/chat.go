package chat

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gohouse/demo/chat/stock"
	"github.com/gohouse/t"
	"github.com/gookit/color"
	"io/ioutil"
	"os"
	"strings"
	"time"
)
var (
	rds      *redis.Client
	chatKey  = "z:chat"
	idKey = "a:id"
	records []string
	lastMsg string
	u string
	l int64
	d time.Duration
)
func init() {
	rds = redis.NewClient(&redis.Options{
		Addr:     "redis-17876.c1.ap-southeast-1-1.ec2.cloud.redislabs.com:17876",
		Password: "DOiWgwDu7RaWdb5HIqh4TkmJbjoXQnt886",
	})
	flag.Int64Var(&l,"l",50,"设定读取消息的条数")
	flag.DurationVar(&d,"d",3,"设定读取消息的时间周期,默认3s")
	flag.Parse()
}
func Run() {
	checkUsername()
	fmt.Println("模式选择: \n1. 聊天模式 \n2. 查看消息模式(3s自动刷新一次)\n请输入选择: ")
	var choice string
	fmt.Scanln(&choice)
	if choice == "1" {// 开始聊天
		var msg string
		for {
			getMsg()
			msg = readStr()
			if strings.HasPrefix(msg, "/gp") {
				space := t.SpliteAndTrimSpace(msg, " ")
				if len(space) != 2 || space[1] == "" {
					msg = fmt.Sprintf("%s\n%s",msg,"命令格式错误")
				} else {
					msg = fmt.Sprintf("%s\n%s (%s)",msg, stock.ReadStock(space[1]),time.Now().Format("2006-01-02 15:04:05"))
				}
				saveMsg(msg)
				continue
			}
			switch msg {
			case "":
				continue
			case `\q`:
				os.Exit(1)
			default:
				saveMsg(msg)
			}
		}
	} else {
		getMsg()
		// 周期获取记录,默认3s
		ticker := time.NewTicker(d * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				getMsg()
			}
		}
	}
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
func readStr() string {
	fmt.Print("_> ")
	cmdReader := bufio.NewReader(os.Stdin)
	if cmdStr, err := cmdReader.ReadString('\n'); err == nil {
		//这里把读取的数据后面的换行去掉，对于Mac是"\r"，Linux下面 //是"\n"，Windows下面是"\r\n"，所以为了支持多平台，直接用"\r\n"作为过滤字符
		cmdStr = strings.Trim(cmdStr, "\r\n")
		return cmdStr
	}
	return ""
}
func saveMsg(msg string) {
	_,err := rds.ZAdd(chatKey, redis.Z{
		Score:  float64(rds.Incr(idKey).Val()),
		Member: fmt.Sprintf("%s: %s", u, msg),
	}).Result()
	if err!=nil {
		panic(err.Error())
	}
}
func getMsg()  {
	records = rds.ZRevRange(chatKey, 0, l).Val()
	if len(records) == 0 {
		return
	}
	if records[0] != lastMsg {
		lastMsg = records[0]
		fmt.Println("----------------------------------------------")
		for i:=len(records)-1;i>=0;i-- {
			fields := strings.Fields(records[i])
			color.Red.Print(fields[0])
			fmt.Println(fields[1:])
			fmt.Println()
		}
	}
}

