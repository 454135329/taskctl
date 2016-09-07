package tasks

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	// StartStatus represents status for started task
	StartStatus = "START"
	// StopStatus represents status for stopped task
	StopStatus = "STOP"
)

var messages = map[string]string{
	"START": "In progress",
	"STOP":  "Stopped",
	"DONE":  "Completed",
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

func getTaskPath(task string) string {
	return getTasksDir() + "/" + task + ".txt"
}

func checkStatys(task string, status string) {
	history := getHistory(task)
	currStatus, _ := getCurrentStatus(history)

	if currStatus == status {
		err := errors.New("This task is already in " + status + " status")
		panic(err)
	}
}

func getHistory(task string) [][]string {
	filename := getTaskPath(task)
	content, err := ioutil.ReadFile(filename)
	check(err)

	lines := strings.Split(string(content), "\n")

	if lines[len(lines)-1] == "" && len(lines) > 0 {
		lines = lines[:len(lines)-1]
	}

	history := make([][]string, len(lines))

	for i, status := range lines {
		history[i] = strings.Split(status, "|")
	}

	return history
}

func fillHistoryGap(history [][]string) [][]string {
	currentStatus, _ := getCurrentStatus(history)

	if currentStatus == StartStatus {
		history = append(history, []string{StopStatus, getCurrentDateTime()})
	}

	return history
}

func getCurrentStatus(history [][]string) (string, string) {
	status := history[len(history)-1]

	return status[0], status[1]
}

func formatDuration(seconds int) string {
	hours := seconds / 3600
	minutes := seconds % 3600 / 60

	return strconv.Itoa(hours) + " h " + strconv.Itoa(minutes) + " m"
}

func getSpentTime(history [][]string) int {
	if len(history)%2 != 0 {
		err := errors.New("Wrong history records number")
		panic(err)
	}

	spentTime := 0

	for i := 0; i < len(history); i += 2 {
		startTime := parseDateTime(history[i][1])
		endTime := parseDateTime(history[i+1][1])
		timeDiff := endTime.Sub(startTime)

		spentTime += int(timeDiff.Seconds())
	}

	return spentTime
}

// WriteStatus writes task status to file
func WriteStatus(task string, status string) {
	filename := getTaskPath(task)

	if fileExists(filename) {
		checkStatys(task, status)
	} else if status == StopStatus {
		err := errors.New("This task does not exist")
		panic(err)
	}

	data := status + "|" + getCurrentDateTime() + "\n"

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	check(err)

	defer f.Close()

	_, err = f.WriteString(data)
	check(err)
}

// ListTasks return list of all task with logged time
func ListTasks() [][]string {
	files, err := ioutil.ReadDir(getTasksDir())
	if err != nil {
		log.Fatal(err)
	}

	data := make([][]string, len(files))

	for i, file := range files {
		filename := file.Name()
		task := strings.TrimSuffix(filename, filepath.Ext(filename))

		history := getHistory(task)
		status, _ := getCurrentStatus(history)
		history = fillHistoryGap(history)
		spentTime := getSpentTime(history)

		data[i] = []string{task, messages[status], formatDuration(spentTime)}
	}

	return data
}
