package tasks

import (
	"encoding/json"
	"io/ioutil"
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
func OpenTask(name string) Task {
	path := getTaskPath(name)

	if fileExists(path) {
		file, _ := ioutil.ReadFile(path)

		var task Task
		json.Unmarshal(file, &task)

		return task
	}

	return Task{Name: name, Status: "todo"}
}

// Close writes changes to file
func (task *Task) Close() {
	path := getTaskPath(task.Name)
	data, _ := json.Marshal(task)

	_ = ioutil.WriteFile(path, data, 0644)
}

// Start changes status to in progress
func (task *Task) Start() {
	status := "start"
	task.Status = status
	err := task.History.LogEvent(status)
	check(err)
}

// Stop changes status to stopped
func (task *Task) Stop() {
	status := "stop"
	task.Status = status
	err := task.History.LogEvent(status)
	check(err)
}

// Done changes status to done
func (task *Task) Done() {
	status := "done"
	task.Status = status
	err := task.History.LogEvent(status)
	check(err)
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
func LoadTasks() []Task {
	files, err := ioutil.ReadDir(getTasksDir())
	check(err)

	var tasks []Task

	for _, file := range files {
		fileName := file.Name()
		fileExtension := filepath.Ext(fileName)
		taskName := strings.TrimSuffix(fileName, fileExtension)

		tasks = append(tasks, OpenTask(taskName))
	}

	return tasks
}
