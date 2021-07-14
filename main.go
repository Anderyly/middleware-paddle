package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"middlewarePaddle/other"
	"net/http"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var Path string

func main() {
	start()
	r := gin.Default()
	r.POST("/api/picture", picture)
	//r.POST("/api/picture/base64", base64ToPicture)
	r.Run(":8080")

}

func start() {
	// 获取根目录
	p, _ := other.Command("pwd")
	root := strings.Replace(p, "\n", "", -1)

	sysType := runtime.GOOS

	if sysType == "windows" {
		Path = "./image/"
	} else {
		Path = root + "/image/"
	}

	row := other.IsFileExist(Path)
	if row == false {
		os.Mkdir(Path, 0666)
	}
}

func picture(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "message": "Please Upload"})
		return
	}
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(999999-100000)+100000) + path.Ext(file.Filename)
	uperr := c.SaveUploadedFile(file, Path+fileName)
	if uperr != nil {
		c.JSON(200, gin.H{"code": http.StatusBadRequest, "message": "Upload failed, please contact the administrator"})
	} else {
		shell := "python3 /home/Projects/PaddleOCR/tools/test_hubserving.py http://127.0.0.1:8868/predict/ocr_system " + Path + fileName
		res, _ := other.Command(shell)
		fmt.Println(shell)
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": strings.Replace(res, "\n", "", -1)})
	}
}

func base64ToPicture(c *gin.Context) {

}