package writer

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func StartWriter(ctx context.Context, pool *pgxpool.Pool) {
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	log.Println("starting writer...")

	for {
		select {
		case <-ticker.C:
			if err := upsertArticle(ctx, pool); err != nil {
				log.Printf("error while upserting article: %v\n", err)
			}
			log.Printf("article upserted at %s", time.Now())
		case <-ctx.Done():
			return
		}
	}
}

func upsertArticle(ctx context.Context, pool *pgxpool.Pool) error {
	title := "Example Article"
	content := fmt.Sprintf("Content updated at %s", time.Now().Format(time.RFC3339))

	query := `
		INSERT INTO articles(id, title, content, created_at, updated_at)
		VALUES (1, $1, $2, NOW(), NOW())
		ON CONFLICT (id)
		DO UPDATE SET 
			content = EXCLUDED.content, 
			updated_at = NOW();
	`

	_, err := pool.Exec(ctx, query, title, content)
	return err
}
