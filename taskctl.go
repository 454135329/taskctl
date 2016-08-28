package main

import (
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
	curTime := time.Now()
	filename := getPath(task)
	data := status + "|" + curTime.Format(time.RFC3339) + "\n"

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	check(err)

	defer f.Close()

	_, err = f.WriteString(data)
	check(err)
}

func getStatuses(task string) []string {
	filename := getPath(task)
	content, err := ioutil.ReadFile(filename)
	check(err)

	return strings.Split(string(content), "\n")
}

func getCurrentStatus(task string) string {
	statuses := getStatuses(task)

	return statuses[len(statuses)-2]
}

func main() {
	task := "ASS-2103"
	status := getCurrentStatus(task)

	fmt.Println(status)
}
