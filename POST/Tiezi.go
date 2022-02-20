package POST

import (
	"awesomeProject/Utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path"
)

type NewPost struct {
	Title    string `form:"title" json:"title" binding:"required"`
	Content  string `form:"content" json:"content" binding:"required"`
	TopicId  string `form:"topic_id" json:"topic_id" binding:"required"`
	Pictures []string
}

type Post struct {
	Name         string //Title
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

func GetPostPostFunc(c *gin.Context) {

	userName, success := Utils.JudgeAccessToken(c.GetHeader("Authorrization"))

	if !success {
		c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{"message": "TokenError"}))
		return
	}

	var newPost Utils.NewPost

	form, _ := c.MultipartForm()
	files, _ := form.File["photo"]

	if err := c.ShouldBind(&newPost); err != nil {
		c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{"message": "缺少参数"}))
		return
	}

	postid := Utils.GetNewPostId()

	if err := Utils.CreateNewPostFileFunc(postid, newPost); err != nil {
		c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{"message": "上传失败"}))
		return
	} //生成json文件

	if err := Utils.CreatePostInDataBase(newPost, postid, userName); err != nil {
		fmt.Println(err)
	} //写入数据库

	ps := make([]string, cap(files))
	{
		folderPath := "./PostsPicture/" + postid
		os.Mkdir(folderPath, 0777)
		os.Chmod(folderPath, 0777) //创建文件夹并分配权限
		commentFolderPictures := "./CommentsSource/CommentsPictures/" + postid
		os.Mkdir(commentFolderPictures, 0777)
		os.Chmod(commentFolderPictures, 0777)
		num := 0

		for _, v := range files {
			if err := c.SaveUploadedFile(v, path.Join("./PostsPicture/"+postid, v.Filename)); err != nil {
				fmt.Println(files)
			}
			ps[num] = "./PostsPicture/" + postid + "/" + v.Filename
			num++
		}
	} //保存图片

	Utils.Posts[postid] = Utils.Post{
		Name:         newPost.Title,
		Post_id:      postid,
		Public_time:  Utils.GetNowTime(),
		Content:      newPost.Content,
		Pictures:     ps,
		Topic_id:     newPost.TopicId,
		User_id:      userName,
		Avatar:       "./Avatar/" + userName + ".png",
		Nickname:     Utils.GetUser(userName).Nickname,
		Praise_count: 0,
	} //Loadin

	c.JSON(200, Utils.GetNormalInfo(map[string]interface{}{"post_id": postid}))

	return

}
