package config

const (
	MYSQL  = "mysql"
	SQLITE = "sqlite"
)

const (
	JSON = "json"
	TOML = "toml"
)

var ConstsType = map[string]string{
	MYSQL:  "driver", // 驱动
	SQLITE: "driver", // 驱动
	JSON:   "file",   // 文件
	TOML:   "file",   // 文件
}
