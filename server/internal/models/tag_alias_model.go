package models

type TagAlias struct {
	TagID string `db:"tag_id"`
	Alias string `db:"alias"`
}
