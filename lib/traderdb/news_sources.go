package traderdb

import (
	"context"
)

type NewsSource struct {
	ID          string
	Name        string
	Description string
}

const userNewsSourcesQuery = `
SELECT ns.id, ns.name, ns.description FROM user_news_sources uns
INNER JOIN news_sources ns ON ns.id = uns.news_source_id
WHERE uns.user_id = $1
`

func GetNewsSourcesWithUserID(db PGConnection, userID int) (newsSources []NewsSource, err error) {
	rows, err := db.Query(context.Background(), userNewsSourcesQuery, userID)
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
			ID:          id,
			Name:        name,
			Description: description,
		})
	}

	if rows.Err() != nil {
		return newsSources, rows.Err()
	}
	return newsSources, err
}

func GetNewsSourceIDsWithUserID(db PGConnection, userID int) (newsSourceIDs []string, err error) {
	newsSources, err := GetNewsSourcesWithUserID(db, userID)
	if err != nil {
		return newsSourceIDs, err
	}
	for _, newsSource := range newsSources {
		newsSourceIDs = append(newsSourceIDs, newsSource.ID)
	}
	return newsSourceIDs, nil
}
