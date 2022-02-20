package GET

import (
	Utils "awesomeProject/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func GetUserTokenFunc(c *gin.Context) {

	id := c.PostForm("id")
	value := c.PostForm("value")

	if !Utils.VerifyCaptcha(id, value) {
		c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{"message": "验证码错误"}))
		return
	}

	var login User

	if err := c.ShouldBind(&login); err != nil {
		c.JSON(403, Utils.GetErrorInfo(map[string]interface{}{"message": "出现了不可控的意外"}))
		return
	}

	if Utils.FindUser(login.Username) {
		if Utils.GetUser(login.Username).Password == login.Password {

			Token, TokenError := Utils.GetToken(login.Username)
			Refresh_token, RTokenError := Utils.GetRefreshToken(login.Username)

			if TokenError != nil || RTokenError != nil {
				c.JSON(403, Utils.GetErrorInfo(map[string]interface{}{"message": "Token生成出错"}))
			} else {
				c.JSON(http.StatusOK, Utils.GetNormalInfo(map[string]interface{}{"token": Token, "refresh_token": Refresh_token}))
				//fmt.Println(Utils.PraseToken(Token))
				//fmt.Println(Utils.PraseToken(Refresh_token))
				Utils.WriteRefreshToken(login.Username, time.Now().Add(Utils.RefreshTokenTime).Unix())
				//fmt.Println(Utils.GetRefreshTokenTime(login.Username))
			}

		} else {
			c.JSON(http.StatusForbidden, Utils.GetMistakeInfo(map[string]interface{}{"message": "密码错误"}))
		}

	} else {
		c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{"message": "用户不存在"}))
	}

}

func GetRefreshTokenFunc(c *gin.Context) {

	refresh_token := c.PostForm("refresh_token")

	if len(refresh_token) == 0 {
		c.JSON(403, Utils.GetErrorInfo(map[string]interface{}{"message": "缺少参数"}))
		return
	}
	theTime, _ := Utils.PraseToken(refresh_token)

	if theTime.TokenType != "RefreshToken" {
		c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{"message": "参数错误"}))
		return
	}

	t := Utils.GetRefreshTokenTime(theTime.Username)

	if t <= time.Now().Unix() {
		c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{"message": "refresh_token已过期"}))
		return
	}

	Token, TokenError := Utils.GetToken(theTime.Username)
	Refresh_token, RTokenError := Utils.GetRefreshToken(theTime.Username)

	if TokenError != nil || RTokenError != nil {
		c.JSON(403, Utils.GetErrorInfo(map[string]interface{}{"message": "Token生成出错"}))
	} else {
		c.JSON(http.StatusOK, Utils.GetNormalInfo(map[string]interface{}{"token": Token, "refresh_token": Refresh_token}))
		Utils.WriteRefreshToken(theTime.Username, time.Now().Add(Utils.RefreshTokenTime).Unix())
	}

}

func GetUserFunc(c *gin.Context) {

	id := c.Param("id")

	var user Utils.User

	if !Utils.FindUser(id) {
		user = Utils.GetUserByUUID(id)
	} else {
		user = Utils.GetUser(id)
	}

	if user.Username == "ERROR" {
		c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{"message": "该用户不存在"}))
	} else {

		c.JSON(200, Utils.GetNormalInfo(map[string]interface{}{"user": user}))

	}

}
