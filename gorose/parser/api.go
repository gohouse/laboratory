package parser

import (
	"errors"
	"github.com/gohouse/laboratory/gorose/config"
)

// 检查解析器是否实现了接口
var jsonParser IParser = &JsonConfigParser{}
var tomlParser IParser = &TomlConfigParser{}

// 注册解析器
var fileParsers = map[string]IParser{
	config.JSON: jsonParser,
	config.TOML: tomlParser,
}

// 对外提供接口
func NewFileParser(file, fileType string) (*FileParser, error) {
	if pr, ok := fileParsers[fileType]; ok {
		return pr.Parse(file)
	}
	return &FileParser{}, errors.New("不支持的配置类型")
}
