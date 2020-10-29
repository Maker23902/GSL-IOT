package v1

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

// @Tags ExaFileUploadAndDownload
// @Summary 上传
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body model.ExaFileUploadAndDownload true "传入文件里面id即可"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /fileUploadAndDownload/UploadCode [post]
func UploadCode(c *gin.Context) {

	data, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Printf("ctx.Request.body: %v", string(data))
	//service.FilerCreate()
}
