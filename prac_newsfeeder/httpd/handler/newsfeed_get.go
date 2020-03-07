package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"prac_newsfeeder/platform/newsfeed"
)

func NewsfeedGet(feed newsfeed.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := feed.GetAll()
		c.JSON(http.StatusOK, result)
	}
}
