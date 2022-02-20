package PUT

import (
	"awesomeProject/Utils"
	"github.com/gin-gonic/gin"
	"path"
)

type ChangePassword struct {
	old_password string `form:"old_password" json:"old_password" binding:"required"`
	Password     string `form:"password" json:"password" binding:"required"`
	Email        string `form:"email"    json:"email"    binding:"-"`
}

func GetChangePasswordFunc(c *gin.Context) {

	userName, success := Utils.JudgeAccessToken(c.GetHeader("Authorrization"))

	old_Password := c.PostForm("old_password")
	new_password := c.PostForm("new_password")

	if !success {
		c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{"message": "TokenError"}))
		return
	} else {

		user := Utils.GetUser(userName)

		if old_Password != user.Password {
			c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{"message": "old_password mismatch"}))
			return
		}
		if !Utils.JudgeStr(new_password) {
			c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{"message": "新密码不合法"}))
			return
		}

		success = Utils.ChangePassword(userName, new_password)

		if success {
			c.JSON(200, Utils.GetNormalInfo(map[string]interface{}{}))
		} else {
			c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{"message": "修改密码失败"}))
		}

	}

}

func GetChangeInfoFunc(c *gin.Context) {

	userName, success := Utils.JudgeAccessToken(c.GetHeader("Authorrization"))

	if !success {
		c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{"message": "TokenError"}))
		return
	}

	mistake := make(map[string]interface{})

	avatar, err := c.FormFile("avator")

	if err == nil {

		if er := c.SaveUploadedFile(avatar, path.Join("./Avatar", userName+".png")); er != nil {
			mistake["avator"] = "Eroor"
		}

	} //Avatar

	nickname := c.PostForm("nickname")
	introduction := c.PostForm("introduction")
	phone := c.PostForm("telephone")
	qq := c.PostForm("qq")
	email := c.PostForm("email")
	birthday := c.PostForm("birthday")

	if len(nickname) != 0 {
		if !Utils.SetUserInfoFunc(userName, "nickname", nickname) {
			mistake["nickname"] = "error"
		}
	}
	if len(introduction) != 0 {
		if !Utils.SetUserInfoFunc(userName, "introduction", introduction) {
			mistake["introduction"] = "error"
		}
	}
	if len(phone) != 0 {

		if !Utils.JudgeNum(phone) {
			mistake["telephone"] = "error"
		} else if !Utils.SetUserInfoFunc(userName, "telephone", phone) {
			mistake["telephone"] = "error"
		}
	}
	if len(qq) != 0 {
		if !Utils.JudgeNum(qq) {
			mistake["qq"] = "error"
		} else if !Utils.SetUserInfoFunc(userName, "qq", qq) {
			mistake["qq"] = "error"
		}
	}
	if len(email) != 0 {
		if !Utils.JudgeStr(email) {
			mistake["email"] = "error"
		} else if !Utils.SetUserInfoFunc(userName, "email", email) {
			mistake["email"] = "error"
		}
	}
	if len(birthday) != 0 {
		if !Utils.JudgeStr(birthday) {
			mistake["birthday"] = "error"
		} else if !Utils.SetUserInfoFunc(userName, "birthday", birthday) {
			mistake["birthday"] = "error"
		}
	}

	if len(mistake) == 0 {
		c.JSON(200, Utils.GetNormalInfo(map[string]interface{}{}))
	} else {
		c.JSON(403, Utils.GetMistakeInfo(mistake))
	}

}
