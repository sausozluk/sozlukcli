package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/damnever/cc"
	"github.com/gosimple/slug"
	"gopkg.in/resty.v0"
	"gopkg.in/urfave/cli.v1"
)

var token string

type CheckResponse struct {
	Success bool       `json:"success"`
	Data    *CheckData `json:"data"`
}

type CheckData struct {
	IsAlive bool   `json: "isAlive"`
	User_Id string `json: "user_id, omitempty"`
	Slug    string `json: "slug, omitempty"`
	Unread  int    `json: "unread`
}

type SearchResponse struct {
	Success bool        `json:"success"`
	Data    *SearchData `json:"data"`
}

type SearchData struct {
	Topics []SearchDataTopics `json:"topics"`
	Users  []SearchDataUsers  `json:"users"`
}

type SearchDataTopics struct {
	ID    int    `json:"id"`
	Slug  string `json:"slug"`
	Title string `json:"title"`
}

type SearchDataUsers struct {
	Username string `json:"username"`
	Slug     string `json:"slug"`
}

type EntryCreateResponse struct {
	Success bool             `json:"success"`
	Data    *EntryCreateData `json:"data"`
}

type EntryCreateData struct {
	ID int `json:"id"`
}

type TopicCreateResponse struct {
	Success bool `json:"success"`
	ID      int  `json:"entry_id"`
}

func getToken() string {
	c, _ := cc.NewConfigFromFile("config.json")
	return c.String("token")
}

func checkToken(token string) (*CheckResponse, bool) {
	resp, err := resty.R().
		SetHeader("token", token).
		SetResult(CheckResponse{}).
		Post("http://sausozluk.net/service/proxy/api/v1/sessions/check")

	checkResp := resp.Result().(*CheckResponse)

	return checkResp, err == nil && checkResp.Data.IsAlive
}

func start() {
	checkResp, done := checkToken(token)

	if done {
		fmt.Printf("Hi %s!\n", checkResp.Data.Slug)
	} else {
		fmt.Printf("This token (%s) not alive! Please check config.json!\n", token)
	}
}

func topicAlreadyExist(topic string) (*SearchDataTopics, bool) {
	resp, err := resty.R().
		SetHeader("token", token).
		SetResult(SearchResponse{}).
		Get("http://sausozluk.net/service/proxy/api/v1/search?q=" + topic)

	searchResp := resp.Result().(*SearchResponse)
	done := err == nil && searchResp.Success

	if done {
		targetSlug := slug.Make(topic)

		for index, element := range searchResp.Data.Topics {
			_ = index
			if element.Slug == targetSlug {
				return &element, true
			}
		}
	}

	return nil, false
}

func withCreateTopic(topic string, entry string) (*TopicCreateResponse, bool) {
	payload := fmt.Sprintf(`{"topic": {"title": "%s"}, "entry": {"text": "%s"}}`, topic, entry)

	resp, err := resty.R().
		SetHeader("token", token).
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		SetResult(TopicCreateResponse{}).
		Post("http://sausozluk.net/service/proxy/api/v1/topics")

	topicCreateResp := resp.Result().(*TopicCreateResponse)
	done := err == nil && topicCreateResp.Success

	return topicCreateResp, done
}

func intoExistTopic(topicID string, entry string) (*EntryCreateResponse, bool) {
	resp, err := resty.R().
		SetHeader("token", token).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{"text": entry, "topic_id": topicID}).
		SetResult(EntryCreateResponse{}).
		Post("http://sausozluk.net/service/proxy/api/v1/entries")

	entryCreateResp := resp.Result().(*EntryCreateResponse)
	done := err == nil && entryCreateResp.Success

	return entryCreateResp, done
}

func writeEntry(topic string, entry string) int {
	foundedTopic, exist := topicAlreadyExist(topic)

	if exist {
		entryCreateResp, done := intoExistTopic(strconv.Itoa(foundedTopic.ID), entry)

		if done {
			return entryCreateResp.Data.ID
		}

		return -1
	}

	topicCreateResp, done := withCreateTopic(topic, entry)

	if done {
		return topicCreateResp.ID
	}

	return -1
}

func main() {
	token = getToken()

	app := cli.NewApp()
	app.EnableBashCompletion = true

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

	var entry string
	var topic string

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
					id := writeEntry(topic, entry)
					fmt.Printf("Entry ready here http://sausozluk.net/entry/%d\n", id)
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
