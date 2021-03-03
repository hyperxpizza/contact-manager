package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hyperxpizza/contact-manager/server/filestorage"
	"github.com/minio/minio-go"
)

type File struct {
	Bucket   string `uri:"bucket" binding:"required"`
	Filepath string `uri:"filepath" binding:"required"`
}

func DownloadFile(c *gin.Context) {
	var file File
	if err := c.ShouldBindUri(&file); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	storage, err := filestorage.NewStorage()
	if err != nil {
		log.Printf("creating filestorage instance failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	stat, err := storage.Client.StatObject(file.Bucket, file.Filepath, minio.StatObjectOptions{})
	if err != nil {
		log.Printf("statObject failed: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	reader, err := storage.Client.GetObject(file.Bucket, file.Filepath, minio.GetObjectOptions{})
	if err != nil {
		log.Printf("getObject failed: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	c.DataFromReader(http.StatusOK, stat.Size, stat.ContentType, reader, nil)
	return
}
