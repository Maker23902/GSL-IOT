package v1

import (
	"fmt"
	"gin-vue-admin/IOTService"
	"gin-vue-admin/global/response"
	"gin-vue-admin/service"

	"github.com/gin-gonic/gin"
)

// @Tags ExaFileUploadAndDownload
// @Summary 上传
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body model.ExaFileUploadAndDownload true "传入文件里面id即可"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /scriptmanager/UploadCode [post]
func UploadCode(c *gin.Context) {

	err := service.FilerWrite(c)
	if err != true {
		panic(err)
	}

	//data, _ := ioutil.ReadAll(c.Request.Body)
	//fmt.Printf("ctx.Request.body: %v", string(data))
	//service.FilerCreate()
}

// @Tags ExaFileUploadAndDownload
// @Summary 上传
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body model.ExaFileUploadAndDownload true "传入文件里面id即可"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /scriptmanager/reqGetcode [get]
func ReqGetCode(c *gin.Context) {
	service.FilerRead(c)
}

// @Tags ExaFileUploadAndDownload
// @Summary 上传
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body model.ExaFileUploadAndDownload true "传入文件里面id即可"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /scriptmanager/reqStartTask [post]
func ReqStartTask(c *gin.Context) {
	err := IOTService.Starttask(c)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("启动任务失败，%v", err), c)
		return
	} else {
		response.Ok(c)
	}
}

// @Tags ExaFileUploadAndDownload
// @Summary 上传
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body model.ExaFileUploadAndDownload true "传入文件里面id即可"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /scriptmanager/reqPuseTask [post]
func ReqPuseTask(c *gin.Context) {

}
