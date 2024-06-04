package teacher

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllSubs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetAllSubs Called"})
}
