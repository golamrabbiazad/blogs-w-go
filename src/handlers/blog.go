package handlers

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"

	. "github.com/golamrabbiazad/blogs-w-go/db"
)

func GetBlogs(c *gin.Context) {
	var blogs []Blog

	gorm := Gorm()
	gorm.Find(&blogs)

	c.JSON(http.StatusOK, gin.H{"data": blogs})
}

type CreateBlogInput struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required"`
	Upvotes int32  `json:"up_votes" binding:"required"`
}

func PostBlog(c *gin.Context) {
	var input CreateBlogInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// create a blog
	blog := Blog{
		ID: rand.Intn(200),
		Author: Author{
			Name:  input.Name,
			Email: input.Email,
		},
		Upvotes: input.Upvotes,
	}
	gorm := Gorm()
	gorm.Create(&blog)

	c.JSON(http.StatusOK, gin.H{
		"data": blog,
	})
}
