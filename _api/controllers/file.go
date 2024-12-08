package controllers

import (
	"app/database"
	"app/models"
	"math/rand"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func fileNameWithoutExtSliceNotation(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}

func generateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func FileIndex(c *gin.Context) {
	// Get all records

	reqId, _ := c.Get("userId")

	if reqId != nil {
		var files []models.File
		database.DB.Find(&files, "user_id = ?", reqId.(uint))
		c.JSON(http.StatusOK, gin.H{
			"files": files,
		})
	}

	c.JSON(http.StatusForbidden, gin.H{
		"info": "Not Authorized",
	})

}

func FileAdd(c *gin.Context) {

	// Source
	file, err := c.FormFile("files_test")

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get the file from form",
		})

	}

	reqId, _ := c.Get("userId")

	if reqId != nil {
		extension := strings.LastIndex(file.Filename, ".")
		println("got the user id")
		if extension == -1 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "File does not have an extesion.",
			})
		}

		ext := file.Filename[extension:] //obtain the extension in ext variable

		filename := generateRandomString(25)
		if err := c.SaveUploadedFile(file, "./views/assets/uploads/"+filename+ext); err != nil {
			c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
			return
		}

		file := models.File{UserId: reqId.(uint), Path: filename, Type: ext, Name: fileNameWithoutExtSliceNotation(file.Filename)}
		result := database.DB.Create(&file)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to create file",
			})
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"info": "cannot get user id from the context so none",
		})
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"OK": "File uploaded succesfully",
	})

}

func FileUpdate(c *gin.Context) {
	/*
		type URI struct {
			Id uint `json:"id" uri:"id"`
		}

		var uri URI
		// binding to URI
		if err := c.BindUri(&uri); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		var file models.File

		result := database.DB.First(&file, uri.Id)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "no file with the mentioned id",
			})

			return
		}

		var body struct {
			Name        string
			Description string
			Image       string
		}

		if c.Bind(&body) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "failed to read body request",
			})

			return
		}

		database.DB.Model(&models.File{ID: uri.Id}).Updates(models.File{
			Path:  body.Description,
			Image: body.Image,
		})
	*/
}

func FileDelete(c *gin.Context) {
	/*
		type URI struct {
			Id uint `json:"id" uri:"id"`
		}

		var uri URI
		// binding to URI
		if err := c.BindUri(&uri); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		var file models.File

		result := database.DB.First(&file, uri.Id)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "no file with the mentioned id",
			})

			return
		}

		database.DB.Delete(&models.File{}, uri.Id)*/
}
