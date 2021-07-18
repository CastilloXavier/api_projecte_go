package creating

import (
	mooc2 "api_project_go/API/internal"
	increasing2 "api_project_go/API/internal/increasing"
	event2 "api_project_go/API/kit/event"
	"context"
	"errors"
)

type IncreaseCoursesCounterOnCourseCreated struct {
	increasingService increasing2.CourseCounterService
}

func NewIncreaseCoursesCounterOnCourseCreated(increaserService increasing2.CourseCounterService) IncreaseCoursesCounterOnCourseCreated {
	return IncreaseCoursesCounterOnCourseCreated{
		increasingService: increaserService,
	}
}

func (e IncreaseCoursesCounterOnCourseCreated) Handle(_ context.Context, evt event2.Event) error {
	courseCreatedEvt, ok := evt.(mooc2.CourseCreatedEvent)
	if !ok {
		return errors.New("unexpected event")
	}

	return e.increasingService.Increase(courseCreatedEvt.ID())
}
