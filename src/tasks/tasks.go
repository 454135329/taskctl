package tasks

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Task is data structure to describe project task
type Task struct {
	Name    string
	Status  string
	History History
}

var messages = map[string]string{
	"todo":  "To do",
	"start": "In progress",
	"stop":  "Stopped",
	"done":  "Completed",
}

// OpenTask reads existing task or create new one
func OpenTask(name string) (Task, error) {
	if len(name) == 0 {
		return Task{}, errors.New("Empty task name")
	}

	path := getTaskPath(name)

	if !fileExists(path) {
		return Task{Name: name, Status: "todo"}, nil
	}

	file, _ := ioutil.ReadFile(path)

	var task Task
	json.Unmarshal(file, &task)

	return task, nil
}

// Close writes changes to file
func (task *Task) Close() {
	path := getTaskPath(task.Name)
	data, _ := json.Marshal(task)

	_ = ioutil.WriteFile(path, data, 0644)
}

// Start changes status to in progress
func (task *Task) Start() error {
	task.Status = StartStatus
	err := task.History.LogEvent(StartStatus)

	return err
}

// Stop changes status to stopped
func (task *Task) Stop() error {
	task.Status = StopStatus
	err := task.History.LogEvent(StopStatus)

	return err
}

// Done changes status to done
func (task *Task) Done() error {
	task.Status = DoneStatus
	err := task.History.LogEvent(DoneStatus)

	return err
}

// Remove deletes task from file system
func (task Task) Remove() {
	path := getTaskPath(task.Name)

	if fileExists(path) {
		os.Remove(path)
	}
}

// ToArray returns array with task name, status and logged time
func (task Task) ToArray() []string {
	return []string{
		task.Name,
		messages[task.Status],
		formatDuration(task.History.GetLoggedTime()),
	}
}

// LoadTasks scans tasks storage and loads all tasks
func LoadTasks() ([]Task, error) {
	var tasks []Task

	files, err := ioutil.ReadDir(getTasksDir())
	if err != nil {
		return tasks, err
	}

	for _, file := range files {
		fileName := file.Name()
		fileExtension := filepath.Ext(fileName)
		taskName := strings.TrimSuffix(fileName, fileExtension)

		task, err := OpenTask(taskName)
		if err == nil {
			tasks = append(tasks, task)
		}
	}

	return tasks, nil
}
