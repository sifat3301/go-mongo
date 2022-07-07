package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-mongo/models"
	"go-mongo/services"
	"net/http"
	"time"
)

var validate = validator.New()

func CreateUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	defer cancel()

	// Validate request body
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//	User input validate library to validate required field
	if validationError := validate.Struct(&user); validationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": validationError.Error(),
		})
	}
	user, err := services.CreateUserService(&user, ctx)
	if err != nil {
		fmt.Println("This is working", err)
		c.JSON(http.StatusConflict, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)

}

func GetUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Param("userId")
	defer cancel()
	user, err := services.GetUserService(&userId, &ctx)
	if err != nil {
		fmt.Println("This is working", err)
		c.JSON(http.StatusConflict, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

func GetAllUsers(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	users, err := services.GetAllUserService(&ctx)
	if err != nil {
		fmt.Println("This is working", err)
		c.JSON(http.StatusConflict, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, users)
}

func EditUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Param("userId")
	defer cancel()
	var user models.User

	// Validate request body
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//	User input validate library to validate required field
	if validationError := validate.Struct(&user); validationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": validationError.Error(),
		})
	}

	updateUser, err := services.UpdateUserService(&userId, &user, &ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, updateUser)
}

func DeleteUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Param("userId")
	defer cancel()
	var user models.User

	// Validate request body
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := services.DeleteUserService(&userId, &ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "Successfully deleted",
	})
}
