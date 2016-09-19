package tasks

import (
	"os"
	"os/user"
	"strconv"
	"strings"
	"time"
)

func getCurrentDateTime() time.Time {
	return time.Now()
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

func getTaskPath(name string) string {
	name = strings.ToUpper(name)

	return getTasksDir() + "/" + name + ".json"
}

func formatDuration(seconds int) string {
	hours := seconds / 3600
	minutes := seconds % 3600 / 60

	return strconv.Itoa(hours) + " h " + strconv.Itoa(minutes) + " m"
}
