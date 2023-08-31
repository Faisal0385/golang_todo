package taskcontroller

import (
	"net/http"
	"todolist/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ListTaskController(c *gin.Context) {

	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		// User not authenticated, redirect to login
		c.Redirect(http.StatusSeeOther, "/login")
		c.Abort() // Stop the request chain
		return
	}

	var tasks []models.Task
	models.DB.Model(&tasks).Where("status = ?", "new").Find(&tasks)

	var newCount int64
	var progressCount int64
	var completedCount int64
	var cancelledCount int64
	models.DB.Model(&tasks).Where("status = ?", "new").Count(&newCount)
	models.DB.Model(&tasks).Where("status = ?", "progress").Count(&progressCount)
	models.DB.Model(&tasks).Where("status = ?", "completed").Count(&completedCount)
	models.DB.Model(&tasks).Where("status = ?", "cancelled").Count(&cancelledCount)

	// Retrieve the message from the query parameter
	message := c.DefaultQuery("message", "")

	data := gin.H{
		"title":          "Task List Page",
		"values":         tasks,
		"message":        message,
		"newCount":       newCount,
		"progressCount":  progressCount,
		"completedCount": completedCount,
		"cancelledCount": cancelledCount,
	}
	c.HTML(http.StatusOK, "index.html", data)
}
