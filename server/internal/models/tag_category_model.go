package models

import "time"

type TagCategory struct {
	ID          string    `db:"id" json:"id"`
	Description string    `db:"description" json:"description"`
	Color       string    `db:"color" json:"color"`
	TagCount    int       `db:"tag_count" json:"tag_count"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
