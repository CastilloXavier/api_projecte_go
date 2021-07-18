package courses

import (
	mooc2 "github.com/CastilloXavier/api_project_go/API/internal"
	creating2 "github.com/CastilloXavier/api_project_go/API/internal/creating"
	command2 "github.com/CastilloXavier/api_project_go/API/kit/command"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)


type createRequest struct {
	ID		 string	`json:"id" binding:"required"`
	Name	 string	`json:"name" binding:"required"`
	Duration string	`json:"duration" binding:"required"`
}

func CreateHandler(commandBus command2.Bus) gin.HandlerFunc{
	return func(ctx *gin.Context) {
		var req createRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err := commandBus.Dispatch(ctx, creating2.NewCourseCommand(
			req.ID,
			req.Name,
			req.Duration,
		))
		//err := creatingCourseService.CreateCourse(ctx, req.ID, req.Name, req.Duration)

		if err != nil {
			switch {
			case errors.Is(err, mooc2.ErrInvalidCourseID),
				errors.Is(err, mooc2.ErrEmptyCourseName), errors.Is(err, mooc2.ErrInvalidCourseID):
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			default:
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		ctx.Status(http.StatusCreated)
	}
}