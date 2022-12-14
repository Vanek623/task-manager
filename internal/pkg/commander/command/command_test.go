package command

import (
	"testing"
	"time"

	"github.com/google/uuid"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestAddCommand_Execute(t *testing.T) {
	t.Run("add task to storage", func(t *testing.T) {
		// arrange
		f := addCommandSetUp(t)
		f.service.EXPECT().AddTask(gomock.Any(), models.NewAddTaskData("test", "test")).
			Return(uint64(1), nil)
		// act
		res := f.command.Execute(f.Ctx, "test test")
		// assert
		assert.Equal(t, res, "Task #1 added")
	})

	t.Run("task doesn't add if has no args", func(t *testing.T) {
		// arrange
		f := addCommandSetUp(t)
		// act
		res := f.command.Execute(f.Ctx, "")
		// assert
		assert.Equal(t, res, "has no enough arguments")
	})

	t.Run("task doesn't add if has internal error", func(t *testing.T) {
		// arrange
		f := addCommandSetUp(t)
		f.service.EXPECT().AddTask(gomock.Any(), models.NewAddTaskData("test", "test")).
			Return(uint64(0), errors.New("internal error"))
		// act
		res := f.command.Execute(f.Ctx, "test test")
		// assert
		assert.Equal(t, res, "internal error")
	})
}

func TestDeleteCommand_Execute(t *testing.T) {
	t.Run("delete task from storage", func(t *testing.T) {
		// arrange
		f := delCommandSetUp(t)
		id := uuid.New()
		f.service.EXPECT().DeleteTask(gomock.Any(), models.NewDeleteTaskData(&id)).Return(nil)
		// act
		res := f.command.Execute(f.Ctx, id.String())
		// assert
		assert.Equal(t, res, "Task deleted")
	})

	t.Run("task doesn't delete if has no args", func(t *testing.T) {
		// arrange
		f := delCommandSetUp(t)
		// act
		res := f.command.Execute(f.Ctx, "")
		// assert
		assert.Equal(t, res, "has no enough arguments")
	})

	t.Run("doesn't delete task if args is invalid", func(t *testing.T) {
		// arrange
		f := delCommandSetUp(t)
		// act
		res := f.command.Execute(f.Ctx, "a")
		// assert
		assert.NotEqual(t, res, "Task deleted")
	})

	t.Run("doesn't delete task if has internal error", func(t *testing.T) {
		// arrange
		f := delCommandSetUp(t)
		id := uuid.New()
		f.service.EXPECT().DeleteTask(gomock.Any(), models.NewDeleteTaskData(&id)).Return(errors.New(""))
		// act
		res := f.command.Execute(f.Ctx, id.String())
		// assert
		assert.Equal(t, res, "internal error")
	})
}

func TestGetCommand_Execute(t *testing.T) {
	t.Run("get task form storage", func(t *testing.T) {
		// arrange
		f := getCommandSetUp(t)
		actual := models.NewDetailedTask("test", "test", time.Now())
		id := uuid.New()
		f.service.EXPECT().GetTask(gomock.Any(), models.NewGetTaskData(&id)).Return(actual, nil)
		// act
		res := f.command.Execute(f.Ctx, id.String())
		// assert
		assert.Equal(t, res, actual.String())
	})

	t.Run("doesn't get task if has internal error", func(t *testing.T) {
		// arrange
		f := getCommandSetUp(t)
		actual := models.NewDetailedTask("test", "test", time.Now())
		id := uuid.New()
		f.service.EXPECT().GetTask(gomock.Any(), models.NewGetTaskData(&id)).Return(nil, errors.New(""))
		// act
		res := f.command.Execute(f.Ctx, id.String())
		// assert
		assert.Equal(t, res, "internal error")
	})

	t.Run("doesn't get task if args is invalid", func(t *testing.T) {
		// arrange
		f := getCommandSetUp(t)
		actual := models.NewDetailedTask("test", "test", time.Now())
		id := uuid.New()
		f.service.EXPECT().GetTask(gomock.Any(), models.NewGetTaskData(&id)).Return(nil, errors.New(""))
		// act
		res := f.command.Execute(f.Ctx, id.String())
		// assert
		assert.NotEqual(t, res, actual.String())
	})
}
