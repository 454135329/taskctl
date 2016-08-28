package main

import (
	"fmt"
	"os/user"
)

func getPath(task string) string {
	usr, _ := user.Current()

	return usr.HomeDir + "/.taskctl/tasks/" + task + ".txt"
}

func main() {
	path := getPath("ASS-0001")
	fmt.Println(path)
}
