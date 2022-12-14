package command

import (
	"context"

	"fmt"

	"strconv"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"
)

type getCommand struct {
	command
}

func (c *getCommand) Execute(ctx context.Context, args string) string {
	argsArr, err := extractArgs(args)
	if err != nil {
		return err.Error()
	} else if len(argsArr) == 0 {
		return ErrNoEnoughArgs.Error()
	}

	id, err := strconv.ParseUint(argsArr[0], 10, 64)
	if err != nil {
		return fmt.Sprintf("Cannot parse %s", argsArr[0])
	}

	uid := enumerator.Get(id)
	if uid == nil {
		return fmt.Sprintf("Cannot find task #%d", id)
	}

	t, err := c.service.GetTask(ctx, models.NewGetTaskData(uid))
	if err != nil {
		return err.Error()
	}

	return t.String()
}

func newGetCommand(s iService) *getCommand {
	return &getCommand{
		command{
			name:        "get",
			description: "get task info",
			subArgs:     "<ID>",
			service:     s},
	}
}
