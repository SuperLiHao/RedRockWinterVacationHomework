package PUT

import (
	"awesomeProject/Utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path"
)

func GetUpdatePostFunc(c *gin.Context) {

	userName, success := Utils.JudgeAccessToken(c.GetHeader("Authorrization"))

	if !success {
		c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{"message": "TokenError"}))
		return
	}

	postid := c.Param("post_id")

	jsonStr, _ := (json.Marshal(Utils.Posts[postid]))
	js := string(jsonStr)

	if len(js) == 176 {
		c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{"message": "帖子不存在或已被删除"}))
		return
	}

	Post := Utils.Posts[postid]

	if Post.User_id != userName && !Utils.IsAdministrator(userName) {
		c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{"message": "无权限"}))
		return
	}

	var newPost Utils.NewPost

	form, _ := c.MultipartForm()
	files, _ := form.File["photo"]

	if err := c.ShouldBind(&newPost); err != nil {
		c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{"message": "缺少参数"}))
		return
	}

	if err := Utils.UpdatePostFileFunc(postid, newPost); err != nil {
		fmt.Println(err)
		c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{"message": "上传失败"}))
		return
	} //生成json文件

	if err := Utils.UpdatePostInDataBase(newPost, postid, userName); err != nil {
		fmt.Println(err)
	} //写入数据库

	ps := make([]string, cap(files))
	{
		folderPath := "./PostsPicture/" + postid
		os.RemoveAll(folderPath)
		os.Mkdir(folderPath, 0777)
		os.Chmod(folderPath, 0777) //创建文件夹并分配权限
		num := 0

		for _, v := range files {
			if err := c.SaveUploadedFile(v, path.Join("./PostsPicture/"+postid, v.Filename)); err != nil {
				fmt.Println(err)
			}
			ps[num] = v.Filename
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
