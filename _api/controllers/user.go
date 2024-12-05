package controllers

import (
	"app/database"
	"app/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

/* -------- main -------- */

func Signup(c *gin.Context) {
	var body struct {
		Username        string
		Email           string
		Password        string
		ConfirmPassword string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body request",
		})

		return
	}

	if body.Password != body.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "passwords do not match",
		})

		return
	}

	//hash passwd

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to hash password",
		})

		return
	}

	//create user

	user := models.User{Username: body.Username, Email: body.Email, Password: string(hash)}
	result := database.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
	}

	//respond

	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {

	//get the email and password off req body

	var body struct {
		Username string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body request",
		})

		return
	}

	//look up on request user
	var user models.User
	database.DB.First(&user, "username = ?", body.Username)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid username",
		})

		return
	}
	//compare sent in pass woth saved user pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid password",
		})

		return
	}
	//generate a jwt tocken
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 4).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("jwt_secret")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})

		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*4, "", "", false, true) //set to true

	c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context) {

	var body struct {
		Id    uint
		Email string
	}

	user, _ := c.Get("user")

	body.Id = user.(models.User).ID
	body.Email = user.(models.User).Email

	c.JSON(http.StatusOK, gin.H{
		"messeage": body,
	})
}

func Logout(c *gin.Context) {
	c.SetCookie("Authorization", "deleted", 0, "", "", false, true) //set to true
	c.JSON(http.StatusOK, gin.H{
		"info": "Successfully logged out",
	})
}

/*
func LogoutWeb(redir string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie("Authorization", "deleted", 0, "", "", false, true)
		if redir != "" {
			c.Redirect(http.StatusFound, redir)
			c.Abort()
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"info": "Successfully logged out",
			})
			return
		}

	}
}
*/
/* ---------------------- */
