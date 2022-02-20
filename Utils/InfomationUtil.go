package Utils

import "github.com/gin-gonic/gin"

func GetNormalInfo(info map[string]interface{}) gin.H {

	return gin.H{
		"state": "10000",
		"info":  "success",
		"data":  info,
	}
}

func GetErrorInfo(info map[string]interface{}) gin.H {

	return gin.H{
		"state": "10002",
		"info":  "error",
		"data":  info,
	}
}

func GetMistakeInfo(info map[string]interface{}) gin.H {

	return gin.H{
		"state": "10001",
		"info":  "failed",
		"data":  info,
	}

}
