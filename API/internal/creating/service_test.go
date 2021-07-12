package creating

import (
	mooc2 "api_project/API/internal"
	storagemocks2 "api_project/API/internal/platform/storage/storagemocks"
	eventmocks2 "api_project/API/kit/event/eventmocks"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_CourseService_CreateCourse_RepositoryError(t *testing.T) {
	courseID := "37a0f027-15e6-47cc-a5d2-64183281087e"
	courseName := "Test Course"
	courseDuration := "10 months"

	course, err := mooc2.NewCourse(courseID, courseName, courseDuration)
	require.NoError(t, err)

	courseRepositoryMock := new(storagemocks2.CourseRepository)
	courseRepositoryMock.On("Save", mock.Anything, course).Return(errors.New("something unexcepted happend"))

	eventBusMock := new(eventmocks2.Bus)

	courseService := NewCourseService(courseRepositoryMock, eventBusMock)

	err = courseService.CreateCourse(context.Background(), courseID, courseName, courseDuration)

	courseRepositoryMock.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_CourseService_CreateCourse_EventsBusError(t *testing.T) {
	courseID := "37a0f027-15e6-47cc-a5d2-64183281087e"
	courseName := "Test Course"
	courseDuration := "10 months"

	courseRepositoryMock := new(storagemocks2.CourseRepository)
	courseRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("mooc.Course")).Return(nil)

	eventBusMock := new(eventmocks2.Bus)
	eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).Return(errors.New("something unexpected happened"))

	courseService := NewCourseService(courseRepositoryMock, eventBusMock)

	err := courseService.CreateCourse(context.Background(), courseID, courseName, courseDuration)

	courseRepositoryMock.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_CourseService_CreateCourse_Succeed(t *testing.T) {
	courseID := "37a0f027-15e6-47cc-a5d2-64183281087e"
	courseName := "Test Course"
	courseDuration := "10 months"

	course, err := mooc2.NewCourse(courseID, courseName, courseDuration)
	require.NoError(t, err)

	courseRepositoryMock := new(storagemocks2.CourseRepository)
	courseRepositoryMock.On("Save", mock.Anything, course).Return(nil)

	eventBusMock := new(eventmocks2.Bus)

	eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).Return(nil)

	courseService := NewCourseService(courseRepositoryMock, eventBusMock)

	err = courseService.CreateCourse(context.Background(), courseID, courseName, courseDuration)

	courseRepositoryMock.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
	assert.NoError(t, err)
}
