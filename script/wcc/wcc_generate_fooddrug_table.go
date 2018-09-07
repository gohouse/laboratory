package main

import (
	"github.com/gohouse/gorose"
	_ "github.com/go-sql-driver/mysql"
	"github.com/robfig/cron"
	"log"
	"time"
	"github.com/BurntSushi/toml"
	"flag"
	"fmt"
	"strings"
)

var config = map[string]string{ // 定义名为 mysql_dev 的数据库配置
	"host": "192.168.200.248",          // 数据库地址
	"username": "gcore",                // 数据库用户名
	"password": "gcore",                // 数据库密码
	"port": "3306",                     // 端口
	"database": "wcc_service_fooddrug_test", // 链接的数据库名字
	"charset": "utf8",                  // 字符集
	"protocol": "tcp",                  // 链接协议
	"prefix": "",                       // 表前缀
	"driver": "mysql",                  // 数据库驱动(mysql,sqlite,postgres,oracle,mssql)
}

var connection gorose.Connection
var fields = []string{
	"check_org_id",
	"check_user_name",
	"check_org_name",
	"del_user",
	"org_id",
	"org_type",
	"reason",
	"resource_org_name",
	"sample_state",
	"scaname",
	"scbid",
	"scbname",
	"status",
	"ueid",
	"user",
	"create_at",
	"update_at",
	"fa_org_name",
	"fa_org_id",
	"org_name",

	"health_code",
	"health_func_cat",
	"rainbowcode",
	"sp_d_28",
	"sp_d_38",
	"sp_d_46",
	"sp_d_86",
	"sp_i_jgback",
	"sp_i_state",
	"sp_s_1",
	"sp_s_13",
	"sp_s_14",
	"sp_s_16",
	"sp_s_17",
	"sp_s_18",
	"sp_s_19",
	"sp_s_2",
	"sp_s_2_1",
	"sp_s_20",
	"sp_s_202",
	"sp_s_215",
	"sp_s_220",
	"sp_s_27",
	"sp_s_3",
	"sp_s_35",
	"sp_s_37",
	"sp_s_4",
	"sp_s_43",
	"sp_s_44",
	"sp_s_45",
	"sp_s_5",
	"sp_s_64",
	"sp_s_68",
	"submit_d_flag",
	"created_at",
	"updated_at",

	"task_id",
}
var fieldsSelect = []string{
	"b.check_org_id",
	"d.tname as check_user_name",
	"b.check_org_name",
	"a.del_user",
	"b.org_id",
	"b.org_type",
	"a.reason",
	"b.resource_org_name",
	"a.sample_state",
	"b.scaname",
	"b.scbid",
	"b.scbname",
	"a.status",
	"b.ueid",
	"b.user",
	"a.create_at",
	"a.update_at",
	"b.fa_org_name",
	"b.fa_org_id",
	"b.org_name",

	"c.health_code",
	"c.health_func_cat",
	"c.rainbowcode",
	"c.sp_d_28",
	"c.sp_d_38",
	"c.sp_d_46",
	"c.sp_d_86",
	"c.sp_i_jgback",
	"c.sp_i_state",
	"c.sp_s_1",
	"c.sp_s_13",
	"c.sp_s_14",
	"c.sp_s_16",
	"c.sp_s_17",
	"c.sp_s_18",
	"c.sp_s_19",
	"c.sp_s_2",
	"c.sp_s_2_1",
	"c.sp_s_20",
	"c.sp_s_202",
	"c.sp_s_215",
	"c.sp_s_220",
	"c.sp_s_27",
	"c.sp_s_3",
	"c.sp_s_35",
	"c.sp_s_37",
	"c.sp_s_4",
	"c.sp_s_43",
	"c.sp_s_44",
	"c.sp_s_45",
	"c.sp_s_5",
	"c.sp_s_64",
	"c.sp_s_68",
	"c.submit_d_flag",
	"c.created_at",
	"c.updated_at",

	"b.fcc_grade_one_name",
	"b.fcc_grade_two_name",
	"b.fcc_grade_three_name",
	"b.fcc_grade_four_name",
	"a.sample_code",
	"a.id as task_id",
}

func init() {
}

func main() {
	//defer connection.Close()
	var err error
	//fmt.Println(decodeDbToml())
	connection, err = gorose.Open(config)
	if err != nil {
		panic(err)
	}
	log.Println("cron running: ", time.Now().Format("2006-01-02 15:04:05"))
	// 开始迁移
	//Crontab()
	Migration()
}

func parseFlag(key string) string {
	s := flag.String(key,
		"/Users/fizz/go/src/github.com/gohouse/laboratory/script/wcc/db.toml", "配置文件地址")
	//"/Users/fizz/go/src/github.com/gohouse/laboratory/script/wcc/db.toml", "配置文件地址")
	flag.Parse()

	return *s
}

// 定义config的struct
type Configer struct {
	Default         string
	SetMaxOpenConns int
	SetMaxIdleConns int
	Connections     map[string]map[string]string
}

func Crontab() {
	i := 0
	c := cron.New()
	// 每分钟的第 44s 执行
	//spec := "44 */1 * * * ?"
	//每3分钟的第一秒执行一次
	spec2 := "1 */3 * * * ?"
	c.AddFunc(spec2, func() {
		i++
		log.Println("cron running: ", time.Now().Format("2006-01-02 15:04:05"), i)
		Migration()
	})
	c.Start()

	select {}
}
func decodeDbToml() map[string]interface{} {
	var conf2 Configer
	var cpath string = parseFlag("conf")
	//fmt.Println(cpath)

	if _, err := toml.DecodeFile(cpath, &conf2); err != nil {
		log.Fatal(err)
	}

	var confReal = map[string]interface{}{
		"Default":         conf2.Default,
		"SetMaxOpenConns": conf2.SetMaxOpenConns,
		"SetMaxIdleConns": conf2.SetMaxIdleConns,
		"Connections":     conf2.Connections,
	}

	return confReal
}

