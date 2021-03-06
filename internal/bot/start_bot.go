package bot

import (
	"context"
	app "github.com/VMAnalytic/alarm-bot/internal"
	"gopkg.in/tucnak/telebot.v2"
)

type StartBot struct {
	storage app.UserStorage
	commonBot
}

func NewStartBot(storage app.UserStorage) *StartBot {
	return &StartBot{storage: storage}
}

func (b *StartBot) Register(ctx context.Context, tgBot *telebot.Bot, errCh chan<- error) {
	b.init(tgBot, errCh)
	tgBot.Handle(CommandStart, func(m *telebot.Message) {
		var (
			err error
			uid = m.Sender.ID
		)

		exists, err := b.storage.Exists(ctx, uid)
		if err != nil {
			b.handleError(err, uid)
			return
		}

		if !exists {
			user := app.NewUser(uid)
			err = b.storage.Add(ctx, user)

			if err != nil {
				b.handleError(err, uid)
				return
			}
		}

		//mainMenu.Inline(
		//	mainMenu.Row(btnAlarm),
		//	mainMenu.Row(btnAddContact),
		//	mainMenu.Row(btnMyInfo, btnHelp),
		//)

		mainMenu.Reply(
			mainMenu.Row(btnAlarm),
		)

		_, err = tgBot.Send(m.Sender, "Welcome to alarm bot!", mainMenu)
		if err != nil {
			b.handleError(err, uid)
		}
	})
}

func (b *StartBot) Help() string {
	panic("implement me")
}
