package main

import (
	"os"
	"os/user"
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
	filename := getPath(task)
	data := status + "|" + "time" + "\n"

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	check(err)

	defer f.Close()

	_, err = f.WriteString(data)
	check(err)
}

func main() {
	task := "ASS-2103"
	status := "START"
	writeStatus(task, status)
}
