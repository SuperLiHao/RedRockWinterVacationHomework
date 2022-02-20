package GET

import (
	"awesomeProject/Utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetPageListFunc(c *gin.Context) {

	pageN := c.Query("page")
	pageS := c.Query("size")

	if len(pageS) == 0 || len(pageN) == 0 {
		c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{"message": "参数错误"}))
		return
	}

	pageNum, err := strconv.Atoi(pageN)
	if err != nil {
		c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{"message": "参数错误"}))
		return
	}
	pageSize, er := strconv.Atoi(pageS)
	if er != nil {
		c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{"message": "参数错误"}))
		return
	}

	ans := make([]Utils.Post, pageSize)

	num := 0
	jl := 0

	for _, v := range Utils.Posts {

		num++

		fmt.Println(v)

		if num >= (pageNum-1)*pageSize && num < pageSize*pageNum {
			ans[jl] = v
			jl++
		}

		if num >= pageSize*pageNum {
			break
		}

	}

	c.JSON(200, Utils.GetNormalInfo(map[string]interface{}{"posts": ans}))

}

func GetFindPostFunc(c *gin.Context) {

	name := c.Param("post_id")

	jsonStr, _ := (json.Marshal(Utils.Posts[name]))
	js := string(jsonStr)

	if len(js) == 176 {
		c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{"message": "No POST"}))
		return
	}

	c.JSON(200, Utils.GetNormalInfo(map[string]interface{}{"posts": Utils.Posts[name]}))
	return
}

func GetSearchPostsFunc(c *gin.Context) {

	key := c.Query("key")
	page := c.Query("page")
	size := c.Query("size")

	pageNum, _ := strconv.Atoi(page)
	sizeNum, _ := strconv.Atoi(size)

	if pageNum == 0 || sizeNum == 0 {
		c.JSON(403, Utils.GetMistakeInfo(map[string]interface{}{"message": "参数错误"}))
		return
	}

	//(pageNum-1)*sizeNum -> pageNum*sizeNum

	ans := make([]Utils.Post, 0)
	num := 0

	for _, v := range Utils.Posts {

		num++

		if num > (pageNum-1)*sizeNum && num < pageNum*sizeNum {
			if Utils.IsContainStr(v, key) {
				ans = append(ans, v)
			}
		}

	}

	c.JSON(200, Utils.GetNormalInfo(map[string]interface{}{"posts": ans}))

}
