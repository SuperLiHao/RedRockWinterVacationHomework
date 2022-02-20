package DELETE

import (
	"awesomeProject/Utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"os"
)

func GetDeletePostFunc(c *gin.Context) {

	postid := c.Param("post_id")

	userName, success := Utils.JudgeAccessToken(c.GetHeader("Authorrization"))

	if !success {
		c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{"message": "TokenError"}))
		return
	}

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

	delete(Utils.Posts, Post.Post_id)  //unload
	Utils.DeletePostInDataBase(postid) //Delete in Database
	folderPath := "./PostsPicture/" + postid
	os.RemoveAll(folderPath) //Delete In picturesFolder
	fp := "./Posts/" + postid + ".json"
	os.RemoveAll(fp) //Delete JSON

	c.JSON(200, Utils.GetNormalInfo(map[string]interface{}{}))

}
