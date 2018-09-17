package parser

import (
	"encoding/json"
	"io/ioutil"
)

type JsonConfigParser struct {
}

func (c JsonConfigParser) Parse(file string) (p *FileParser, err error) {
	var fp []byte
	fp, err = ioutil.ReadFile(file)
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(fp), &p)
	return
}
