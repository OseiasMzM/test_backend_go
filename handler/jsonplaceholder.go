package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test_backend_go/service"
)

func GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := service.FetchUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}

	c.JSON(http.StatusOK, user)
}
