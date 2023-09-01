package main

import (
	"net/http"
	logincontroller "todolist/controllers/loginController"
	registercontroller "todolist/controllers/registerController"
	taskcontroller "todolist/controllers/taskController"
	usercontroller "todolist/controllers/userController"
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

	// Define a route that renders the template
	router.GET("/status/wise", func(c *gin.Context) {
		var tasks []models.Task
		models.DB.Model(&tasks).Find(&tasks)

		var users []models.User
		models.DB.Model(&users).Find(&users)

		var newCount int64
		var progressCount int64
		var completedCount int64
		var cancelledCount int64
		models.DB.Model(&tasks).Where("status = ?", "new").Count(&newCount)
		models.DB.Model(&tasks).Where("status = ?", "progress").Count(&progressCount)
		models.DB.Model(&tasks).Where("status = ?", "completed").Count(&completedCount)
		models.DB.Model(&tasks).Where("status = ?", "cancelled").Count(&cancelledCount)

		data := gin.H{
			"title":          "Task Status Wise Page",
			"values":         tasks,
			"users":          users,
			"newCount":       newCount,
			"progressCount":  progressCount,
			"completedCount": completedCount,
			"cancelledCount": cancelledCount,
		}
		c.HTML(http.StatusOK, "status_wise.html", data)
	})

	router.GET("/forgot/password", func(c *gin.Context) {
		data := gin.H{
			"title": "Forgot Password Page",
		}
		c.HTML(http.StatusOK, "forgot_password.html", data)
	})

	router.GET("/logout", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear() // Clear session data
		session.Save()

		// Redirect to the login page or any other page
		c.Redirect(http.StatusSeeOther, "/login")
	})

	router.GET("/profile", usercontroller.ProfileController)

	router.Run(":5000")
}
