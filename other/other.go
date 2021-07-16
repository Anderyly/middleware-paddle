package other

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"time"
)

func Command(cmd string) (string, error) {
	c := exec.Command("bash", "-c", cmd)
	output, err := c.CombinedOutput()

	return string(output), err
}

//IsFileExist
func IsFileExist(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true

}

//WriteFile base64 to picture
func WriteFile(path string, base64_image_content string) (bool, string) {
	b, _ := regexp.MatchString(`^data:\s*image\/(\w+);base64,`, base64_image_content)
	if !b {
		return false, ""
	}
	re, _ := regexp.Compile(`^data:\s*image\/(\w+);base64,`)
	allData := re.FindAllSubmatch([]byte(base64_image_content), 2)
	fileType := string(allData[0][1]) //png ，jpeg 后缀获取
	base64Str := re.ReplaceAllString(base64_image_content, "")
	file := path + "/" + strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(999999-100000)+100000) + "." + fileType
	byte, _ := base64.StdEncoding.DecodeString(base64Str)

	err := ioutil.WriteFile(file, byte, 0666)
	if err != nil {
		log.Println(err)
	}

	return true, file
}
