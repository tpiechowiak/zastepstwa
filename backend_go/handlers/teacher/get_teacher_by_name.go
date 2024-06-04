package teacher

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTeacherByName(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "GetTeacherByName" + id + " Called"})
}
