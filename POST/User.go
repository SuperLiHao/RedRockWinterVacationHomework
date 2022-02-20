package POST

import (
	"awesomeProject/Operate"
	Utils "awesomeProject/Utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

type RegUser struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Email    string `form:"email"    json:"email"    binding:"-"`
}

func GetUserRegisterFunc(c *gin.Context) {
	var reg RegUser
	if err := c.ShouldBind(&reg); err != nil {
		c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{"message": "缺少参数"}))
		return
	}

	if !Utils.JudgeStr(reg.Password) {
		c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{"message": "密码不合法"}))
		return
	}

	if Utils.FindUser(reg.Username) {
		c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{"message": "账号已经存在"}))
	} else {
		Token, TokenError := Utils.GetToken(reg.Username)
		Refresh_token, RTokenError := Utils.GetRefreshToken(reg.Username)

		if TokenError != nil || RTokenError != nil {
			c.JSON(403, Utils.GetErrorInfo(map[string]interface{}{"message": "Token生成出错"}))
		} else {
			Utils.AddUser(reg.Username, reg.Password, reg.Email)
			c.JSON(http.StatusOK, Utils.GetNormalInfo(map[string]interface{}{"token": Token, "refresh_token": Refresh_token}))
			Utils.WriteRefreshToken(reg.Username, time.Now().Add(Utils.RefreshTokenTime).Unix())
			_, err := os.Create("./Users/" + Utils.GetUser(reg.Username).UUID + ".json")
			Operate.WriteUserInfo(reg.Username, Operate.UserInfo{})
			if err != nil {
				fmt.Println(err)
			}
		}

	}

}
