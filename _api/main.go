package main

import (
	//"app/services"
	"app/controllers"
	"app/database"
	"app/environment"
	"app/middleware"
	"html/template"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	environment.Load()
	database.Connect()
	database.Sync()
	database.Admin()
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Static("/css", "./views/css")
	r.Static("/js", "./views/js")
	r.Static("/img", "./views/assets/img")
	r.Static("/uploads", "./views/assets/uploads")

	//r.MaxMultipartMemory = 8 << 20 // 8 MiB

	api := r.Group("/api")
	{

	}

	admin := api.Group("/admin")
	{
		admin.GET("/validate", middleware.Auth(""), controllers.Validate)
	}

	projects := api.Group("/project")
	{
		projects.GET("", controllers.ProjectIndex)
		projects.POST("/add", middleware.Auth(""), controllers.ProjectAdd)
		projects.PUT("/update/:id", middleware.Auth(""), controllers.ProjectUpdate)
		projects.DELETE("/delete/:id", middleware.Auth(""), controllers.ProjectDelete)
	}

	file := api.Group("/file")
	{
		file.GET("", middleware.Auth(""), controllers.FileIndex)
		file.POST("/add", middleware.Auth(""), controllers.FileAdd)
		file.PUT("/update/:id", middleware.Auth(""), controllers.FileUpdate)
		file.DELETE("/delete/:id", middleware.Auth(""), controllers.FileDelete)
	}

	users := api.Group("/user")
	{
		users.POST("/login", controllers.Login)
		users.POST("/create_account", middleware.Auth(""), controllers.Signup)
		users.POST("/logout", middleware.Auth(""), controllers.Logout)
	}

	renderer := r.Group("/")
	{
		renderer.GET("/", func(c *gin.Context) {
			t := template.Must(template.ParseFiles("views/index.html"))
			t.Execute(c.Writer, nil)

		})

		renderer.GET("/user", middleware.Auth(""), func(c *gin.Context) {
			reqId, _ := c.Get("userId")
			reqUsername, _ := c.Get("userUsername")
			reqRole, _ := c.Get("userRole")

			if reqId != nil && reqUsername != nil && reqRole != nil {
				c.JSON(http.StatusOK, gin.H{
					"id":       reqId,
					"username": reqUsername,
					"role":     reqRole,
				})
			}

			c.JSON(http.StatusOK, gin.H{
				"data": "nope",
			})

		})

		renderer.GET("/pdf", func(c *gin.Context) {
			t := template.Must(template.ParseFiles("views/pdf.html"))
			t.Execute(c.Writer, nil)

		})

		renderer.GET("/auth", func(c *gin.Context) {
			t := template.Must(template.ParseFiles("views/auth.html"))
			t.Execute(c.Writer, nil)

		})

		renderer.GET("/admin", middleware.Auth("/auth"), func(c *gin.Context) {
			t := template.Must(template.ParseFiles("views/admin.html"))
			t.Execute(c.Writer, nil)

		})

		renderer.GET("/admin/project/add", middleware.Auth("/auth"), func(c *gin.Context) {
			t := template.Must(template.ParseFiles("views/addproject.html"))
			t.Execute(c.Writer, nil)

		})

		renderer.GET("/admin/file/add", middleware.Auth("/auth"), func(c *gin.Context) {
			t := template.Must(template.ParseFiles("views/addfile.html"))
			t.Execute(c.Writer, nil)

		})

		renderer.GET("/admin/technology/add", middleware.Auth("/auth"), func(c *gin.Context) {
			t := template.Must(template.ParseFiles("views/addtechnology.html"))
			t.Execute(c.Writer, nil)

		})

		renderer.GET("/admin/language/add", middleware.Auth("/auth"), func(c *gin.Context) {
			t := template.Must(template.ParseFiles("views/addlanguage.html"))
			t.Execute(c.Writer, nil)

		})

	}

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(os.Getenv("port"))
}
