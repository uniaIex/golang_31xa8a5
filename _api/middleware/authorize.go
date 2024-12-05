package middleware

import (
	"app/database"
	"app/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

/* ----------- "" means the response is json, otherwise it redirects to another web page ---------- */

func Auth(redir string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//get cookie off request
		tokenString, err := c.Cookie("Authorization")

		if err != nil {
			if redir != "" {
				c.Redirect(http.StatusFound, redir)
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "No auth token",
				})
			}

			c.Abort()
			return
			//c.AbortWithStatus(http.StatusUnauthorized)
		}

		//decode, validate

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(os.Getenv("jwt_secret")), nil
		})

		if err != nil {
			if redir != "" {
				c.Redirect(http.StatusFound, redir)
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "failed to parse jwt ticket",
				})
			}

			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {

			if redir != "" {
				c.Redirect(http.StatusFound, redir)
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "tocken claims are off",
				})
			}

			c.Abort()
			return
		}

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			if redir != "" {
				c.Redirect(http.StatusFound, redir)
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "ticket expired",
				})
			}

			c.Abort()
			return
		}

		//find the user with token sub
		var user models.User
		database.DB.First(&user, claims["sub"])
		//user := controllers.UserFromId(claims["sub"].(uint))

		if user.ID == 0 {
			if redir != "" {
				c.Redirect(http.StatusFound, redir)
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "could not find user",
				})
			}

			c.Abort()
			return
		}

		// attach to req
		c.Set("user", user)

		//continue
		c.Next()

	}
}
