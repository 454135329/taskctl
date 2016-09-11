package tasks

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
	"strconv"
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
	task.Status = "start"
}

// Stop changes status to stopped
func (task *Task) Stop() {
	task.Status = "stop"
}

// Done changes status to done
func (task *Task) Done() {
	task.Status = "done"
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func getCurrentDateTime() string {
	return time.Now().Format(time.RFC3339)
}

func parseDateTime(dateTime string) time.Time {
	parsedTime, err := time.Parse(time.RFC3339, dateTime)
	check(err)

	return parsedTime
}

func fileExists(path string) bool {
	_, err := os.Stat(path)

	if err != nil {
		return false
	}

	return true
}

func getTasksDir() string {
	usr, _ := user.Current()

	return usr.HomeDir + "/.taskctl/tasks"
}

func getTaskPath(name string) string {
	name = strings.ToUpper(name)

	return getTasksDir() + "/" + name + ".txt"
}

func formatDuration(seconds int) string {
	hours := seconds / 3600
	minutes := seconds % 3600 / 60

	return strconv.Itoa(hours) + " h " + strconv.Itoa(minutes) + " m"
}
