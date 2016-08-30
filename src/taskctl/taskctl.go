package taskctl

import (
	"errors"
	"io/ioutil"
	"os"
	"os/user"
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

func getPath(task string) string {
	usr, _ := user.Current()
	path := usr.HomeDir + "/.taskctl/tasks/" + task + ".txt"

	return path
}

func checkStatys(task string, status string) {
	currStatus, _ := GetCurrentStatus(task)

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

// GetCurrentStatus returns current status of given task
func GetCurrentStatus(task string) (string, string) {
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
