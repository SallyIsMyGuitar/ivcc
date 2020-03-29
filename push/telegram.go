package push

import (
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Telegram implements the Telegram messenger
type Telegram struct {
	sync.Mutex
	bot   *tgbotapi.BotAPI
	chats map[int64]struct{}
}

type telegramConfig struct {
	Token string
	Chats []int64
}

// NewTelegramMessenger creates new pushover messenger
func NewTelegramMessenger(token string, chats []int64) *Telegram {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.FATAL.Fatal("telegram: missing token")
	}

	m := &Telegram{
		bot:   bot,
		chats: make(map[int64]struct{}),
	}

	for _, chat := range chats {
		m.chats[chat] = struct{}{}
	}

	go m.trackChats()

	return m
}

// trackChats captures ids of all chats that bot participates in
func (m *Telegram) trackChats() {
	conf := tgbotapi.NewUpdate(0)
	conf.Timeout = 1000

	updates, err := m.bot.GetUpdatesChan(conf)
	if err != nil {
		log.ERROR.Printf("telegram: %v", err)
	}

	for update := range updates {
		m.Lock()
		if _, ok := m.chats[update.Message.Chat.ID]; !ok {
			log.INFO.Printf("telegram: new chat id: %d", update.Message.Chat.ID)
			// m.chats[update.Message.Chat.ID] = struct{}{}
		}
		m.Unlock()
	}
}

// Send sends to all receivers
func (m *Telegram) Send(event Event, title, msg string) {
	m.Lock()
	for chat := range m.chats {
		go func(chat int64) {
			log.TRACE.Printf("telegram: sending to %d", chat)

			msg := tgbotapi.NewMessage(chat, msg)
			if _, err := m.bot.Send(msg); err != nil {
				log.ERROR.Print(err)
			}
		}(chat)
	}
	m.Unlock()
}
