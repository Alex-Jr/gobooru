package models

type TagImplication struct {
	TagID         string `db:"tag_id"`
	ImplicationID string `db:"implication_id"`
}
