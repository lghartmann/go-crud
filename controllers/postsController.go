package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lghartmann/go-crud/initializers"
	"github.com/lghartmann/go-crud/models"
)

type PostBody struct {
	Body  string
	Title string
}

func PostsCreate(c *gin.Context) {
	body := PostBody{}
	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{"post": post})
}

func GetPosts(c *gin.Context) {
	posts := []models.Post{}

	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{"posts": posts})
}

func GetPostById(c *gin.Context) {
	id := c.Param("id")

	post := models.Post{}

	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{"post": post})
}

func PostsUpdate(c *gin.Context) {
	id := c.Param("id")

	body := PostBody{}
	c.Bind(&body)

	post := models.Post{}
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	c.JSON(200, gin.H{"post": post})
}

func PostsDelete(c *gin.Context) {
	id := c.Param("id")

	post := models.Post{}

	initializers.DB.Delete(&post, id)

	c.JSON(200, gin.H{"message": "Successfully deleted"})
}
