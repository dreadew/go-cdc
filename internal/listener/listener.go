package listener

import (
	"context"
	"log"
	"time"

	"github.com/dreadew/go-cdc/internal/client/db"
	"github.com/dreadew/go-cdc/internal/config"
	"github.com/lib/pq"
)

func StartListener(ctx context.Context, cfg *config.Config) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := logArticles(ctx, cfg); err != nil {
				log.Println(err)
			}
		case <-ctx.Done():
			return
		}
	}
}

func logArticles(ctx context.Context, cfg *config.Config) error {
	dsn := db.FormatDSN(cfg)

	listener := pq.NewListener(dsn, time.Second*5, time.Minute, func(event pq.ListenerEventType, err error) {
		switch event {
		case pq.ListenerEventConnected:
			log.Println("listener connected")
		case pq.ListenerEventDisconnected:
			log.Println("listener disconnected")
		case pq.ListenerEventReconnected:
			log.Println("listener reconnected")
		case pq.ListenerEventConnectionAttemptFailed:
			log.Println("listener connection attempt failed")
		}
	})

	const chann = "events"
	if err := listener.Listen(chann); err != nil {
		return err
	}
	if err := listener.Ping(); err != nil {
		return err
	}

	for {
		select {
		case n := <-listener.Notify:
			log.Printf("notification received: %s", n.Extra)
		}
	}

	return nil
}
