package helpers

import "github.com/gin-gonic/gin"

func ElorResponse(c *gin.Context, message string) {
	c.JSON(400, gin.H{
		"status":  "Elor",
		"message": message,
	})
}

func TokenSuksesResponse(c *gin.Context, message string, token string) {
	c.JSON(201, gin.H{
		"status":  "Sukses",
		"message": message,
		"token":   token,
	})
}

func SuksesWithData(c *gin.Context, message string, data interface{}) {
	c.JSON(201, gin.H{
		"status":  "Sukses",
		"message": message,
		"data":    data,
	})
}
