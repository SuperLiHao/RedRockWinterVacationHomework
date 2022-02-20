package Utils

import (
	"strings"
	"time"
)

func JudgeStr(str string) bool {

	if len(str) == 0 {
		return false
	}

	for _, val := range str {
		if !(int(val) > 1 && int(val) < 128 && int(val) != 32) {
			return false
		}
	}
	return true

}

func JudgeNum(str string) bool {

	if len(str) == 0 {
		return false
	}

	for _, val := range str {
		if !(int(val) >= '0' && int(val) <= '9') {
			return false
		}
	}

	return true

}

func NewUserInfo(nickname string, introducion string, phone string, qq string, email string, birthday string) UserInfo {

	var user UserInfo

	user.Nickname = nickname
	user.Introduction = introducion
	user.Phone = phone
	user.QQ = qq
	user.Email = email
	user.Birthday = birthday

	return user

}

func GetUserAvatorFunc(username string) string {
	return "./Avatar/" + username + ".png"
}

func GetNowTime() string {

	t1 := time.Now().Year()
	t2 := time.Now().Month()
	t3 := time.Now().Day()
	t4 := time.Now().Hour()
	t5 := time.Now().Minute()
	t6 := time.Now().Second()
	t7 := time.Now().Nanosecond()

	currentTimeData := time.Date(t1, t2, t3, t4, t5, t6, t7, time.Local)

	return currentTimeData.String()[:19]

}

func IsAdministrator(userName string) bool {

	return false //帮助测试
}

func IsContainStr(post Post, key string) bool {

	if strings.Contains(post.Content, key) {
		return true
	}
	if strings.Contains(post.Nickname, key) {
		return true
	}
	if strings.Contains(post.Name, key) {
		return true
	}
	if strings.Contains(post.User_id, key) {
		return true
	}

	return false

}
