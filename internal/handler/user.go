package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/gin-scaffold/internal/cookie"
	"github.com/zenpk/gin-scaffold/pkg/ep"
)

type User struct{}

func (u User) Register(c *gin.Context) {
	// TODO
	response(c, packer.Pack(ep.ErrOK))
}

func (u User) Logout(c *gin.Context) {
	cookie.ClearAllUserInfos(c)
	errPack := ep.ErrOK
	errPack.Msg = "successfully logged out"
	response(c, CommonResp{
		Err: errPack,
	})
}
