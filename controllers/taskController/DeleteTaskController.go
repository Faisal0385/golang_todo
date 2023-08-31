package taskcontroller

import (
	"net/http"
	"net/url"
	"todolist/models"

	"github.com/gin-gonic/gin"
)

func DeleteTaskController(c *gin.Context) {

	params := c.Param("id")

	// Fetch the existing task
	var task models.Task
	models.DB.Where("id = ?", params).Find(&task)

	// Delete models
	models.DB.Delete(&task, params)
	// Set the message
	msg := "Item Deleted Successfully!"

	c.Redirect(http.StatusSeeOther, "/?message="+url.QueryEscape(msg))
}
