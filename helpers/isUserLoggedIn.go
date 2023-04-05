package helpers

import "github.com/gin-gonic/gin"

func IsUserLoggedIn(c *gin.Context) bool {
	return (c.GetUint64("user_id") > 0)
}
