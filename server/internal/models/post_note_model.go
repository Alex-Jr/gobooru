package models

import (
	"encoding/json"
	"time"
)

type PostNote struct {
	ID        int       `db:"id" json:"id"`
	PostID    int       `db:"post_id" json:"post_id"`
	Body      string    `db:"body" json:"body"`
	X         int       `db:"x" json:"x"`
	Y         int       `db:"y" json:"y"`
	Width     int       `db:"width" json:"width"`
	Height    int       `db:"height" json:"height"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type PostNoteList []PostNote

func (list *PostNoteList) Scan(src interface{}) error {
	if data, ok := src.([]byte); ok && len(data) > 0 {
		if err := json.Unmarshal(data, list); err != nil {
			return err
		}
	}
	return nil
}
