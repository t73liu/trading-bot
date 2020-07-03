package traderdb

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type NewsSource struct {
	Id          string
	Name        string
	Description string
}

const userNewsSourcesQuery = `
SELECT ns.id, ns.name, ns.description FROM user_news_sources uns
INNER JOIN news_sources ns ON ns.id = uns.news_source_id
WHERE uns.user_id = $1
`

func GetNewsSourcesByUserId(db *pgxpool.Pool, userId int) (newsSources []NewsSource, err error) {
	rows, err := db.Query(context.Background(), userNewsSourcesQuery, userId)
	if err != nil {
		return newsSources, err
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var name string
		var description string
		if err = rows.Scan(&id, &name, &description); err != nil {
			return newsSources, err
		}
		newsSources = append(newsSources, NewsSource{
			Id:          id,
			Name:        name,
			Description: description,
		})
	}

	if rows.Err() != nil {
		return newsSources, rows.Err()
	}
	return newsSources, err
}

func GetNewsSourceIdsByUserId(db *pgxpool.Pool, userId int) (newsSourceIds []string, err error) {
	newsSources, err := GetNewsSourcesByUserId(db, userId)
	if err != nil {
		return newsSourceIds, err
	}
	for _, newsSource := range newsSources {
		newsSourceIds = append(newsSourceIds, newsSource.Id)
	}
	return newsSourceIds, nil
}
