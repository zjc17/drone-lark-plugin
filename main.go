package main

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type (
	Plugin struct {
		Token string // token for custom lark bot
	}
)

var (
	version = "unknown"
)

func main() {
	app := cli.NewApp()
	app.Name = "drone lark plugin"
	app.Usage = "drone lark plugin"
	app.Action = run
	app.Version = version
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "token",
			Usage:   "you can get the access token when you add a bot in a group.",
			EnvVars: []string{"TOKEN", "PLUGIN_TOKEN"},
		},
		&cli.StringFlag{
			Name:    "color",
			Usage:   "color the title for easier identification.",
			EnvVars: []string{"MESSAGE_COLOR", "PLUGIN_COLOR"},
		},
		&cli.StringFlag{
			Name:    "title",
			Usage:   "message title.",
			EnvVars: []string{"MESSAGE_TITLE", "PLUGIN_TITLE"},
		},
		&cli.StringFlag{
			Name:    "content",
			Usage:   "message content.",
			EnvVars: []string{"MESSAGE_CONTENT", "PLUGIN_CONTENT"},
		},
		&cli.StringFlag{
			Name:    "commit.sha",
			Usage:   "git commit sha",
			EnvVars: []string{"DRONE_COMMIT_SHA"},
			Value:   "00000000",
		},
		&cli.StringFlag{
			Name:    "commit.ref",
			Usage:   "git commit ref",
			EnvVars: []string{"DRONE_COMMIT_REF"},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) (err error) {
	template := DefaultTemplate{
		Color:           c.String("color"),
		Title:           c.String("title"),
		MarkdownContent: c.String("content"),
	}
	value, err := json.Marshal(template.MarkdownContent)
	if err == nil {
		template.MarkdownContent = string(value)
	}
	content, err := template.Content()
	fmt.Println("body", content)
	if err != nil {
		log.Fatal(err)
	}
	url := fmt.Sprintf("https://open.feishu.cn/open-apis/bot/v2/hook/%s", c.String("token"))
	contentType := "application/json"
	resp, err := http.Post(url, contentType, strings.NewReader(content))
	if err != nil {
		return err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	log.Println(bodyString)
	return
}
