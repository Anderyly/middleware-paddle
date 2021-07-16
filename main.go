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

const PanddlePath = "/home/Projects/PaddleOCR"

var Path string

func main() {
	start()
	r := gin.Default()
	r.POST("/api/picture", picture)
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

	cType := c.PostForm("type")

	fileName := ""
	if cType == "1" {
		res := uploadFile(fileName, c)
		if res == "" {
			return
		}
	} else {
		base := c.PostForm("base")
		_, fileName = other.WriteFile(Path, base)
		fmt.Println(fileName)
	}

	shell := `python3 ` + PanddlePath +
		`PaddleOCR/tools/infer/predict_system.py --image_dir="` + Path + fileName +
		`" --det_model_dir="` + PanddlePath +
		`/inference/ch_ppocr_mobile_v2.0_det_infer/"  --rec_model_dir="` + PanddlePath +
		`/inference/ch_ppocr_mobile_v2.0_rec_infer/" --cls_model_dir="` + PanddlePath +
		`/inference/ch_ppocr_mobile_v2.0_cls_infer/" --use_angle_cls=True --use_space_char=True --use_gpu=False`
	res, _ := other.Command(shell)
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "success", "data": strings.Replace(res, "\\", ``, -1)})

}

func uploadFile(fileName string, c *gin.Context) string {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "msg": "Please Upload"})
		return ""
	}
	fileExt := strings.ToLower(path.Ext(file.Filename))
	if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" {
		c.JSON(200, gin.H{"code": http.StatusBadRequest, "msg": "Upload file suffix not allowed"})
		return ""
	}

	fileName = strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(999999-100000)+100000) + path.Ext(file.Filename)

	upFile := c.SaveUploadedFile(file, Path+fileName)
	if upFile != nil {
		c.JSON(200, gin.H{"code": http.StatusBadRequest, "msg": "Upload failed, please contact the administrator"})
	}
	return fileName
}
