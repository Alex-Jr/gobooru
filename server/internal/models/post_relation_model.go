package models

import (
	"encoding/json"
	"time"
)

type PostRelation struct {
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	OtherPostID int       `db:"other_post_id" json:"other_post_id"`
	OtherPost   Post      `db:"other_post" json:"other_post"`
	PostID      int       `db:"post_id" json:"post_id"`
	Similarity  int       `db:"similarity" json:"similarity"`
	Type        string    `db:"type" json:"type"`
}

type PostRelationList []PostRelation

func (list *PostRelationList) Scan(src interface{}) error {
	if data, ok := src.([]byte); ok && len(data) > 0 {
		if err := json.Unmarshal(data, list); err != nil {
			return err
		}
	}
	return nil
}
