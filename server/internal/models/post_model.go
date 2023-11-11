package models

import (
	"encoding/json"
	"time"
)

type Post struct {
	CreatedAt   time.Time `db:"created_at"`
	Description string    `db:"description"`
	ID          int       `db:"id"`
	Pools       PoolList  `db:"pools"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type PostList []Post

func (list *PostList) Scan(src interface{}) error {
	if data, ok := src.([]byte); ok && len(data) > 0 {
		if err := json.Unmarshal(data, list); err != nil {
			return err
		}
	}
	return nil
}
