package Utils

import (
	"database/sql"
	"errors"
	"fmt"
)

type Post struct {
	Name         string
	Post_id      string
	Public_time  string
	Content      string
	Pictures     []string
	Topic_id     string
	User_id      string
	Avatar       string
	Nickname     string
	Praise_count int
	Is_focus     bool
	Is_praised   bool
}

var Posts map[string]Post

func LoadAllPosts() error {

	Posts = make(map[string]Post)

	db, err := sql.Open("mysql", SQLurl+SQLName)

	if err != nil {
		fmt.Println(err)
		return errors.New("connection Error")
	}

	defer db.Close()

	sqlStr := "select name,postid,publictime,content,Topic_id,user_id,praise_count from " + TieziExcel

	rows, err := db.Query(sqlStr)

	if err != nil {
		fmt.Println(err)
	}

	var p Post

	for rows.Next() {

		err = rows.Scan(&p.Name, &p.Post_id, &p.Public_time, &p.Content, &p.Topic_id, &p.User_id, &p.Praise_count)
		if err != nil {
			fmt.Println(err)
		} else {
			user := GetUser(p.User_id)
			p.Avatar = GetUserAvatorFunc(user.Username)
			p.Pictures = GetPostPictures(p.Post_id)
			p.Nickname = user.Nickname
			Posts[p.Post_id] = p
		}

	}

	return nil

}

func CreatePostInDataBase(post NewPost, postid string, userName string) error {

	db, err := sql.Open("mysql", SQLurl+SQLName)

	defer db.Close()

	if err != nil {
		return err
	}

	sqlStr := "insert into " + TieziExcel + " (name,publictime,postid,content,topic_id,user_id,praise_count) value(?,?,?,?,?,?,?)"

	_, err = db.Exec(sqlStr, post.Title, GetNowTime(), postid, post.Content, post.TopicId, userName, 0)

	if err != nil {
		return err
	} else {
		return nil
	}

}

func UpdatePostInDataBase(post NewPost, postid string, userName string) error {

	db, err := sql.Open("mysql", SQLurl+SQLName)

	defer db.Close()

	if err != nil {
		return err
	}

	sqlStr := "update " + TieziExcel + " set name=?,content=?,topic_id=?,user_id=? where postid = ?"

	_, err = db.Exec(sqlStr, post.Title, post.Content, post.TopicId, userName, postid)

	if err != nil {
		return err
	} else {
		return nil
	}

}

func DeletePostInDataBase(postid string) error {

	db, err := sql.Open("mysql", SQLurl+SQLName)

	defer db.Close()

	if err != nil {
		return err
	}

	sqlStr := "delete from " + TieziExcel + " where postid = ?"

	_, err = db.Exec(sqlStr, postid)

	if err != nil {
		return err
	} else {
		return nil
	}

}
