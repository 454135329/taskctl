package tasks

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

// StartStatus represents status for started task
const StartStatus = "START"

// StopStatus represents status for stopped task
const StopStatus = "STOP"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func getCurrentDateTime() string {
	return time.Now().Format(time.RFC3339)
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

func getPath(task string) string {
	return getTasksDir() + "/" + task + ".txt"
}

func checkStatys(task string, status string) {
	statuses := getStatuses(task)
	currStatus, _ := getCurrentStatus(statuses)

	if currStatus == status {
		err := errors.New("This task is already in " + status + " status")
		panic(err)
	}
}

func getStatuses(task string) [][]string {
	filename := getPath(task)
	content, err := ioutil.ReadFile(filename)
	check(err)

	lines := strings.Split(string(content), "\n")

	if lines[len(lines)-1] == "" && len(lines) > 0 {
		lines = lines[:len(lines)-1]
	}

	statuses := make([][]string, len(lines))

	for i, status := range lines {
		statuses[i] = strings.Split(status, "|")
	}

	return statuses
}

func fillStatusesGap(statuses [][]string) [][]string {
	currentStatus, _ := getCurrentStatus(statuses)

	if currentStatus == StartStatus {
		statuses = append(statuses, []string{StopStatus, getCurrentDateTime()})
	}

	return statuses
}

func getCurrentStatus(statuses [][]string) (string, string) {
	status := statuses[len(statuses)-1]

	return status[0], status[1]
}

// WriteStatus writes task status to file
func WriteStatus(task string, status string) {
	filename := getPath(task)

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
		statuses := getStatuses(task)
		status, _ := getCurrentStatus(statuses)

		data[i] = []string{task, status, "10 h"}
	}

	return data
}
