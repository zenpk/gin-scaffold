package middleware

import (
	"github.com/gin-gonic/gin"
)

// RequireLogin if no token then abort
func RequireLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		//token := cookie.GetAccessToken(c)
		//packer := ep.Packer{V: handler.CommonResp{}}
		//if token == "" { // no access_token
		//	refreshToken := cookie.GetRefreshToken(c)
		//	if refreshToken == "" { // no refresh_token
		//		c.JSON(http.StatusOK, packer.Pack(ep.ErrNotLogin))
		//		c.Abort()
		//		return
		//	}
		//	req := &pb.GenAccessTokenRequest{RefreshToken: refreshToken}
		//	tokenResp, err := rpc.Client.Token.GenAccessToken(req)
		//	if err != nil {
		//		c.JSON(http.StatusOK, packer.PackWithError(err))
		//		c.Abort()
		//		return
		//	}
		//	accessToken := tokenResp.AccessToken
		//	cookie.SetAccessToken(c, accessToken)
		//	if err := cookie.SetAllFromAccessToken(c, accessToken); err != nil {
		//		c.JSON(http.StatusOK, packer.PackWithError(err))
		//		c.Abort()
		//		return
		//	}
		//} else if err := cookie.SetAllFromAccessToken(c, token); err != nil {
		//	c.JSON(http.StatusOK, packer.PackWithError(err))
		//	c.Abort()
		//	return
		//}
		c.Next()
	}
}
