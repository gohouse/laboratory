package config

import (
	"io/ioutil"
	"encoding/json"
)

type JsonConfigParser struct {
}

func (c JsonConfigParser) Parse(file string) (p *ConfigParser, err error) {
	var fp []byte
	fp, err = ioutil.ReadFile(file)
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(fp), &p)
	return
}
