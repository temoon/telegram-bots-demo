package cmd

import (
	"context"

	"github.com/temoon/go-telegram-bots-api"
	. "github.com/temoon/go-telegram-bots-api/helpers"
	"github.com/temoon/go-telegram-bots-api/requests"
	"github.com/temoon/telegram-bots"

	. "github.com/temoon/telegram-demo-bots/vars"
)

func RenderStart(ctx context.Context, h *bots.Handler, u *bots.Update) (err error) {
	text := "Главное меню"

	dataMe := bots.CallbackData{
		Command: CmdMeCode,
	}

	dataSettings := bots.CallbackData{
		Command: CmdSettingsCode,
	}

	dataHelp := bots.CallbackData{
		Command: CmdHelpCode,
	}

	keyboard := telegram.InlineKeyboardMarkup{
		InlineKeyboard: [][]telegram.InlineKeyboardButton{
			{
				{
					Text:         CmdMeText,
					CallbackData: String(dataMe.String()),
				},
			},
			{
				{
					Text:         CmdSettingsText,
					CallbackData: String(dataSettings.String()),
				},
				{
					Text:         CmdHelpText,
					CallbackData: String(dataHelp.String()),
				},
			},
		},
	}

	var message telegram.Request
	if u.NewMessageRequired() {
		message = &requests.SendMessage{
			ChatId:      u.ChatId,
			ParseMode:   String(telegram.ParseModeMarkdown),
			Text:        text,
			ReplyMarkup: &keyboard,
		}
	} else {
		message = &requests.EditMessageText{
			ChatId:      u.ChatId,
			MessageId:   Int32(u.MessageId),
			ParseMode:   String(telegram.ParseModeMarkdown),
			Text:        text,
			ReplyMarkup: &keyboard,
		}
	}

	if _, err = message.Call(ctx, h.GetBot()); err != nil {
		return
	}

	return
}
