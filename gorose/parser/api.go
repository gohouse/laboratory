package parser

import (
	"errors"
	"github.com/gohouse/laboratory/gorose/config"
)

var fileParsers = map[string]IParser{
	config.JSON: JsonConfigParser{},
	config.TOML: TomlConfigParser{},
}

func NewFileParser(file, fileType string) (*FileParser, error) {
	if pr, ok := fileParsers[fileType]; ok {
		return pr.Parse(file)
	}
	return &FileParser{}, errors.New("不支持的配置类型")
}
