package main

import (
	"errors"
	"fmt"
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

func getPath(task string) string {
	usr, _ := user.Current()

	return usr.HomeDir + "/.taskctl/tasks/" + task + ".txt"
}

func writeStatus(task string, status string) {
	checkStatys(task, status)

	curTime := time.Now()
	filename := getPath(task)
	data := status + "|" + curTime.Format(time.RFC3339) + "\n"

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	check(err)

	defer f.Close()

	_, err = f.WriteString(data)
	check(err)
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

func main() {
	task := "ASS-2103"
	status, time := getCurrentStatus(task)

	fmt.Println("Switched to \"" + status + "\" at " + time)

	writeStatus(task, status)
}
