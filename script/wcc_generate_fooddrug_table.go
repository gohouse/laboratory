package main

import (
	"github.com/gohouse/gorose"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

var config = map[string]string{ // 定义名为 mysql_dev 的数据库配置
	"host": "192.168.200.248",          // 数据库地址
	"username": "gcore",                // 数据库用户名
	"password": "gcore",                // 数据库密码
	"port": "3306",                     // 端口
	"database": "wcc_service_fooddrug", // 链接的数据库名字
	"charset": "utf8",                  // 字符集
	"protocol": "tcp",                  // 链接协议
	"prefix": "",                       // 表前缀
	"driver": "mysql",                  // 数据库驱动(mysql,sqlite,postgres,oracle,mssql)
}

var connection gorose.Connection
var fields = []string{
	"b.check_org_id",
	//"d.check_user_name",
	"b.check_org_name",
	"a.del_user",
	"c.health_code",
	"c.health_func_cat",
	"b.org_id",
	"b.org_type",
	"c.rainbowcode",
	"a.reason",
	"b.resource_org_name",
	"a.sample_state",
	"b.scaname",
	"b.scbid",
	"b.scbname",
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
	"c.sp_s_71",
	"a.status",
	"c.submit_d_flag",
	"b.ueid",
	"b.user",
	"a.create_at",
	"a.update_at",
	"b.fa_org_name",
	"b.fa_org_id",
	"b.org_name",
}

func init() {
	var err error
	connection, err = gorose.Open(config)
	if err != nil {
		panic(err)
	}
}

func main() {
	defer connection.Close()
	var lastDate string
	// 获取最后一条插入的时间
	lastUpDate, _ := M("fd_sample_list").Order("update_at desc").Value("update_at")
	if lastUpDate == nil {
		//lastDate = time.Now().Format("2006-01-02 15:04:05")
		lastDate = "2017-02-01"
	}
	db := connection.GetInstance()
	db.Table("fd_sample_code a").
		Fields(strings.Join(fields, ",")).
		Join("fd_task_down b on a.tdid=b.tdid").
		Join("fd_samples c on a.sample_code=c.sp_s_16").
		Where("a.tdid", "!=", "0").
		Where("a.update_at", ">", lastDate)

	db.Chunk(100, func(data []map[string]interface{}) {
		if len(data) > 0 {
			for _, item := range data {
				sp_s_16 := item["sp_s_16"]
				// 检查是否已经插入数据库
				count, _ := M("fd_sample_list").Where("sp_s_16", sp_s_16).Count()
				db2 := connection.GetInstance()
				item = MapValueNilToNull(item)
				if count == 0 {
					db2.Table("fd_sample_list").Data(item).Insert()
				} else {
					delete(item, "sp_s_16")
					db2.Table("fd_sample_list").Where("sp_s_16", sp_s_16).Data(item).Update()
				}
			}
		}
	})
}

func M(table string) *gorose.Database {
	db := connection.GetInstance()
	return db.Table(table)
}

func MapValueNilToNull(res map[string]interface{}) map[string]interface{} {
	if len(res) > 0 {
		for k, v := range res {
			switch v.(type) {
			case nil:
				res[k] = "null"}
		}
	}
	return res
}
