package router

import (
	"context"

	"github.com/temoon/telegram-bots"

	"github.com/temoon/telegram-demo-bots/router/cmd"
	. "github.com/temoon/telegram-demo-bots/vars"
)

func OnUpdate(ctx context.Context, h *bots.Handler, u *bots.Update) (err error) {
	switch {
	case u.CheckCommand(CmdStart, CmdStartCode):
		err = cmd.RenderStart(ctx, h, u)
	case u.CheckCommand(CmdSettings, CmdSettingsCode):
		err = cmd.RenderSettings(ctx, h, u)
	case u.CheckCommand(CmdHelp, CmdHelpCode):
		err = cmd.RenderHelp(ctx, h, u)
	case u.CheckCommand(CmdMe, CmdMeCode):
		err = cmd.RenderMe(ctx, h, u)
	}

	return
}
