package main

import (
	"image"
	"image/jpeg"  // Encoding to JPEG
	_ "image/png" // Decode PNGs
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize" //Resizing library
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
	uploadedFile, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file received"})
		return
	}

	// Opens uploaded file to allow for resize and conversion
	file, err := uploadedFile.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open file"})
		return
	}
	defer file.Close()

	// Decode image into image.Image object
	img, _, err := image.Decode(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to decode image"})
		return
	}

	// Resize to max width of 1024
	resizedImg := resize.Resize(1024, 0, img, resize.Lanczos3)

	// Define the directory and ensures it exists
	uploadDir := "./uploads"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create upload directory"})
		return
	}

	// Create new file to save processed image
	processedFilename := "processed_" + strings.TrimSuffix(uploadedFile.Filename, filepath.Ext(uploadedFile.Filename)) + ".jpeg"
	destination := filepath.Join("./uploads/", processedFilename)

	outFile, err := os.Create(destination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create output file"})
		return
	}
	defer outFile.Close()

	// Encode resized image as a JPEG with 75% quality
	jpeg.Encode(outFile, resizedImg, &jpeg.Options{Quality: 75})

	c.JSON(http.StatusOK, gin.H{
		"message":  "File uploaded successfully",
		"filename": processedFilename,
	})
}
