package taskcontroller

import (
	"net/http"
	"net/url"
	"todolist/models"

	"github.com/gin-gonic/gin"
)

func StatusController(c *gin.Context) {

	params := c.Param("id")

	// Fetch the existing category
	var category models.Task
	models.DB.Where("id = ?", params).Find(&category)

	if category.Status == "active" {
		category.Status = "deactive"
	} else {
		category.Status = "active"
	}

	// Update the item in the database
	models.DB.Save(&category)

	// Set the message
	msg := "Status Updated!!"

	c.Redirect(http.StatusSeeOther, "/category/list?message="+url.QueryEscape(msg))
}
