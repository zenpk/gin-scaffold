package controller

import "github.com/zenpk/gin-scaffold/internal/handler"

// AllHandler Gin HTTP request handler
type AllHandler struct {
	User handler.User
	Temp handler.Temp
}

var ginHandler AllHandler
