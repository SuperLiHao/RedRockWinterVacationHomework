package Utils

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	base64Captcha "github.com/mojocn/base64Captcha"
	"io/ioutil"
	"strings"
)

const (
	ImgHeight = 43
	ImgWidth  = 200
	KeyLong   = 5
)

func CreateNewCaptcha() (string, string, error) {

	var store = base64Captcha.DefaultMemStore

	driver := base64Captcha.NewDriverDigit(ImgHeight, ImgWidth, KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)

	if id, b64s, err := cp.Generate(); err != nil {
		return "", "", err
	} else {
		b64s = strings.ReplaceAll(b64s, "data:image/png;base64,", "")
		//fmt.Println(id, "-", b64s)
		//SavePicture(b64s)
		return id, b64s, nil
	}

}

func SavePicture(str string) {
	ddd, _ := base64.StdEncoding.DecodeString(str)
	err := ioutil.WriteFile("./output.jpg", ddd, 0777)
	fmt.Println(err)
}

func VerifyCaptcha(id, VerifyValue string) bool {
	return base64Captcha.DefaultMemStore.Verify(id, VerifyValue, true)
}

func GetNewCptchaFunc(c *gin.Context) {

	id, data, err := CreateNewCaptcha()

	if err != nil {
		c.JSON(403, GetMistakeInfo(map[string]interface{}{"message": "获取验证码失败"}))
		return
	} else {
		c.JSON(200, GetNormalInfo(map[string]interface{}{"id": id, "pngstr": data}))
		return
	}

}
