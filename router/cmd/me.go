package cmd

import (
	"context"
	"strconv"
	"strings"

	"github.com/temoon/go-telegram-bots-api"
	. "github.com/temoon/go-telegram-bots-api/helpers"
	"github.com/temoon/go-telegram-bots-api/requests"
	"github.com/temoon/telegram-bots"

	. "github.com/temoon/telegram-demo-bots/vars"
)

func RenderMe(ctx context.Context, h *bots.Handler, u *bots.Update) (err error) {
	text := "*Информация о пользователе*\n\n"
	text += "*ID:* " + MarkdownUserMention(strconv.FormatInt(u.UserId, 10), u.UserId) + "\n"
	text += "*Пользователь:* " + Markdown(StringOrEmpty(u.User.Username)) + "\n"
	text += "*Имя:* " + Markdown(strings.TrimSpace(u.User.FirstName+" "+StringOrEmpty(u.User.LastName))) + "\n"
	text += "*Язык:* " + StringOrEmpty(u.User.LanguageCode)

	dataStart := bots.CallbackData{
		Command: CmdStartCode,
	}

	keyboard := telegram.InlineKeyboardMarkup{
		InlineKeyboard: [][]telegram.InlineKeyboardButton{
			{
				{
					Text:         CmdStartText,
					CallbackData: String(dataStart.String()),
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
