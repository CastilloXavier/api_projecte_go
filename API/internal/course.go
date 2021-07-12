package mooc

import (
	event2 "api_project/API/kit/event"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

var ErrInvalidCourseID = errors.New("invalid Course ID")

type CourseID struct {
	value string
}

func NewCourseID(value string) (CourseID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return CourseID{}, fmt.Errorf("%w: %s", ErrInvalidCourseID, value)
	}

	return CourseID{
		value : v.String(),
	}, nil
}

func (id CourseID) String() string {
	return id.value
}

var ErrEmptyCourseName = errors.New("the field Course Name can not be empty")

type CourseName struct {
	value string
}

func NewCourseName(value string) (CourseName, error) {
	if value == "" {
		return CourseName{}, ErrEmptyCourseName
	}

	return CourseName{
		value: value,
	}, nil
}

func (name CourseName) String() string {
	return name.value
}

var ErrEmptyDuration = errors.New("The field Duration can not be empty")

type CourseDuration struct {
	value string
}

func NewCourseDuration(value string) (CourseDuration, error) {
	if value == "" {
		return CourseDuration{}, ErrEmptyDuration
	}

	return CourseDuration{
		value: value,
	}, nil
}

func (duration CourseDuration) String() string {
	return duration.value
}

// Course is the data structure that represents a course.
type Course struct {
	id       CourseID
	name     CourseName
	duration CourseDuration

	events []event2.Event
}

// CourseRepository defines the expected behaviour from a course storage.
type CourseRepository interface {
	Save(ctx context.Context, course Course) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=CourseRepository

func NewCourse(id, name, duration string) (Course, error) {
	idVO, err := NewCourseID(id)
	if err != nil {
		return Course{}, err
	}

	nameVO, err := NewCourseName(name)
	if err != nil {
		return Course{}, err
	}

	durationVO, err := NewCourseDuration(duration)
	if err != nil {
		return Course{}, err
	}

	return  Course{
		id: idVO,
		name: nameVO,
		duration: durationVO,
	}, nil
}

// ID returns the course unique identifier.
func (c Course) ID() CourseID {
	return c.id
}

// Name returns the course name.
func (c Course) Name() CourseName {
	return c.name
}

// Duration returns the course duration.
func (c Course) Duration() CourseDuration {
	return c.duration
}

// Record records a new domain event.
func (c *Course) Record(evt event2.Event) {
	c.events = append(c.events, evt)
}

// PullEvents returns all the recorded domain events.
func (c Course) PullEvents() []event2.Event {
	evt := c.events
	c.events = []event2.Event{}

	return evt
}
