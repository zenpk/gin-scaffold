package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"strings"
	"time"
)

// CORSFilter cross-origin resource sharing
func CORSFilter() gin.HandlerFunc {
	return cors.New(cors.Config{
		//AllowOrigins:     []string{"https://localhost:5173", "http://localhost:5173"}, // ignored because of AllowOriginFunc
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Origin", "x-requested-with"},
		ExposeHeaders:    []string{},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			// requests from the same server
			if strings.Contains(origin, viper.GetString("server.domain")) {
				return true
			}
			// requests from whitelist
			whitelist := viper.GetStringSlice("cors.whitelist")
			for _, allow := range whitelist {
				if strings.Contains(origin, allow) {
					return true
				}
			}
			return false
		},
		MaxAge: 12 * time.Hour,
	})
}
