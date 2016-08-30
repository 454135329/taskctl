package main

import "./taskctl"

func main() {
	task := "ASS-2103"

	taskctl.WriteStatus(task, "START")

	// status, time := taskctl.GetCurrentStatus(task)

	// fmt.Println("Switched to \"" + status + "\" at " + time)

}
