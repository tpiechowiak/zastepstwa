package teacher

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTeachers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "teachers"})
}
