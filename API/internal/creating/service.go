package creating

import (
	mooc2 "api_project_go/API/internal"
	event2 "api_project_go/API/kit/event"
	"context"
)

type CourseService struct {
	courseRepository mooc2.CourseRepository
	eventBus         event2.Bus
}

func NewCourseService(courseRepository mooc2.CourseRepository, eventBus event2.Bus) CourseService {
	return CourseService{
		courseRepository: courseRepository,
		eventBus: 		  eventBus,
	}
}

func (s CourseService) CreateCourse(ctx context.Context, id, name, duration string) error {
	course, err := mooc2.NewCourse(id, name, duration)
	if err != nil {
		return err
	}

	if err := s.courseRepository.Save(ctx, course); err != nil {
		return err
	}

	return s.eventBus.Publish(ctx, course.PullEvents())
}
