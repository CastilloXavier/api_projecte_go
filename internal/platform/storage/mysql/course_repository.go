package mysql

import (

	"context"
	"database/sql"
	"fmt"
	"time"

	mooc "api_project/internal"
	"github.com/huandu/go-sqlbuilder"
)

type CourseRepository struct {
	db *sql.DB
	dbTimeout time.Duration
}

func NewCourseRepository(db *sql.DB, dbTimeout time.Duration) *CourseRepository {
	return &CourseRepository{
		db: db,
		dbTimeout: dbTimeout,
	}
}

func (r *CourseRepository) Save(ctx context.Context, course mooc.Course) error {
	courseSQLStruct := sqlbuilder.NewStruct(new(sqlCourse))
	query, args := courseSQLStruct.InsertInto(sqlCourseTable, sqlCourse{
		ID: course.ID().String(),
		Name:  course.Name().String(),
		Duration: course.Duration().String(),
	}).Build()

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error tying to persit course on database: %v", err)
	}
	return nil
}