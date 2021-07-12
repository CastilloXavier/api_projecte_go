package mooc

import (
	event2 "api_project/API/kit/event"
)

const CourseCreatedEventType event2.Type = "events.course.created"

type CourseCreatedEvent struct {
	event2.BaseEvent
	id       string
	name     string
	duration string
}

func NewCourseCreatedEvent(id, name, duration string) CourseCreatedEvent {
	return CourseCreatedEvent{
		id:       id,
		name:     name,
		duration: duration,

		BaseEvent: event2.NewBaseEvent(id),
	}
}

func (e CourseCreatedEvent) Type() event2.Type {
	return CourseCreatedEventType
}

func (e CourseCreatedEvent) CourseID() string {
	return e.id
}

func (e CourseCreatedEvent) CourseName() string {
	return e.name
}

func (e CourseCreatedEvent) CourseDuration() string {
	return e.duration
}
