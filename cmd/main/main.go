package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/kardianos/osext"
	"github.com/sausozluk/sozlukcli"
	"gopkg.in/ini.v1"
	"gopkg.in/urfave/cli.v1"
)

// ConfData :
const ConfData = `
[sozluk]
TOKEN = 
LASTENTRY = 
`

func getValueFromConf(key string) string {
	file := getCurrentDir() + "/sozluk.ini"

	cfg, _ := ini.Load([]byte(ConfData), file)

	value := cfg.Section("sozluk").Key(key).Value()

	return value
}

func setValueToConf(key string, value string) {
	file := getCurrentDir() + "/sozluk.ini"

	cfg, _ := ini.Load([]byte(ConfData), file)

	keyObj := cfg.Section("sozluk").Key(key)

	keyObj.SetValue(value)

	cfg.SaveTo(file)
}

func getCurrentDir() string {
	filename, _ := osext.ExecutableFolder()
	return filename
}

func main() {
	app := cli.NewApp()

	ready := sozlukcli.NewSozluk(getValueFromConf("TOKEN"))

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
	var email, password string

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
				if !ready {
					cli.ShowCommandHelp(c, "login")
				} else if entry != "" && topic != "" {
					id := sozlukcli.CreateEntry(topic, entry)

					if id != -1 {
						setValueToConf("LASTENTRY", strconv.Itoa(id))

						fmt.Printf("%s\n - http://sausozluk.net/entry/%d\n", sozlukcli.GetSlug(), id)
					}
				} else {
					cli.ShowCommandHelp(c, "write")
				}

				return nil
			},
		},
		{
			Name:    "login",
			Aliases: []string{"l"},
			Usage:   "Login without pain",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "email",
					Value:       "",
					Usage:       "email address",
					Destination: &email,
				},
				cli.StringFlag{
					Name:        "password",
					Value:       "",
					Usage:       "best secret password ever",
					Destination: &password,
				},
			},
			Action: func(c *cli.Context) error {
				if email != "" && password != "" {
					token, error := sozlukcli.DoLogin(email, password)

					if error == nil {
						setValueToConf("TOKEN", token)

						fmt.Printf("- Auth succes :)\n")
					} else {
						fmt.Printf("- Auth Failed :/\n")
					}
				} else {
					cli.ShowCommandHelp(c, "login")
				}

				return nil
			},
		},
		{
			Name:  "logout",
			Usage: "Just logout",
			Action: func(c *cli.Context) error {
				sozlukcli.DoLogout(getValueFromConf("TOKEN"))

				setValueToConf("TOKEN", "")

				fmt.Printf("- Done !\n")

				return nil
			},
		},
		{
			Name:  "delete",
			Usage: "Delete last entry",
			Action: func(c *cli.Context) error {
				lastEntry, error := strconv.Atoi(getValueFromConf("LASTENTRY"))

				if error != nil {
					fmt.Printf("- NOT Done :(\n")

					return nil
				}

				sozlukcli.DeleteEntry(lastEntry)

				setValueToConf("LASTENTRY", "")

				fmt.Printf("- Done !\n")

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