func Migration() {
	var lastDate string
	// 获取最后一条插入的时间
	lastUpDate, err3 := M("fd_sample_list").Order("update_at desc").Value("update_at")
	//fmt.Println(lastUpDate, err3)
	//return
	if err3 != nil {
		// 记录log
		M("fd_sample_list_log").Data(map[string]interface{}{
			"sample_code": "",
			"mark":        "未获取到插入历史数据",
			"errors":      err3.Error(),
		}).Insert()
	}
	//fmt.Println(lastUpDate)
	if lastUpDate == nil {
		//lastDate = time.Now().Format("2006-01-02 15:04:05")
		lastDate = "2016-02-01"
	} else {
		lastDate = lastUpDate.(string)
	}
	//lastDate = "2016-02-01"
	//fmt.Println(lastDate)
	db := connection.GetInstance()
	db.Table("fd_sample_code a").
		LeftJoin("fd_task_down b on a.tdid=b.tdid").
		LeftJoin("fd_samples c on a.sample_code=c.sp_s_16").
		LeftJoin("fd_user_info d on b.ueid=d.id and b.org_type=d.user_type").
		Fields(strings.Join(fieldsSelect, ",")).
	//Where("a.tdid", "!=", "0").
		Where("a.update_at", ">=", lastDate)
	resTest, err5 := db.First()
	//resTest33, _ := db.Count()
	//fmt.Println(resTest33)
	//fmt.Println(db.LastSql)
	//os.Exit(1)
	if err5 != nil {
		// 记录log
		M("fd_sample_list_log").Data(map[string]interface{}{
			"sample_code": "",
			"mark":        "相关条件未查出数据" + db.LastSql,
			"errors":      err5.Error(),
		}).Insert()
	}

	if resTest == nil {
		fmt.Println("相关条件未查出数据")
		return
	}

	db.Chunk(1000, func(data []map[string]interface{}) {
		if len(data) > 0 {
			var result int
			var err error
			for _, item := range data {
				sp_s_16 := item["sample_code"]
				task_id := item["task_id"]
				// 检查是否已经插入数据库
				count, err2 := M("fd_sample_list").Where("task_id", task_id).Count()
				if err2 != nil {
					// 记录log
					M("fd_sample_list_log").Data(map[string]interface{}{
						"sample_code": sp_s_16,
						"mark":        "统计出错",
						"errors":      err2.Error(),
					}).Insert()
				}
				// item做兼容
				if item["sp_s_17"]==nil || item["sp_s_17"].(string) == "" {
					item["sp_s_17"] = item["fcc_grade_one_name"]
				}
				if item["sp_s_18"]==nil || item["sp_s_18"].(string) == "" {
					item["sp_s_18"] = item["fcc_grade_two_name"]
				}
				if item["sp_s_19"]==nil || item["sp_s_19"].(string) == "" {
					item["sp_s_19"] = item["fcc_grade_three_name"]
				}
				if item["sp_s_20"]==nil || item["sp_s_20"].(string) == "" {
					item["sp_s_20"] = item["fcc_grade_four_name"]
				}
				if item["sp_s_16"]==nil || item["sp_s_16"].(string) == "" {
					item["sp_s_16"] = sp_s_16
				}
				// 抽样人员名字
				if item["sp_s_37"]==nil || item["sp_s_37"].(string) == "" {
					item["sp_s_37"] = item["check_user_name"]
				}
				itemTmp := MapValueNilToNull(item)
				//fmt.Println(item["sample_code"], itemTmp["sp_s_16"])
				//fmt.Println(item)
				//os.Exit(1)
				db2 := connection.GetInstance()
				if count == 0 {
					result, err = db2.Table("fd_sample_list").Data(itemTmp).Insert()
					fmt.Println(db2.LastSql)
				} else {
					delete(item, "task_id")
					result, err = db2.Table("fd_sample_list").
						Where("task_id", task_id).Data(itemTmp).Update()
				}
				if err != nil || result == 0 {
					// 记录log
					M("fd_sample_list_log").Data(map[string]interface{}{
						"sample_code": sp_s_16,
						"mark":        "同步失败:" + db2.LastSql,
						"errors":      fmt.Sprint(err),
					}).Insert()
				} else {
					//fmt.Println("成功:", sp_s_16)
				}
			}
		} else {
			// 记录log
			M("fd_sample_list_log").Data(map[string]interface{}{
				"sample_code": "",
				"mark":        "暂无数据",
				"errors":      "",
			}).Insert()
		}
	})
}

func M(table string) *gorose.Database {
	db := connection.GetInstance()
	return db.Table(table)
}

func MapValueNilToNull(res map[string]interface{}) map[string]interface{} {
	if len(res) > 0 {
		tmpResult := make(map[string]interface{})
		//for k, v := range res {
		//	switch v.(type) {
		//	case nil:
		//		res[k] = "null"
		//	}
		//}
		for _,k := range fields{
			switch res[k].(type) {
			case nil:
				tmpResult[k] = "null"
			default:
				tmpResult[k] = res[k]
			}
		}
		return tmpResult
	}
	return res
}
