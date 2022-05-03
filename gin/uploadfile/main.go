package main

//文件上传

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
)

func main() {

	r := gin.Default()

	//设置文件上传的最大大小
	r.MaxMultipartMemory = 32 << 20 //32M
	r.POST("/uploadone", uploadOne)
	r.POST("/uploadmore", moreFile)

	err := r.Run(":9000")
	if err != nil {
		panic(err)
	}
}

// 单个文件上传

func uploadOne(c *gin.Context) {
	//name属性对应文件上传的key
	file, err := c.FormFile("uploadfile")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	//	上传文件到指定目录
	dst := filepath.Join("./uploadfile", file.Filename)
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "upload success",
	})
}

//上传多个文件

func moreFile(c *gin.Context) {
	//MultipartForm 读取多个文件
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
	}

	files := form.File["uploadfile"]

	for _, file := range files {
		log.Println(file.Filename)
		//	上传到之地文件夹
		dst := filepath.Join("./uploadfile", file.Filename)
		err := c.SaveUploadedFile(file, dst)
		if err != nil {
			panic(err)
		}

	}

	c.JSON(http.StatusOK, gin.H{
		"message": "upload success",
	})
}
