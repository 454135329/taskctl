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

func check(err error) {
	if err != nil {
		panic(err)
	}
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
	currStatus, _ := getCurrentStatus(task)

	if currStatus == status {
		err := errors.New("This task is already in " + status + " status")
		panic(err)
	}
}

func getStatuses(task string) []string {
	filename := getPath(task)
	content, err := ioutil.ReadFile(filename)
	check(err)

	return strings.Split(string(content), "\n")
}

func getCurrentStatus(task string) (string, string) {
	statuses := getStatuses(task)
	status := strings.Split(statuses[len(statuses)-2], "|")

	return status[0], status[1]
}

// WriteStatus writes task status to file
func WriteStatus(task string, status string) {
	filename := getPath(task)

	if fileExists(filename) {
		checkStatys(task, status)
	} else if status == "STOP" {
		err := errors.New("This task does not exist")
		panic(err)
	}

	curTime := time.Now()
	data := status + "|" + curTime.Format(time.RFC3339) + "\n"

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
		name := strings.TrimSuffix(filename, filepath.Ext(filename))
		status, _ := getCurrentStatus(name)

		data[i] = []string{name, status, "10 h"}
	}

	return data
}
