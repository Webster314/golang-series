package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main(){
    r := gin.Default();

    r.GET("/", func(c * gin.Context){
        c.String(200, "hello gin\n");
    })

    r.POST("/form", func(c * gin.Context){
        username := c.PostForm("username");
        password := c.DefaultPostForm("password", "000000")
        c.JSON(http.StatusOK, gin.H{
            "username" : username,
            "password" : password,
        })
    })

    r.GET("/redirect", func(c * gin.Context){
        c.Redirect(http.StatusMovedPermanently, "index");
    })

    r.GET("/goindex", func(c * gin.Context){
        c.Request.URL.Path = "/"
        r.HandleContext(c);
    })

    v1 := r.Group("/v1")
    {
        v1.GET("/get", func(c * gin.Context){

        })
        v1.GET("/series", func(c * gin.Context){

        })
    }
    v2 := r.Group("/v2")
    {
        v2.GET("/get", func(c * gin.Context){
        });
    }
    r.Run()
}
