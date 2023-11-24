package models

import (
	"encoding/json"
	"time"

	"github.com/lib/pq"
)

type Pool struct {
	ID          int            `db:"id" json:"id"`
	Name        string         `db:"name" json:"name"`
	PostCount   int            `db:"post_count" json:"post_count"`
	Description string         `db:"description" json:"description"`
	Custom      pq.StringArray `db:"custom" json:"custom"`
	CreatedAt   time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at" json:"updated_at"`
	Posts       PostList       `db:"posts" json:"posts"`
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
	PoolID    int       `db:"pool_id"`
	Position  int       `db:"position"`
	PostID    int       `db:"post_id"`
	UpdatedAt time.Time `db:"updated_at"`
}
