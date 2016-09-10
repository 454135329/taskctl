package tasks

import (
	"os"
	"os/user"
	"strconv"
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

	return Task{Name: "TSK-001", Status: "start"}
}

// Close writes changes to file
func (task Task) Close() {

}

// Start changes status to in progress
func (task *Task) Start() {

}

// Stop changes status to stopped
func (task *Task) Stop() {

}

// Done changes status to done
func (task *Task) Done() {

}

func (task Task) getTaskPath() string {
	return getTasksDir() + "/" + task.Name + ".txt"
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

func formatDuration(seconds int) string {
	hours := seconds / 3600
	minutes := seconds % 3600 / 60

	return strconv.Itoa(hours) + " h " + strconv.Itoa(minutes) + " m"
}
