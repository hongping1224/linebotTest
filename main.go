package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	client := &http.Client{}
	bot, err := linebot.New("6635ba429ac2a1639a9d06562c0b843f", "h6gNcNbSRgJ3xioe0Ygh6w2qxgt3m4aNhZssGKJG0QtEgbv9EPvr+wLti6Ij9aUjbLWsBAed5BPA2oU4f0omWoSOwK5OYtRrA7OQeE4Edf4dTEDcXXmuMM/XMVPlamqSYnc6Ts0Z4c1pBh39F0ul9wdB04t89/1O/w1cDnyilFU=", linebot.WithHTTPClient(client))
	if err != nil {
		fmt.Println(err)
		return
	}
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
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						log.Print(err)
					}
				case *linebot.StickerMessage:
					replyMessage := fmt.Sprintf(
						"sticker id is %s, stickerResourceType is %s", message.StickerID, message.StickerResourceType)
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})
	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	if err := http.ListenAndServe(":65000", nil); err != nil {
		log.Fatal(err)
	}
}
