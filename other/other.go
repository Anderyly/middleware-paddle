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
	date := time.Now().Format("2006-01-02")
	if ok := IsFileExist(path + "/" + date); !ok {
		os.Mkdir(path+"/"+date, 0666)
	}

	curFileStr := strconv.FormatInt(time.Now().UnixNano(), 10)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(99999)

	var file string = path + "/" + date + "/" + curFileStr + strconv.Itoa(n) + "." + fileType
	byte, _ := base64.StdEncoding.DecodeString(base64Str)

	err := ioutil.WriteFile(file, byte, 0666)
	if err != nil {
		log.Println(err)
	}

	return true, "/" + date + "/" + curFileStr + strconv.Itoa(n) + "." + fileType
}
