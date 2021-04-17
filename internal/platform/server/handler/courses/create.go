package courses

import (
	mooc "api_project/internal"
	"github.com/gin-gonic/gin"
	"net/http"
)


/*const (
	dbUser = "root"
	dbPass = "gs3458"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "codely"
)*/

type createRequest struct {
	ID		 string	`json:"id" binding:"required"`
	Name	 string	`json:"name" binding:"required"`
	Duration string	`json:"duration" binding:"required"`
}

func CreateHandler(courseRepository mooc.CourseRepository) gin.HandlerFunc{
	return func(ctx *gin.Context) {
		var req createRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		course, err := mooc.NewCourse(req.ID, req.Name, req.Duration)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		if err := courseRepository.Save(ctx, course); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Status(http.StatusCreated)
	}
}