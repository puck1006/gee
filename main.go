package main

import (
	"gee"
	"net/http"
)

func main() {

	r := gee.New()

	r.GET("/", func(c *gee.Context) {
		c.Data(200, []byte("haha harden"))
	})
	r.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello is test Query %s \n , and the path is %s", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	r.Run(":9999")
}
