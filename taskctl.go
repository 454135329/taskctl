package main

import (
	"fmt"
	"os"

	"github.com/d-kurochkin/taskctl/src/tasks"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "start working on a task",
			Action: func(c *cli.Context) error {
				task := c.Args().First()
				tasks.WriteStatus(task, "START")
				fmt.Println("started task: ", task)
				return nil
			},
		},
		{
			Name:  "stop",
			Usage: "complete a task",
			Action: func(c *cli.Context) error {
				task := c.Args().First()
				tasks.WriteStatus(task, "STOP")
				fmt.Println("stopped task: ", task)
				return nil
			},
		},
	}

	app.Run(os.Args)
}
