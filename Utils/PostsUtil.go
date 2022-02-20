package Utils

import (
	"encoding/json"
	"mime/multipart"
	"os"
)

type NewPost struct {
	Title   string                 `form:"title" json:"title" binding:"required"`
	Content string                 `form:"content" json:"content" binding:"required"`
	TopicId string                 `form:"topic_id" json:"topic_id" binding:"required"`
	Photo   []multipart.FileHeader `form:"photo" json:"photo" binding:"required"`
}

func CreateNewPostFileFunc(postid string, value NewPost) error {

	f, err := os.Create("./Posts/" + postid + ".json")
	if err != nil {
		return err
	}

	jsonStr, _ := (json.Marshal(value))
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

func UpdatePostFileFunc(postid string, value NewPost) error {

	f, err := os.OpenFile("./Posts/"+postid+".json", 1, 0777)
	if err != nil {
		return err
	}

	jsonStr, _ := (json.Marshal(value))
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
