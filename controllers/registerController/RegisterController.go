package registercontroller

import (
	"net/http"
	"todolist/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const PasswordMinLength = 6

// NoError represents the absence of an error.
var NoError = ErrorMessage{Error: ""}

type ErrorMessage struct {
	Error string
}

func RegisterController(c *gin.Context) {

	session := sessions.Default(c)
	user := session.Get("user")
	if user != nil {
		// User not authenticated, redirect to login
		c.Redirect(http.StatusSeeOther, "/")
		c.Abort() // Stop the request chain
		return
	}

	data := gin.H{
		"title": "Registration Page",
	}
	c.HTML(http.StatusOK, "register.html", data)
}

func RegisterStoreController(c *gin.Context) {

	name := c.PostForm("name")
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	password := c.PostForm("password")

	validationErrors := validateInputs(name, email, phone, password)
	if validationErrors != NoError {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"error": validationErrors.Error,
		})
		return
	}

	newUser := models.User{
		Name:     name,
		Phone:    phone,
		Email:    email,
		Password: hashPassword(password),
		Status:   "active",
	}

	result := models.DB.Create(&newUser)
	if result.Error != nil {
		data := ErrorMessage{"Somthing went wrong"}
		c.HTML(http.StatusOK, "register.html", data)
	} else {
		data := gin.H{"message": "Registration successfully. Pls login. Thank you."}
		c.HTML(http.StatusOK, "login.html", data)
	}
}

func validateInputs(name, email, phone, password string) ErrorMessage {

	if name == "" {
		return ErrorMessage{"Name field is required!!"}
	}

	if phone == "" {
		return ErrorMessage{"Phone field is required!!"}
	}

	if email == "" {
		return ErrorMessage{"Email field is required!!"}
	}

	if checkDuplicateEmail(email) {
		return ErrorMessage{"Email Already Exist!!"}
	}

	if checkDuplicatePhone(phone) {
		return ErrorMessage{"Phone Number Already Exist!!"}
	}

	if password == "" {
		return ErrorMessage{"Password field is required!!"}
	}

	if len(password) < PasswordMinLength {
		return ErrorMessage{"Password should be more than 5 characters!!"}
	}

	return NoError
}

func checkDuplicateEmail(email string) bool {
	var checkEmail int64
	models.DB.Model(&models.User{}).Where("email = ?", email).Count(&checkEmail)
	return checkEmail > 0
}

func checkDuplicatePhone(phone string) bool {
	var checkNumber int64
	models.DB.Model(&models.User{}).Where("phone = ?", phone).Count(&checkNumber)
	return checkNumber > 0
}

func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic("Password failed to hash!!")
	}
	return string(hashedPassword)
}

func RegisterUpdateController(c *gin.Context) {

	session := sessions.Default(c)
	email := session.Get("user")
	if email == nil {
		// User not authenticated, redirect to login
		c.Redirect(http.StatusSeeOther, "/login")
		c.Abort() // Stop the request chain
		return
	}

	var user []models.User
	models.DB.Model(&user).Where("email = ?", email).Find(&user)

}
