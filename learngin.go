package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
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

func PostHandler(c *gin.Context)  {
	type JsonHandler struct {
		Id int `json:"id"`
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
