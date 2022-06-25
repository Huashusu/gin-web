package initialiaze

import (
	"gin-web/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.New()
	Router.Use(middleware.ZapRecovery(true))
	Router.Use(middleware.ZapLogger())

	PublicGroup := Router.Group("")
	{
		PublicGroup.GET("/health", func(c *gin.Context) { c.String(http.StatusOK, "ok") })
	}
	//{
		//otherRouter.InitOtherRouter(PublicGroup)
	//}
	//PrivateGroup := Router.Group("")
	return Router
}
