package router

import (
	v1 "gin-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitScriptManager(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("scriptmanager")
	// .Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{

		ApiRouter.POST("/uploadcode", v1.UploadCode) //上传脚本文件
	}
}
