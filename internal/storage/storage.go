package storage

import (
	"github.com/pkg/errors"
	"log"
	"strconv"
)

var tasks map[uint]*Task

func init() {
	tasks = make(map[uint]*Task)

	if t, e := NewTask("Create new Task",
		"Test task creating"); e != nil {
		log.Panicln(e)
	} else if e = Add(t); e != nil {
		log.Panicln(e)
	}
}

func Tasks() []*Task {
	res := make([]*Task, 0, len(tasks))

	for _, t := range tasks {
		res = append(res, t)
	}

	return res
}

func Add(t *Task) error {
	if _, ok := tasks[t.Id()]; ok {
		return makeTaskExistError(true, t.Id())
	}

	tasks[t.Id()] = t
	return nil
}

func Update(t *Task) error {
	if _, ok := tasks[t.Id()]; !ok {
		return makeTaskExistError(false, t.Id())
	}

	tasks[t.Id()] = t
	return nil
}

func Delete(id uint) error {
	if _, ok := tasks[id]; !ok {
		return makeTaskExistError(false, id)
	}

	delete(tasks, id)
	return nil
}

func Get(id uint) (*Task, error) {
	if t, ok := tasks[id]; ok {
		return t, nil
	}

	return nil, makeTaskExistError(false, id)
}

var TaskExistError = errors.New("task already exist")
var TaskNotExistError = errors.New("task doesn't exist")

func makeTaskExistError(isExist bool, id uint) error {
	var e error
	if isExist {
		e = TaskExistError
	} else {
		e = TaskNotExistError
	}

	return errors.Wrap(e, strconv.FormatUint(uint64(id), 10))
}