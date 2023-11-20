package models

import (
	"encoding/json"
	"time"

	"github.com/lib/pq"
)

type Post struct {
	ID          int            `db:"id" json:"id"`
	Rating      string         `db:"rating" json:"rating"`
	Description string         `db:"description" json:"description"`
	TagIDs      pq.StringArray `db:"tag_ids" json:"tag_ids"`
	TagCount    int            `db:"tag_count" json:"tag_count"`
	PoolCount   int            `db:"pool_count" json:"pool_count"`
	CreatedAt   time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at" json:"updated_at"`
	Pools       PoolList       `db:"pools" json:"pools"`
	Tags        TagList        `db:"tags" json:"tags"`
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
