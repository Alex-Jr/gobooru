package models

import (
	"encoding/json"
	"time"

	"github.com/lib/pq"
)

type Pool struct {
	CreatedAt   time.Time      `db:"created_at" json:"created_at"`
	Custom      pq.StringArray `db:"custom" json:"custom"`
	Description string         `db:"description" json:"description"`
	ID          int            `db:"id" json:"id"`
	Name        string         `db:"name" json:"name"`
	PostCount   int            `db:"post_count" json:"post_count"`
	Posts       PostList       `db:"posts" json:"posts"`
	UpdatedAt   time.Time      `db:"updated_at" json:"updated_at"`
}

type PoolList []Pool

func (list *PoolList) Scan(src interface{}) error {
	if data, ok := src.([]byte); ok && len(data) > 0 {
		if err := json.Unmarshal(data, list); err != nil {
			return err
		}
	}
	return nil
}

type PoolPost struct {
	CreatedAt time.Time `db:"created_at"`
	Pool      Pool      `db:"pool"`
	PoolID    int       `db:"pool_id"`
	Position  int       `db:"position"`
	Post      Post      `db:"post"`
	PostID    int       `db:"post_id"`
	UpdatedAt time.Time `db:"updated_at"`
}
