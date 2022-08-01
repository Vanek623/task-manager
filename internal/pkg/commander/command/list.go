package command

import (
	"strings"
)

type listCommand struct {
	command
}

func (c *listCommand) Execute(args string) string {
	tasks := c.manager.List()

	if len(tasks) == 0 {
		return "There are no tasks!"
	}

	out := make([]string, 0, len(tasks))
	for _, t := range tasks {
		out = append(out, t.String())
	}

	return strings.Join(out, "\n")
}

func newListCommand(m iTaskManager) *listCommand {
	return &listCommand{
		command{
			name:        "list",
			description: "tasks list",
			manager:     m},
	}
}
