package controllers

import (
	"net/http"
	"task5-pbi-btpns-ArifBudiantoPratomo/app"
	"task5-pbi-btpns-ArifBudiantoPratomo/database"
	"task5-pbi-btpns-ArifBudiantoPratomo/helpers"
	"task5-pbi-btpns-ArifBudiantoPratomo/models"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	userInput := app.UserRegister{
		Username: c.PostForm("username"),
		Email:    c.PostForm("email"),
		Password: c.PostForm("password"),
	}

	if _, err := govalidator.ValidateStruct(userInput); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	hashedPassword, _ := helpers.HashPassword(userInput.Password)

	user := models.User{
		Username: userInput.Username,
		Email:    userInput.Email,
		Password: hashedPassword,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "New user created"})
}

func Login(c *gin.Context) {
	userInput := app.UserLogin{
		Email:    c.PostForm("email"),
		Password: c.PostForm("password"),
	}

	if _, err := govalidator.ValidateStruct(userInput); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	var user models.User
	if err := database.DB.First(&user, "email = ?", userInput.Email).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Invalid email"})
		return
	}

	if err := helpers.CheckPasswordHash(userInput.Password, user.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid password"})
		return
	}

	accessToken, _ := helpers.GenerateToken(user.ID)

	c.JSON(http.StatusCreated, gin.H{"message": "Login Success", "userID": user.ID, "token": accessToken})
}

func GetUsers(c *gin.Context) {
	var users []models.User
	database.DB.Preload("Photos").Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	
	var user models.User
	if err := database.DB.Preload("Photos").First(&user, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	userInput := app.UserUpdate{
		Username: c.PostForm("username"),
		Email:    c.PostForm("email"),
		Password: c.PostForm("password"),
	}

	if _, err := govalidator.ValidateStruct(userInput); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	hashedPassword, _ := helpers.HashPassword(userInput.Password)

	user := models.User{
		Username: userInput.Username,
		Email:    userInput.Email,
		Password: hashedPassword,
	}

	if err := database.DB.Model(&user).Where("id = ?", id).Updates(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated"})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	if err := database.DB.Unscoped().Delete(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
