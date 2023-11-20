package models

import (
	"encoding/json"
	"time"
)

type Tag struct {
	ID          string    `db:"id" json:"id"`
	Description string    `db:"description" json:"description"`
	PostCount   int       `db:"post_count" json:"post_count"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type TagList []Tag

func (list *TagList) Scan(src interface{}) error {
	if data, ok := src.([]byte); ok && len(data) > 0 {
		if err := json.Unmarshal(data, list); err != nil {
			return err
		}
	}
	return nil
}
