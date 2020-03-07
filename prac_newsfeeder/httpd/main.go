package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"prac_newsfeeder/httpd/handler"
	"prac_newsfeeder/platform/newsfeed"
)

func main() {
	fmt.Println("# # # Run pran_newsfeeder main.go # # #")

	feed := newsfeed.New()
	r := gin.Default()

	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsfeedGet(feed))
	r.POST("/newsfeed", handler.NewsfeedPost(feed))

	r.Run()

	/*
		feed := newsfeed.New()
		fmt.Println(feed)
		feed.Add(newsfeed.Item{"Hello", "How ya' doing mate?"})
	*/
}
