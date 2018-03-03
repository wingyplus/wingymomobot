package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/line/line-bot-sdk-go/linebot/httphandler"
)

var (
	channelAccessToken = "u3sBY3X239Uk1JlkcWHPMlMLX+aRICHK23Aik1ltAmtbAe4o67mW2iSL98voefpZs4MtaE8A8zXeEyg4kX2zuo+CNHJdu94xaPdhKgWSkDJtvicyUoq1qdoHhaH83zqkDezkiMW6NARQC0sTh/a7bgdB04t89/1O/w1cDnyilFU="
	channelSecret      = "0f1959754517b93dec632f1877ba618a"
)

func main() {
	var (
		handler *httphandler.WebhookHandler
		bot     *linebot.Client
		err     error
	)

	if handler, err = httphandler.New(channelSecret, channelAccessToken); err != nil {
		panic(err)
	}
	if bot, err = handler.NewClient(); err != nil {
		panic(err)
	}

	handler.HandleEvents(NewBot(bot).HandleEvents)

	http.Handle("/callback", handler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func NewBot(bot *linebot.Client) *WingyMomoBot {
	return &WingyMomoBot{
		LineBot: bot,
	}
}

type WingyMomoBot struct {
	LineBot *linebot.Client
}

func (bot *WingyMomoBot) HandleEvents(events []*linebot.Event, r *http.Request) {
	log.Println("event: received")

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch msg := event.Message.(type) {
			case *linebot.TextMessage:
				switch strings.Trim(msg.Text, " ") {
				case "ปุจฉา":
					bot.reply(event.ReplyToken, linebot.NewTemplateMessage("ทดลองส่ง", linebot.NewButtonsTemplate("", "วิสัจฉนา", "คุณต้องการค้นหาสิ่งใด", hospital(), clinic())))
				default:
					fmt.Println(msg.Text)
					if _, err := bot.reply(event.ReplyToken, linebot.NewTextMessage("รักเธอ <3")); err != nil {
						log.Println("text message reply error:", err)
					}
				}

			case *linebot.LocationMessage:
				fmt.Println("your location", msg.Latitude, msg.Longitude)
				if _, err := bot.reply(event.ReplyToken, linebot.NewTextMessage(fmt.Sprintf("location ของคุณคือ (%f, %f)", msg.Latitude, msg.Longitude))); err != nil {
					log.Println("location reply error:", err)
				}
			}
		} else if event.Type == linebot.EventTypePostback {
			if _, err := bot.reply(event.ReplyToken, linebot.NewTextMessage(fmt.Sprintf("คุณเลือก %s", event.Postback.Data))); err != nil {
				log.Println("postback reply error:", err)
			}
		}
	}
}

func (bot *WingyMomoBot) reply(replyToken string, messages ...linebot.Message) (*linebot.BasicResponse, error) {
	return bot.LineBot.ReplyMessage(replyToken, messages...).Do()
}

func hospital() *linebot.PostbackTemplateAction {
	return &linebot.PostbackTemplateAction{
		Label: "โรงพยาบาล",
		Data:  "hospital",
	}
}

func clinic() *linebot.PostbackTemplateAction {
	return &linebot.PostbackTemplateAction{
		Label: "คลีนิค",
		Data:  "clinic",
	}
}
