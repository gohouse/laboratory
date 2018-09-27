package main

import (
	"fmt"
)

type Gatway interface {
	CreateInputParser() Parser
	CreeteOutPutEncoder() Encoder
}

type Parser interface {
	ParseData()
}

type Encoder interface {
	EncodeData()
}

type XmlParser struct {
}

func (x *XmlParser) ParseData() {
	fmt.Println("parse xml data ok")
}

type JsonParser struct {
}

func (j *JsonParser) ParseData() {
	fmt.Println("parse Json data ok")
}

type XmlEncoder struct {
}

func (x *XmlEncoder) EncodeData() {
	fmt.Println("Encode xml data ok")
}

type JsonEncoder struct {
}

func (j *JsonEncoder) EncodeData() {
	fmt.Println("Encode Json data ok")
}

type GatwayHuawei struct {
}

func (g *GatwayHuawei) CreateInputParser() Parser {
	return &XmlParser{}
}

func (g *GatwayHuawei) CreeteOutPutEncoder() Encoder {
	return &XmlEncoder{}
}

type GatwayTencent struct {
}

func (g *GatwayTencent) CreateInputParser() Parser {
	return &JsonParser{}
}

func (g *GatwayTencent) CreeteOutPutEncoder() Encoder {
	return &JsonEncoder{}
}

func main() {
	var gatway Gatway

	gatway = &GatwayHuawei{}
	gatway.CreateInputParser().ParseData()
	gatway.CreeteOutPutEncoder().EncodeData()

	gatway = &GatwayTencent{}
	gatway.CreateInputParser().ParseData()
	gatway.CreeteOutPutEncoder().EncodeData()
}
