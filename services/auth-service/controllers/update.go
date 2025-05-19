package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vidinine-ecommerce/auth-service/models"
	"github.com/vidinine-ecommerce/auth-service/services"
)

func UpdateHandler(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		controllerLogger.Error("User not authenticated")
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "unauthenticated",
			"message": "User not authenticated",
		})
		return
	}

	var input models.UpdateUser
	if err := c.ShouldBindJSON(&input); err != nil {
		controllerLogger.Errorf("Invalid user data: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid_data",
			"message": "Invalid user data: " + err.Error(),
		})
		return
	}

	updatedUser, err := services.UpdateUser(userID.(uint), input)
	if err != nil {
		controllerLogger.Errorf("Failed to update user: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "update_failed",
			"message": err.Error(),
		})
		return
	}

	controllerLogger.Successf("User updated successfully (ID: %d)", updatedUser.ID)
	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"data": gin.H{
			"id":    updatedUser.ID,
			"name":  updatedUser.Name,
			"email": updatedUser.Email,
		},
	})
}
