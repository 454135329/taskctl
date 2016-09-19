package tasks

import (
	"encoding/json"
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
func OpenTask(name string) Task {
	path := getTaskPath(name)

	if !fileExists(path) {
		return Task{Name: name, Status: "todo"}
	}

	file, _ := ioutil.ReadFile(path)

	var task Task
	json.Unmarshal(file, &task)

	return task
}

// Close writes changes to file
func (task *Task) Close() {
	path := getTaskPath(task.Name)
	data, _ := json.Marshal(task)

	_ = ioutil.WriteFile(path, data, 0644)
}

// Start changes status to in progress
func (task *Task) Start() {
	task.Status = StartStatus
	err := task.History.LogEvent(StartStatus)
	check(err)
}

// Stop changes status to stopped
func (task *Task) Stop() {
	task.Status = StopStatus
	err := task.History.LogEvent(StopStatus)
	check(err)
}

// Done changes status to done
func (task *Task) Done() {
	task.Status = DoneStatus
	err := task.History.LogEvent(DoneStatus)
	check(err)
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
