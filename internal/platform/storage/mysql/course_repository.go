package mysql

import (
	mooc "api_project/internal"
	"context"
	"database/sql"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
)

type CourseRepository struct {
	db *sql.DB
}

func NewCourseRepositroy(db *sql.DB) *CourseRepository {
	return &CourseRepository{
		db: db,
	}
}

func (r *CourseRepository) Save(ctx context.Context, course mooc.Course) error {
	courseSQLStruct := sqlbuilder.NewStruct(new(sqlCourse))
	query, args := courseSQLStruct.InsertInto(sqlCourseTable, sqlCourse{
		ID: course.ID().String(),
		Name:  course.Name().String(),
		Duration:  course.Duration().String(),
	}).Build()

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error tying to persit course on database: %v", err)
	}
	return nil
}