package GET

import (
	"awesomeProject/Utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Topic struct {
	Topicid      string
	Logourl      string
	Topicname    string
	Introduction string
}

func GetTopicListFunc(c *gin.Context) {

	str, err := Utils.GetStrinFile("./Config/topics.json")

	if err != nil {
		fmt.Println(err)
		c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{"message": "读取主题错误"}))
		return
	}

	ans := make([]Topic, 0)

	json.Unmarshal([]byte(str), &ans)

	c.JSON(200, Utils.GetNormalInfo(map[string]interface{}{"topics": ans}))

	return

}
