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
	Name	 string	`json:"id" binding:"required"`
	Duration string	`json:"id" binding:"required"`
}

func CreateHandler(courseRepository mooc.CourseRepository) gin.HandlerFunc{
	return func(ctx *gin.Context) {
		var req createRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		course := mooc.NewCourse(req.ID, req.Name, req.Duration)
		if err := courseRepository.Save(ctx, course); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		//course := mooc.NewCourse(req.ID, req.Name, req.Duration)

		//mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
		//db, err := sql.Open("mysql", mysqlURI)
		//if err != nil {
		//	ctx.JSON(http.StatusInternalServerError, err.Error())
		//	return
		//}

		//courseRepository := mysql.NewCourseRepositroy(db)

		//if err := courseRepository.Save(ctx, course); err != nil {
		//	ctx.JSON(http.StatusInternalServerError, err.Error())
		//	return
		//}

		ctx.Status(http.StatusCreated)
	}
}