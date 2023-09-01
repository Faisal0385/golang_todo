package usercontroller

import (
	"net/http"
	"todolist/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ProfileController(c *gin.Context) {

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
	data := gin.H{
		"title":  "Task List Page",
		"values": user[0],
	}
	c.HTML(http.StatusOK, "profile.html", data)
}
