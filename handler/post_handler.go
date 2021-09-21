package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"helloGin/model"
	"net/http"
)

func PostTest(c *gin.Context) {
	defer fmt.Println("post test is finished")

	var postParamModel model.PostTestModel

	name := c.PostForm("name")
	age, err := c.GetPostForm("age")

	if err := c.ShouldBindJSON(&postParamModel); err != nil {

	}

	if err {
		fmt.Printf("the age field is empty")
		return
	}
	fmt.Printf("name is %s\n, age is %s\n", name, age)
	c.JSON(http.StatusOK,gin.H{
		"code":http.StatusOK,
		"message":"insert 2 DB SUCCESS",

	})
}
