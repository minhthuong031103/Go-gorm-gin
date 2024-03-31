package taghandler

import (
	appctx "go-gorm-gin/components/appCtx"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	tags := router.Group("/tags")
	{
		tags.GET("", FindAllTag(appCtx))
		tags.POST("", CreateTag(appCtx))
	}
}
