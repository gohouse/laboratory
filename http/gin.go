package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gohouse/gorose"
	"strconv"
)

func GetList(where [][]interface{}, ext map[string]int)  {
	gorose.Connection{}.NewSession().Table("xxx").Where(where).
		Limit(ext["limit"]).Offset(ext["offset"]).Get()
}

func GetTaskList(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	getType,_ := strconv.Atoi(c.DefaultQuery("getType", ""))
	lat,_ := strconv.Atoi(c.DefaultQuery("lat", ""))
	lng,_ := strconv.Atoi(c.DefaultQuery("lng", ""))
	var where [][]interface{}
	if getType>0 {
		where = append(where, []interface{}{"status",">",0})
	}
	// 经纬度
	where = append(where, []interface{}{fmt.Sprintf("POINT(%f, %f)<5000", lat,lng)})
	ext := map[string]int{
		"limit": limit,
		"page":  page,
	}
	GetList(where, ext)
}
