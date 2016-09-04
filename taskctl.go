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
		{
			Name:  "list",
			Usage: "list all tasks",
			Action: func(c *cli.Context) error {
				data := tasks.ListTasks()

				table := tablewriter.NewWriter(os.Stdout)
				table.SetRowLine(true)
				table.SetRowSeparator("-")
				table.SetHeader([]string{"Task", "Status", "Logged time"})

				for _, v := range data {
					table.Append(v)
				}
				table.Render()
				return nil
			},
		},
	}

	app.Run(os.Args)
}
