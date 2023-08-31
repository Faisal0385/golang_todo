package logincontroller

import (
	"net/http"
	"todolist/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginController(c *gin.Context) {

	session := sessions.Default(c)
	user := session.Get("user")
	if user != nil {
		// User not authenticated, redirect to login
		c.Redirect(http.StatusSeeOther, "/")
		c.Abort() // Stop the request chain
		return
	}

	data := gin.H{
		"title": "Login Page",
	}
	c.HTML(http.StatusOK, "login.html", data)
}

type ValidationError struct {
	Fields map[string]string
}

func validateEmptyField(value, fieldName string) (bool, string) {
	if value == "" {
		return false, fieldName + " field is required!!"
	}
	return true, ""
}

func LoginStoreController(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	validationErr := &ValidationError{}
	// In your LoginStoreController function
	isEmailValid, emailError := validateEmptyField(email, "Email")
	isPasswordValid, passwordError := validateEmptyField(password, "Password")

	if !isEmailValid || !isPasswordValid {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": emailError + " " + passwordError,
		})
		return
	}

	if len(validationErr.Fields) > 0 {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": validationErr,
		})
		return
	}

	// Authentication logic ...
	var user models.User
	models.DB.Where("email = ?", email).First(&user)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": "Invalid credentials",
		})
		return
	}

	// Set flash message
	session := sessions.Default(c)
	session.AddFlash("Login successful!", "message")
	session.Set("user", user.Email)
	session.Save()

	c.Redirect(http.StatusSeeOther, "/")
}
