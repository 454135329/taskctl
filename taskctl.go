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

				task, err := tasks.OpenTask(name)
				if err != nil {
					return err
				}
				defer task.Close()

				err = task.Start()
				if err != nil {
					return err
				}

				fmt.Println("started task: ", name)
				return nil
			},
		},
		{
			Name:  "stop",
			Usage: "stop working on a task",
			Action: func(c *cli.Context) error {
				name := c.Args().First()

				task, err := tasks.OpenTask(name)
				if err != nil {
					return err
				}
				defer task.Close()

				err = task.Stop()
				if err != nil {
					return err
				}

				fmt.Println("stopped task: ", name)
				return nil
			},
		},
		{
			Name:  "done",
			Usage: "complete a task",
			Action: func(c *cli.Context) error {
				name := c.Args().First()

				task, err := tasks.OpenTask(name)
				if err != nil {
					return err
				}
				defer task.Close()

				err = task.Done()
				if err != nil {
					return err
				}

				fmt.Println("completed task: ", name)
				return nil
			},
		},
		{
			Name:  "remove",
			Usage: "remove a task",
			Action: func(c *cli.Context) error {
				name := c.Args().First()

				task, err := tasks.OpenTask(name)
				if err != nil {
					return err
				}

				task.Remove()

				fmt.Println("removed task: ", name)
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

				tasks, err := tasks.LoadTasks()
				if err != nil {
					return err
				}

				for _, task := range tasks {
					table.Append(task.ToArray())
				}

				table.Render()

				return nil
			},
		},
	}

	app.Run(os.Args)
}
