package courses

import (
	mooc "api_project/internal"
	"api_project/internal/creating"
	"api_project/kit/command"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)


type createRequest struct {
	ID		 string	`json:"id" binding:"required"`
	Name	 string	`json:"name" binding:"required"`
	Duration string	`json:"duration" binding:"required"`
}

func CreateHandler(commandBus command.Bus) gin.HandlerFunc{
	return func(ctx *gin.Context) {
		var req createRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err := commandBus.Dispatch(ctx, creating.NewCourseCommand(
			req.ID,
			req.Name,
			req.Duration,
		))
		//err := creatingCourseService.CreateCourse(ctx, req.ID, req.Name, req.Duration)

		if err != nil {
			switch {
			case errors.Is(err, mooc.ErrInvalidCourseID),
				errors.Is(err, mooc.ErrEmptyCourseName), errors.Is(err, mooc.ErrInvalidCourseID):
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