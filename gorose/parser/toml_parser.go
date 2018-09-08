package parser

import (
	"io/ioutil"
	"github.com/BurntSushi/toml"
)

type TomlConfigParser struct {
}

func (c TomlConfigParser) Parse(file string) (p *FileParser, err error) {
	var fp []byte
	fp, err = ioutil.ReadFile(file)
	if err != nil {
		return
	}

	err = toml.Unmarshal([]byte(fp), &p)
	return
}
