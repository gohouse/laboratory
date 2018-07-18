package bootService

import (
	"github.com/gohouse/gorose"
	"github.com/gin-gonic/gin"
)
type BootEngin struct {
	// gorose orm
	Connection gorose.Connection

	// gin framwork
	Router  *gin.Engine
}

var Booter = &BootEngin{}

// Boot : 驱动入口
func Boot(options ...func(*BootEngin)) *BootEngin {
	//srv := &Booter{}
	for _, option := range options {
		option(Booter)
	}

	return Booter
}

