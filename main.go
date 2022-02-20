package main

import (
	"awesomeProject/Comment"
	"awesomeProject/DELETE"
	"awesomeProject/GET"
	"awesomeProject/Operate"
	"awesomeProject/POST"
	"awesomeProject/PUT"
	DataBaseUtil "awesomeProject/Utils"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/unrolled/secure"
	"net/http"
)

func main() {

	r := gin.Default()

	DataBaseUtil.CreateNewCaptcha()

	//Init
	{
		{
			err := DataBaseUtil.LoadAllPosts()
			if err != nil {
				fmt.Println(err)
			}
		} //LoadPosts
	}

	r.LoadHTMLGlob("HTML/*")
	r.Static("/js", "./HTML")

	//PAGE
	{
		r.GET("/", func(c *gin.Context) {

			c.HTML(http.StatusOK, "index.html", gin.H{

				"state": "10000",
				"info":  "success",
				"data":  DataBaseUtil.FineTiezis(),
			})

		})                                                                                         //主页
		r.GET("/login", func(c *gin.Context) { c.HTML(http.StatusOK, "login.html", gin.H{}) })     //登录
		r.GET("/registe", func(c *gin.Context) { c.HTML(http.StatusOK, "registe.html", gin.H{}) }) //注册
	}

	//Captcha
	{
		r.GET("/captcha", DataBaseUtil.GetNewCptchaFunc)
	}

	// /user/……
	{
		userGroup := r.Group("/user")
		//GET
		{
			userGroup.GET("/token", GET.GetUserTokenFunc)            //Login
			userGroup.GET("/token/refresh", GET.GetRefreshTokenFunc) //refresh
			userGroup.GET("/info/:id", GET.GetUserFunc)              //FindUser
		}
		//POST
		{
			userGroup.POST("/register", POST.GetUserRegisterFunc) //register
		}

		//PUT
		{
			userGroup.PUT("/password", PUT.GetChangePasswordFunc) //changePassword
			userGroup.PUT("/info", PUT.GetChangeInfoFunc)         //ChangeInfomation
		}
	}

	// /post/……
	{
		post := r.Group("/post")

		//GET
		{
			post.GET("/list", GET.GetPageListFunc)
			post.GET("/single/:post_id", GET.GetFindPostFunc)
			post.GET("/search", GET.GetSearchPostsFunc)
		}
		//POST
		{
			post.POST("/single", POST.GetPostPostFunc)
		}
		//PUT
		{
			post.PUT("/single/:post_id", PUT.GetUpdatePostFunc)
		}
		//DELETE
		{
			post.DELETE("/single/:post_id", DELETE.GetDeletePostFunc)
		}
	}

	// /topic/……
	{
		r.GET("/topic/list", GET.GetTopicListFunc)
	}

	// /operate/……
	{
		operate := r.Group("/operate")
		//PUT
		{
			operate.PUT("/praise", Operate.GetPraisePostFunc)
			operate.PUT("/focus", Operate.GetFocusUserFunc)
			operate.PUT("/collect", Operate.GetCollectUserFunc)
		}
		//GET
		{
			operate.GET("/collect/list", Operate.GetCollectListFunc)
			operate.GET("/focus/list", Operate.GetFocusListFunc)
		}
	}

	// /comment
	{
		comment := r.Group("/comment")

		comment.POST("", Comment.GetPostCommentFunc)
		comment.GET("", Comment.GetGetCommentsFunc)
		comment.PUT("/:comment_id", Comment.GetUpdateCommentFunc)
		comment.DELETE("/:comment_id", Comment.GETDeleteCommentFunc)
	}

	{
		r.Use(TlsHandler())

		r.RunTLS(":443", "./HTML/https.cer", "./HTML/https.key")

	} //https

}

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     ":443",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)
		if err != nil {
			return
		}

		c.Next()
	}
}
