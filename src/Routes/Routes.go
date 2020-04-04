package Routes

import (
	"../Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("urls", Controllers.GetUrls)
		v1.POST("urls", Controllers.CreateUrl)
		v1.GET("urls/:id", Controllers.GetUrl)
		v1.DELETE("urls/:id", Controllers.DeleteUrl)
	}
	return r
}