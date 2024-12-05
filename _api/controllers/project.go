package controllers

import (
	"app/database"
	"app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProjectIndex(c *gin.Context) {
	// Get all records
	var projects []models.Project
	database.DB.Find(&projects)
	// SELECT * FROM users;

	c.JSON(http.StatusOK, gin.H{
		"data": projects,
	})
}

func ProjectAdd(c *gin.Context) {

	var body struct {
		Name        string
		Description string
		Github      string
		Deployment  string
		Technology  string
		Platform    string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body request",
		})

		return
	}

	project := models.Project{Name: body.Name, Description: body.Description, Github: body.Github, Deployment: body.Deployment, Technology: body.Technology, Platform: body.Platform}
	result := database.DB.Create(&project)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create project",
		})
	}

	//respond

	c.JSON(http.StatusOK, gin.H{
		"info": "project created",
	})
}

func ProjectUpdate(c *gin.Context) {

	type URI struct {
		Id uint `json:"id" uri:"id"`
	}

	var uri URI
	// binding to URI
	if err := c.BindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var project models.Project

	result := database.DB.First(&project, uri.Id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no project with the mentioned id",
		})

		return
	}

	var body struct {
		Name        string
		Description string
		Github      string
		Deployment  string
		Technology  string
		Platform    string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body request",
		})

		return
	}

	database.DB.Model(&models.Project{ID: uri.Id}).Updates(models.Project{
		Name:        body.Name,
		Description: body.Description,
		Github:      body.Github,
		Deployment:  body.Deployment,
		Technology:  body.Technology,
		Platform:    body.Platform,
	})

}

func ProjectDelete(c *gin.Context) {

	type URI struct {
		Id uint `json:"id" uri:"id"`
	}

	var uri URI
	// binding to URI
	if err := c.BindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var project models.Project

	result := database.DB.First(&project, uri.Id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no project with the mentioned id",
		})

		return
	}

	database.DB.Delete(&models.Project{}, uri.Id)
}
