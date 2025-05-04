package controllers

// verificación básica de que el servidor backend está funcionando correctamente.
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
