package tasks

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"
)

// Event describes change of status
type Event struct {
	Status string
	Time   time.Time
}

// Task is data structure to describe project task
type Task struct {
	Name    string
	Status  string
	History []Event
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

	return Task{Name: name, Status: "todo", History: []Event{}}
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

	length := len(task.History)
	if length > 0 && task.History[length-1].Status == status {
		err := errors.New("This task already has " + messages[status] + " status")
		check(err)
	}

	event := Event{Status: status, Time: getCurrentDateTime()}

	task.Status = status
	task.History = append(task.History, event)
}

// Stop changes status to stopped
func (task *Task) Stop() {
	status := "stop"

	length := len(task.History)
	if length > 0 && task.History[length-1].Status == status {
		err := errors.New("This task already has " + messages[status] + " status")
		check(err)
	}

	event := Event{Status: status, Time: getCurrentDateTime()}

	task.Status = status
	task.History = append(task.History, event)
}

// Done changes status to done
func (task *Task) Done() {
	status := "done"
	task.Status = status

	length := len(task.History)
	if length > 0 && task.History[length-1].Status == status {
		err := errors.New("This task already has " + messages[status] + " status")
		check(err)
	}

	if length > 0 && task.History[length-1].Status == "stop" {
		task.History[length-1].Status = status
		return
	}

	event := Event{Status: status, Time: getCurrentDateTime()}

	task.History = append(task.History, event)
}

// ToArray returns array with task name, status and logged time
func (task Task) ToArray() []string {
	return []string{
		task.Name,
		messages[task.Status],
		formatDuration(task.getLoggedTime()),
	}
}

func (task Task) getLoggedTime() int {
	history := task.History

	if len(history)%2 != 0 {
		err := errors.New("Wrong history records number")
		panic(err)
	}

	loggedTime := 0

	for i := 0; i < len(history); i += 2 {
		startTime := history[i].Time
		endTime := history[i+1].Time
		timeDiff := endTime.Sub(startTime)

		loggedTime += int(timeDiff.Seconds())
	}

	return loggedTime
}

// LoadTasks scans tasks storage and loads all tasks
func LoadTasks() []Task {
	files, err := ioutil.ReadDir(getTasksDir())
	check(err)

	tasks := make([]Task, len(files))

	for i, file := range files {
		fileName := file.Name()
		taskName := strings.TrimSuffix(fileName, filepath.Ext(fileName))

		tasks[i] = OpenTask(taskName)
	}

	return tasks
}
