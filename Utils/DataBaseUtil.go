package Utils

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

const SQLurl = "root:214315487@tcp(localhost:3306)/"
const SQLName = "redrockbbs"
const UsersExcel = "users"
const TieziExcel = "tiezis"
const TokenExcel = "refreshtoken"

type User struct {
	Username     string `form:"username" json:"username" binding:"required"`
	Password     string `form:"password" json:"password" binding:"required"`
	Email        string `form:"email"    json:"email"    binding:"required"`
	UUID         string
	QQ           string
	Avatar       string
	Phone        string
	Nickname     string
	Introduction string
	Birthday     string
}

type UserInfo struct {
	Email        string
	QQ           string
	Avatar       string
	Phone        string
	Nickname     string
	Introduction string
	Birthday     string
}

func AddUser(userName string, password string, email string) bool {

	db, err := sql.Open("mysql", SQLurl+SQLName)

	defer db.Close()

	if err != nil {
		fmt.Println(err)
		return false
	}

	sqlStr := "insert into " + UsersExcel + "(username,password,email,UUID) value(?,?,?,?)"

	_, err = db.Exec(sqlStr, userName, password, email, GetNewUUID())

	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}

}

func FindUser(userName string) bool {

	db, err := sql.Open("mysql", SQLurl+SQLName)

	defer db.Close()

	if err != nil {
		fmt.Println(err)
		return false
	}

	sqlStr := "select username,password,email from " + UsersExcel + " where username ='" + userName + "'"

	var password string
	var email string
	var username string

	err = db.QueryRow(sqlStr).Scan(&username, &password, &email)

	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}

}

func GetUser(userName string) User {

	var user User

	db, err := sql.Open("mysql", SQLurl+SQLName)

	defer db.Close()

	if err != nil {
		fmt.Println(err)
	}

	sqlStr := "select username,password,email,UUID,QQ,avatar,phone,nickname,introduction,birthday from " + UsersExcel + " where username = ?"

	err = db.QueryRow(sqlStr, userName).Scan(&user.Username, &user.Password, &user.Email, &user.UUID, &user.QQ, &user.Avatar, &user.Phone, &user.Nickname, &user.Introduction, &user.Birthday)

	if err != nil {
		user.Username = "ERROR"
	}

	return user

}

func FineTiezis() map[string]interface{} {

	db, err := sql.Open("mysql", SQLurl+SQLName)

	if err != nil {
		fmt.Println(err)
		return gin.H{}
	}

	ans := gin.H{}

	defer db.Close()

	sqlStr := "select name,time,dianzan from " + TieziExcel

	rows, err := db.Query(sqlStr)

	if err != nil {
		fmt.Println(err)
	}

	var num int64
	num = 1

	for rows.Next() {

		var Name string
		var Time string
		var DianZan int
		err = rows.Scan(&Name, &Time, &DianZan)
		if err != nil {
			fmt.Println(err)
		} else {
			ans["No"+strconv.FormatInt(num, 10)] = map[string]interface{}{
				"name":    Name,
				"time":    Time,
				"dianzan": DianZan,
			}
			num += 1
		}

	}

	return ans

}

func ChangePassword(userName string, password string) bool {

	db, err := sql.Open("mysql", SQLurl+SQLName)

	defer db.Close()

	if err != nil {
		fmt.Println(err)
		return false
	}

	sqlStr := "update " + UsersExcel + " set password = ? where username = ?"

	_, err = db.Exec(sqlStr, password, userName)

	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}

}

func WriteRefreshToken(userName string, time int64) bool {

	db, err := sql.Open("mysql", SQLurl+SQLName)

	defer db.Close()

	if err != nil {
		fmt.Println(err)
		return false
	}

	sqlStr := "update " + TokenExcel + " set time = ? where username = ?"

	if GetRefreshTokenTime(userName) == -1 {
		sqlStr = "insert into " + TokenExcel + " (time,username) VALUES (?,?)"
	}

	_, err = db.Exec(sqlStr, time, userName)

	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}

}

func GetRefreshTokenTime(userName string) int64 {

	db, err := sql.Open("mysql", SQLurl+SQLName)

	defer db.Close()

	if err != nil {
		fmt.Println(err)
		return -1
	}

	sqlStr := "select time from " + TokenExcel + " where username ='" + userName + "'"

	var time int64

	err = db.QueryRow(sqlStr).Scan(&time)

	if err != nil {
		fmt.Println(err)
		return -1
	} else {
		return time
	}

}

func GetUserByUUID(uuid string) User {

	var user User

	db, err := sql.Open("mysql", SQLurl+SQLName)

	defer db.Close()

	if err != nil {
		fmt.Println(err)
	}

	sqlStr := "select username,password,email,UUID,QQ,avatar,phone,nickname,introduction,birthday from " + UsersExcel + " where UUID ='" + uuid + "'"

	err = db.QueryRow(sqlStr).Scan(&user.Username, &user.Password, &user.Email, &user.UUID, &user.QQ, &user.Username, &user.Phone, &user.Nickname, &user.Introduction, &user.Birthday)

	if err != nil {
		user.Username = "ERROR"
	}

	return user

}

func SetUserInfoFunc(username string, key string, value string) bool {

	db, err := sql.Open("mysql", SQLurl+SQLName)

	defer db.Close()

	if err != nil {
		fmt.Println(err)
		return false
	}

	sqlStr := "update " + UsersExcel + " set " + key + " = ? where username = ?"

	_, err = db.Exec(sqlStr, value, username)

	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}

}
