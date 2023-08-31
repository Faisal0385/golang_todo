package taskcontroller

import (
	"net/http"
	"todolist/models"

	"github.com/gin-gonic/gin"
)

func UpdateTaskController(c *gin.Context) {

	id := c.Param("id")
	// Fetch the existing task
	var task models.Task
	models.DB.Where("id = ?", id).Find(&task)

	newtask := c.PostForm("task")
	description := c.PostForm("description")

	progress := c.PostForm("taskStatus")

	if newtask == "" {
		data := gin.H{"error": "Task field is required!!", "values": task}
		c.HTML(http.StatusBadRequest, "edit_task.html", data)
		return
	}

	// Update the item in the database
	task.Task = newtask
	task.Description = description
	task.Status = progress
	models.DB.Save(&task)

	data := gin.H{
		"message": "Data Updated!!",
		"values":  task,
	}

	c.HTML(http.StatusOK, "edit_task.html", data)
}
