package main

import (
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {
	bot, _ := linebot.New("06ba3775aca06dc01fcf536baf3145c7", "/ppMUopJaNfW1sbqXSAd8fncob/p+yfD1r37r5P3SkMA0LHWUfjycs2c9gyykwpmLlFa8cKgF8cZfrZYZZxlQawzj43YNZd5IPWNx31WI42s6is5b9aCvBASgNT+Jt7z3X+ORJuSEvsPPlPlXymc9AdB04t89/1O/w1cDnyilFU=")

	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if _, err = bot.ReplyMessage(event.ReplyToken,
						linebot.NewTextMessage(message.Text)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})
}
