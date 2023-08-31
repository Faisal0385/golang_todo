package main

import (
	"log"
	"net/http"
	logincontroller "todolist/controllers/loginController"
	registercontroller "todolist/controllers/registerController"
	taskcontroller "todolist/controllers/taskController"
	"todolist/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-gonic/gin"
)

func init() {
	models.InitDB()
}

func main() {
	var router = gin.Default()
	router.LoadHTMLGlob("templates/**/*")

	err := models.InitDB()
	if err != nil {
		panic("Failed to connect to the database")
	}
	defer models.CloseDB()

	// Initialize the session store
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.GET("/", taskcontroller.ListTaskController)
	router.GET("/progress/task", taskcontroller.TaskProgressController)
	router.GET("/completed/task", taskcontroller.TaskCompletedController)
	router.GET("/cancelled/task", taskcontroller.TaskCancelledController)

	router.GET("/add/task", taskcontroller.AddTaskController)
	router.GET("/edit/task/:id", taskcontroller.EditTaskController)
	router.GET("/delete/:id", taskcontroller.DeleteTaskController)

	router.POST("/store/task", taskcontroller.StoreTaskController)
	router.POST("/update/task/:id", taskcontroller.UpdateTaskController)

	router.GET("/login", logincontroller.LoginController)
	router.GET("/register", registercontroller.RegisterController)

	router.POST("/login", logincontroller.LoginStoreController)
	router.POST("/store/register", registercontroller.RegisterStoreController)
	router.POST("/update/register", registercontroller.RegisterUpdateController)

	router.GET("/logout", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear() // Clear session data
		session.Save()

		// Redirect to the login page or any other page
		c.Redirect(http.StatusSeeOther, "/login")
	})

	router.GET("/profile", func(c *gin.Context) {

		session := sessions.Default(c)
		userData := session.Get("user")
		if userData == nil {
			// User not authenticated, redirect to login
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort() // Stop the request chain
			return
		}

		var user []models.User

		models.DB.Model(&user).Where("email = ?", userData).First(&user)
		log.Println(user[0])
		data := gin.H{
			"title":  "Task List Page",
			"values": user[0],
		}
		c.HTML(http.StatusOK, "profile.html", data)
	})

	router.Run(":5000")
}
