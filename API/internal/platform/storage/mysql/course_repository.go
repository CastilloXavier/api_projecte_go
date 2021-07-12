package mysql

import (
	mooc2 "api_project/API/internal"
	"context"
	"database/sql"
	"fmt"
	"time"

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

func (r *CourseRepository) Save(ctx context.Context, course mooc2.Course) error {
	courseSQLStruct := sqlbuilder.NewStruct(new(sqlCourse))
	query, args := courseSQLStruct.InsertInto(sqlCourseTable, sqlCourse{
		ID: course.ID().String(),
		Name:  course.Name().String(),
		Duration: course.Duration().String(),
	}).Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("error tying to persit course on database: %v", err)
	}
	return nil
}