package IOTService

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

type fileForm struct {
	APPID   string `form:"appid" binding:"required"`
	Cyclems string `form:"cyclems" binding:"required"`
}

func Starttask(c *gin.Context) error {

	var form fileForm

	if c.ShouldBind(&form) == nil {
		fmt.Println(form.APPID)
		fmt.Println(form.Cyclems)
	}
	fmt.Print("hello-cron---")

	c := cron.New(cron.WithSeconds())

	return nil
}
