package taskcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddTaskController(c *gin.Context) {
	data := gin.H{
		"title": "Task Add Page",
	}
	c.HTML(http.StatusOK, "add_task.html", data)
}
