package postfile

import (
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	Id         int                   `uri:"id"`
	Name       string                `form:"name"`
	Email      string                `form:"email"`
	DockerFile *multipart.FileHeader `form:"file"`
}

func UploadFile() {
	r := gin.Default()
	r.POST("/upload", func(c *gin.Context) {
		var userObj user
		if err := c.ShouldBind(&userObj); err != nil {
			c.String(http.StatusBadRequest, "bad request 1")
			return
		}

		if err := c.ShouldBindUri(&userObj); err != nil {
			c.String(http.StatusBadRequest, "bad request 2")
			return
		}

		err := c.SaveUploadedFile(userObj.DockerFile, userObj.DockerFile.Filename)
		if err != nil {
			c.String(http.StatusInternalServerError, "unknown error")
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   userObj,
		})
	})
	// r.Static("fileDocker", "./fileDocker")

	r.Run("localhost:8080")
}
