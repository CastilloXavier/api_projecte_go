package creating

import (
	command2 "github.com/api_project_go/API/kit/command"
	"context"
	"errors"
)

const CourseCommandType command2.Type = "command.creating.course"

// CourseCommand is the command dispatched to create a new course.
type CourseCommand struct {
	id 		 string
	name 	 string
	duration string
}

// NewCourseCommand creates a new CourseCommand
func NewCourseCommand(id, name, duration string) CourseCommand {
	return CourseCommand{
		id: id,
		name: name,
		duration: duration,
	}
}

func (c CourseCommand) Type() command2.Type {
	return CourseCommandType
}

// CourseCommandHandler is the command handler
// responsible for creating courses.
type CourseCommandHandler struct {
	service CourseService
}

// NewCourseCommandHandler initializes a new CourseCommandHandler.
func NewCourseCommandHandler(service CourseService) CourseCommandHandler {
	return CourseCommandHandler{
		service: service,
	}
}

// Handle implements the command.Handler interface.
func (h CourseCommandHandler) Handle(ctx context.Context, cmd command2.Command) error {
	createCourseCmd, ok := cmd.(CourseCommand)
	if !ok {
		return errors.New("unexpected command")
	}

	return h.service.CreateCourse(
		ctx,
		createCourseCmd.id,
		createCourseCmd.name,
		createCourseCmd.duration,
	)
}

