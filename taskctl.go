package main

import (
	"fmt"
	"os"

	"github.com/d-kurochkin/taskctl/src/tasks"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "start working on a task",
			Action: func(c *cli.Context) error {
				name := c.Args().First()

				task := tasks.OpenTask(name)
				defer task.Close()

				task.Start()

				fmt.Println("started task: ", name)
				return nil
			},
		},
		{
			Name:  "stop",
			Usage: "stop working on a task",
			Action: func(c *cli.Context) error {
				name := c.Args().First()

				task := tasks.OpenTask(name)
				defer task.Close()

				task.Stop()

				fmt.Println("stopped task: ", name)
				return nil
			},
		},
		{
			Name:  "done",
			Usage: "complete a task",
			Action: func(c *cli.Context) error {
				name := c.Args().First()

				task := tasks.OpenTask(name)
				defer task.Close()

				task.Done()

				fmt.Println("completed task: ", name)
				return nil
			},
		},
		{
			Name:  "list",
			Usage: "list all tasks",
			Action: func(c *cli.Context) error {
				table := tablewriter.NewWriter(os.Stdout)
				table.SetRowLine(true)
				table.SetRowSeparator("-")
				table.SetHeader([]string{"Task", "Status", "Logged time"})

				for _, task := range tasks.LoadTasks() {
					table.Append(task.ToArray())
				}

				table.Render()

				return nil
			},
		},
	}

	app.Run(os.Args)
}
