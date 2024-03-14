package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/list"
	"github.com/jedib0t/go-pretty/v6/text"
)

type StatusType string

const (
	NotStarted StatusType = "NotStarted"
	InProgress StatusType = "InProgress"
	Done       StatusType = "Done"
)

type TaskUpdate struct {
	Description string    `json:"description"`
	WrittenAt   time.Time `json:"writtenAt"`
}

type Task struct {
	Title       string       `json:"title"`
	Updates     []TaskUpdate `json:"updates"`
	Status      StatusType   `json:"status"`
	CompletedAt time.Time    `json:"completedAt"`
}

type TaskList []Task

func (tasks *TaskList) Add(title string) {

	todo := Task{
		Title:       title,
		Status:      NotStarted,
		CompletedAt: time.Time{},
	}

	*tasks = append(*tasks, todo)
}

func (tasks *TaskList) AddTask(task Task) {
	*tasks = append(*tasks, task)
}

func (tasks *TaskList) Get(id int) *Task {
	return &(*tasks)[id]
}

func (tasks *TaskList) Update(id int, update string) {
	(*tasks)[id].Updates = append((*tasks)[id].Updates, TaskUpdate{Description: update, WrittenAt: time.Now()})
}

func (tasks *TaskList) ChangeStatus(id int, status StatusType) {
	(*tasks)[id].Status = status

	if status == Done {
		(*tasks)[id].CompletedAt = time.Now()
	}
}

func (tasks *TaskList) Load(filename string) error {
	file, err := os.ReadFile(filename)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, tasks)
	if err != nil {
		return err
	}

	return nil
}

func (tasks *TaskList) Store(filename string) error {

	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func (tasks *TaskList) List() {
	l := list.NewWriter()
	l.SetStyle(list.StyleConnectedRounded)
	for tid, task := range *tasks {
		l.AppendItem(fmt.Sprintf("#%d: (%s - %s) %s", tid, text.FgHiGreen.Sprint(task.Status), task.CompletedAt.Format("Jan 2, 06"), text.FgHiBlue.Sprint(task.Title)))
		l.Indent()
		for uid, update := range task.Updates {
			l.AppendItem(fmt.Sprintf("##%d: %s %s", uid, update.WrittenAt.Format("Jan 2, 06"), text.FgHiMagenta.Sprint(update.Description)))
		}
		l.UnIndent()
	}

	fmt.Println(l.Render())
}
