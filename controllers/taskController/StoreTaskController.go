package taskcontroller

import (
	"net/http"
	"todolist/models"

	"github.com/gin-gonic/gin"
)

func StoreTaskController(c *gin.Context) {

	task := c.PostForm("task")
	description := c.PostForm("description")

	if task == "" {
		data := gin.H{"error": "Task field is required!!"}
		c.HTML(http.StatusBadRequest, "add_task.html", data)
		return
	}

	newTask := models.Task{
		Task:        task,
		Description: description,
		Status:      "new",
	}
	result := models.DB.Create(&newTask)

	if result != nil {
		data := gin.H{"message": "Data inserted successfully!!"}
		c.HTML(http.StatusOK, "add_task.html", data)
	} else {
		data := gin.H{"error": "Somthing went wrong"}
		c.HTML(http.StatusOK, "add_task.html", data)
	}
}
