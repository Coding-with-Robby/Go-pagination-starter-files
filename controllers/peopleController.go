package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robbyklein/pages/initializers"
	"github.com/robbyklein/pages/models"
)

func PeopleIndexGET(c *gin.Context) {
	// Get the people
	var people []models.Person
	initializers.DB.Find(&people)

	// Render the page
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"people": people,
	})
}
