package telegram

import (
	"log"
	"strings"

	tb "gopkg.in/tucnak/telebot.v2"
)

func addTorrent(m *tb.Message) {
	magnetlink := strings.TrimPrefix(m.Text, "/add_torrent ")
	log.Println(magnetlink)
}
