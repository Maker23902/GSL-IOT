package service

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// FilerCreate 创建文件
func FilerCreate(c *gin.Context) {
	iscreate := true //创建文件是否成功
	//创建文件
	f, err := os.Create("Logic/lua/log.text")
	if err != nil {
		iscreate = false
	}
	defer f.Close()
	fmt.Println(f)
	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"path":    "static/txtfile/log.text",
		"success": iscreate,
	})
}

type fileForm struct {
	APPID   string `form:"appid" binding:"required"`
	Cyclems string `form:"cyclems" binding:"required"`
	Content string `form:"content" binding:"required"`
}

// FilerWrite 将内容写入文件
func FilerWrite(c *gin.Context) bool {
	var f *os.File
	var err1 error
	var form fileForm

	iswrite := true //写入文件是否成功

	if c.ShouldBind(&form) == nil {
		fmt.Println(form.APPID)
		fmt.Println(form.Cyclems)
		fmt.Println(form.Content)
	}
	filename := form.APPID + ".lua"
	if checkFileIsExist(filename) { //如果文件存在
		f, err1 = os.OpenFile(filename, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766) //打开文件 os.O_APPEND
		fmt.Println("文件存在")
	} else {
		f, err1 = os.Create(filename) //创建文件
		fmt.Println("文件不存在")
	}
	defer f.Close()
	check(err1)
	//data, _ := ioutil.ReadAll(c.Request.Body)
	n, err1 := io.WriteString(f, form.Content) //写入文件(字符串)
	check(err1)

	//需要写入到文件的内容
	// d1 := []byte(content)
	// err := ioutil.WriteFile(path, d1, 0644)

	// if err != nil {
	// 	iswrite = false
	// }
	iswrite = true
	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"success": iswrite,
		"info":    n,
	})

	return iswrite
}

// FilerRead 读取文件内容.
func FilerRead(c *gin.Context) {
	isread := true //读取文件是否成功
	path := c.PostForm("path")
	//文件读取任务是将文件内容读取到内存中。
	info, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		isread = false
	}
	fmt.Println(info)
	result := string(info)

	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"content": result,
		"success": isread,
	})
}

// FilerDelete 删除文件
func FilerDelete(c *gin.Context) {

	isremove := false          //删除文件是否成功
	path := c.PostForm("path") //源文件路径

	//删除文件
	cuowu := os.Remove(path)

	if cuowu != nil {
		//如果删除失败则输出 file remove Error!
		fmt.Println("file remove Error!")
		//输出错误详细信息
		fmt.Printf("%s", cuowu)
	} else {
		//如果删除成功则输出 file remove OK!
		fmt.Print("file remove OK!")
		isremove = true
	}
	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"success": isremove,
	})
}
