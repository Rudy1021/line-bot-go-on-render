package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {
	bot, _ := linebot.New("06ba3775aca06dc01fcf536baf3145c7",
		"/ppMUopJaNfW1sbqXSAd8fncob/p+yfD1r37r5P3SkMA0LHWUfjycs2c9gyykwpmLlFa8cKgF8cZfrZYZZxlQawzj43YNZd5IPWNx31WI42s6is5b9aCvBASgNT+Jt7z3X+ORJuSEvsPPlPlXymc9AdB04t89/1O/w1cDnyilFU=")
	var timestamp time.Time
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

					if message.Text == "打卡" {
						loc, _ := time.LoadLocation("Asia/Taipei")

						// 轉換為UTC+8時間
						timestamp = event.Timestamp.In(loc)

						newTimestamp := timestamp.Add(time.Minute * 10)
						newTimestamp = newTimestamp.Add(time.Hour * 9)

						msg := "上崗時間：" + timestamp.Format("15:04:05") + "\n" + "下崗時間" + newTimestamp.Format("15:04:05")

						if _, err = bot.ReplyMessage(event.ReplyToken,
							linebot.NewTextMessage(msg)).Do(); err != nil {
							log.Print(err)
						}

					}
					if _, err = bot.ReplyMessage(event.ReplyToken,
						linebot.NewTextMessage("????")).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
