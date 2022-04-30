package main

import (
	"context"

	"github.com/temoon/telegram-bots"

	"github.com/temoon/telegram-demo-bots/router"
)

func init() {
	bots.InitLog()
}

//goland:noinspection GoUnusedExportedFunction
func YandexHandler(ctx context.Context, req *bots.YandexRequest) (res *bots.YandexResponse, err error) {
	return bots.YandexHandler(ctx, req, router.OnUpdate)
}
