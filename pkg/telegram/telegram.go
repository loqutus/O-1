package telegram

import (
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	tb "gopkg.in/tucnak/telebot.v2"
)

func StartBot(wg *sync.WaitGroup) {
	defer wg.Done()
	token, err := getToken()
	if err != nil {
		log.Panic(err)
	}
	log.Println("Starting telegram bot")
	b, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "Hello World!")
	})
	b.Handle("/add_torrent", addTorrent)

	b.Start()
}
