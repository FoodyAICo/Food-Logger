package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			// Message is the variable name
			"message": "ok",
		})
	})

	router.POST("/analyze", handleAnalyze)

	// Uncomment when ready for release
	// gin.SetMode(gin.ReleaseMode)
	router.Run(":8080")
}

func handleAnalyze(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file received"})
		return
	}

	destination := "./uploads/" + file.Filename

	if err := c.SaveUploadedFile(file, destination); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "File uploaded successfully",
		"filename": file.Filename,
	})
}
