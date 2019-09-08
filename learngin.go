package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-common-master/library/database/sql"
	"net/http"
)

func main() {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()

	router.Use(Middleware)

	router.GET("/test/server/get", GetHandler)
	router.POST("/test/server/post", PostHandler)
	router.PUT("/test/server/put", PutHandler)
	router.DELETE("/test/server/delete", DeleteHandler)

	//路由组
	v1 := router.Group("/v1")
	{
		v1.POST("/post", PostHandler)
		//重定向
		v1.GET("/redirect/github", RedirectGitHub)
	}

	v2 := router.Group("v2")
	{
		v2.GET("/student/get", StudentGet)
	}
	//监听端口
	http.ListenAndServe(":8005", router)

}

func Middleware(c *gin.Context) {

	fmt.Println("this is a middleware")
}

func GetHandler(c *gin.Context) {
	v, exist := c.GetQuery("key")
	if !exist {
		v = "the key is not exist!"
	}

	c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("get success! %s\n", v)))

	return
}

func PostHandler(c *gin.Context) {
	type JsonHandler struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	holder := JsonHandler{Id: 1, Name: "my name"}

	c.JSON(http.StatusOK, holder)

	return
}

func PutHandler(c *gin.Context) {
	c.Data(http.StatusOK, "text/plain", []byte("put success!\n"))
	return
}
func DeleteHandler(c *gin.Context) {
	c.Data(http.StatusOK, "text/plain", []byte("delete success!\n"))
	return
}

func RedirectGitHub(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "https://github.com/skyhee/gin-doc-cn")
}

func StudentGet(c *gin.Context) {
	db, err := sql.Open("mysql", "ucenter_t_rw:TVlo_teB3gTxl1RvS@tcp(rm-2zekwe9l6v3tkxixgo.mysql.rds.aliyuncs.com:3306)/ucenter_t_rw")
	result, err := db.QueryRow("SELECT * FROM student limit 1 ")
	PrintArr(result)
	c.Data(http.StatusOK, "text/plain", []byte("success!\n"))
	return
}
