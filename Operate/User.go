package Operate

import (
	"awesomeProject/Utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"strconv"
)

type UserInfo struct {
	PostPraise    []string
	CommentPraise []string
	Collect       []string
	Focus         []string
}

func GetPraisePostFunc(c *gin.Context) {

	userName, success := Utils.JudgeAccessToken(c.GetHeader("Authorrization"))

	if !success {
		c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{"message": "TokenError"}))
		return
	}

	model := c.PostForm("model")
	target := c.PostForm("target_id")

	if len(model) == 0 || len(target) == 0 {
		c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{"message": "缺少参数"}))
		return
	}

	modelNum, err := strconv.Atoi(model)

	if err != nil || !(modelNum == 1 || modelNum == 2) {
		c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{"message": "参数错误"}))
		return
	}

	data, err := GetUserInfo(userName)

	if err != nil {
		fmt.Println(err)
	}

	if modelNum == 1 {

		jsonStr, _ := (json.Marshal(Utils.Posts[target]))
		js := string(jsonStr)

		if len(js) == 176 {
			c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{"message": "帖子不存在"}))
			return
		}
		data.PostPraise = append(data.PostPraise, target)
	} else {

		data.CommentPraise = append(data.CommentPraise, target)

	}

	WriteUserInfo(userName, data)
	//data.
	c.JSON(200, Utils.GetNormalInfo(map[string]interface{}{}))
} //点赞

func WriteUserInfo(userName string, user UserInfo) error {

	userIn := Utils.GetUser(userName).UUID

	f, err := os.OpenFile("./Users/"+userIn+".json", 1, 0777)
	if err != nil {
		return err
	}

	jsonStr, _ := (json.Marshal(user))
	js := string(jsonStr)

	_, err = f.WriteString(js)
	if err != nil {
		f.Close()
		return err
	}
	//fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		return err
	}

	return nil

}

func GetUserInfo(userName string) (UserInfo, error) {

	userIn := Utils.GetUser(userName).UUID
	var user UserInfo

	data, _ := ioutil.ReadFile("./Users/" + userIn + ".json")

	json.Unmarshal(data, &user)

	return user, nil

}

func GetCollectListFunc(c *gin.Context) {

	userName, success := Utils.JudgeAccessToken(c.GetHeader("Authorrization"))

	if !success {
		c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{"message": "TokenError"}))
		return
	}

	var ans []Utils.Post

	userIn, err := GetUserInfo(userName)

	if err != nil {
		fmt.Println(err)
	}

	for _, v := range userIn.Collect {

		ans = append(ans, Utils.Posts[v])

	}

	c.JSON(200, Utils.GetNormalInfo(map[string]interface{}{"collections": ans}))
	return
} //收藏列表

func GetFocusUserFunc(c *gin.Context) {

	userName, success := Utils.JudgeAccessToken(c.GetHeader("Authorrization"))

	if !success {
		c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{"message": "TokenError"}))
		return
	}

	user_id := c.PostForm("user_id")

	if len(user_id) == 0 || !Utils.FindUser(user_id) {
		c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{"messgae": "No user"}))
		return
	}

	userIn, _ := GetUserInfo(userName)

	userIn.Focus = append(userIn.Focus, user_id)

	WriteUserInfo(userName, userIn)

	c.JSON(200, Utils.GetNormalInfo(map[string]interface{}{}))

	return

} //关注用户

func GetCollectUserFunc(c *gin.Context) {

	userName, success := Utils.JudgeAccessToken(c.GetHeader("Authorrization"))

	if !success {
		c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{"message": "TokenError"}))
		return
	}

	user_id := c.PostForm("post_id")

	if len(user_id) == 0 {
		c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{}))
		return
	}

	userIn, _ := GetUserInfo(userName)

	userIn.Collect = append(userIn.Collect, user_id)

	WriteUserInfo(userName, userIn)

	c.JSON(200, Utils.GetNormalInfo(map[string]interface{}{}))

	return

} //收藏帖子

func GetFocusListFunc(c *gin.Context) {

	userName, success := Utils.JudgeAccessToken(c.GetHeader("Authorrization"))

	if !success {
		c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{"message": "TokenError"}))
		return
	}

	userIn, _ := GetUserInfo(userName)

	ans := make([]Utils.User, 0)

	for _, uName := range userIn.Focus {
		ans = append(ans, Utils.GetUser(uName))
	}

	c.JSON(200, Utils.GetNormalInfo(map[string]interface{}{"users": ans}))

	return

} //关注列表
