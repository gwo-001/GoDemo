package main

import (
	"github.com/gin-gonic/gin"
	"helloGin/handler"
	"net/http"
)



func main() {
	r:=gin.Default()

	r.GET("/get",func (c *gin.Context)  {
		name :=c.DefaultQuery("name","")
		age,ok:=c.GetQuery("age")
		if ok {
			c.JSON(200,gin.H{"帅比":name,"age":age,})
		}else {
			c.JSON(http.StatusOK,gin.H{"name":name})
		}
	})
	r.POST("/post",handler.PostTest)

	r.Run(":8080")
}
