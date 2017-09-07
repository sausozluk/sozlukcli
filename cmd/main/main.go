package main

import (
	"fmt"
	"os"
	"time"

	"github.com/damnever/cc"
	"github.com/sausozluk/sozlukcli"
	"gopkg.in/urfave/cli.v1"
)

func getToken() string {
	c, _ := cc.NewConfigFromFile("config.json")
	return c.String("token")
}

func main() {
	app := cli.NewApp()

	ready := sozlukcli.NewSozluk(getToken())

	if !ready {
		return
	}

	app.Name = "sozluk-cli"
	app.Usage = "saü sözlük command-line interface for n3rds"
	app.Version = "0.0.1"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Eray Arslan",
			Email: "relfishere@gmail.com",
		},
	}

	var entry, topic string

	app.Commands = []cli.Command{
		{
			Name:    "write",
			Aliases: []string{"w"},
			Usage:   "Write entry into topic",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "entry",
					Value:       "",
					Usage:       "write something for entry",
					Destination: &entry,
				},
				cli.StringFlag{
					Name:        "topic",
					Value:       "",
					Usage:       "target topic name",
					Destination: &topic,
				},
			},
			Action: func(c *cli.Context) error {
				if entry != "" && topic != "" {
					id := sozlukcli.CreateEntry(topic, entry)

					if id != -1 {
						fmt.Printf("Entry ready here http://sausozluk.net/entry/%d\n", id)
					}
				} else {
					cli.ShowCommandHelp(c, "write")
				}

				return nil
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		cli.ShowAppHelp(c)
		return nil
	}

	app.Run(os.Args)
}
