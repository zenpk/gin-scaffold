package handler

import "github.com/gin-gonic/gin"

type Temp struct {
}

func (t Temp) TempHandler(c *gin.Context) {
	response(c, "hello world")
}
