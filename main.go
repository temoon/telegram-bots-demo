package main

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/temoon/telegram-bots"

	"github.com/temoon/telegram-demo-bots/router"
)

func init() {
	bots.InitLog()
}

func main() {
	h := bots.Handler{}

	ctx, cancel := context.WithCancel(context.Background())
	defer h.Shutdown(cancel)

	go h.OnInterrupt(cancel)

	if err := h.Run(ctx, router.OnUpdate); err != nil {
		log.WithError(err).Fatal("main")
	}
}
