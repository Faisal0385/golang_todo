package taskcontroller

import (
	"net/http"
	"todolist/models"

	"github.com/gin-gonic/gin"
)

func EditTaskController(c *gin.Context) {

	params := c.Param("id")

	// Fetch the existing task
	var task models.Task
	models.DB.Where("id = ?", params).Find(&task)

	data := gin.H{
		"title":  "Task Edit Page",
		"values": task,
	}
	c.HTML(http.StatusOK, "edit_task.html", data)
}
