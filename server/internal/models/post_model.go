package models

import (
	"encoding/json"
	"time"

	"github.com/lib/pq"
)

type Post struct {
	ID               int              `db:"id" json:"id"`
	Rating           string           `db:"rating" json:"rating"`
	Description      string           `db:"description" json:"description"`
	TagIDs           pq.StringArray   `db:"tag_ids" json:"tag_ids"`
	TagCount         int              `db:"tag_count" json:"tag_count"`
	PoolCount        int              `db:"pool_count" json:"pool_count"`
	MD5              string           `db:"md5" json:"md5"`
	FileExt          string           `db:"file_ext" json:"file_ext"`
	FileSize         int              `db:"file_size" json:"file_size"`
	FilePath         string           `db:"file_path" json:"file_path"`
	FileOriginalName string           `db:"file_original_name" json:"file_original_name"`
	ThumbPath        string           `db:"thumb_path" json:"thumb_path"`
	Sources          pq.StringArray   `db:"sources" json:"sources"`
	Custom           pq.StringArray   `db:"custom" json:"custom"`
	CreatedAt        time.Time        `db:"created_at" json:"created_at"`
	UpdatedAt        time.Time        `db:"updated_at" json:"updated_at"`
	Pools            PoolList         `db:"pools" json:"pools"`
	Tags             TagList          `db:"tags" json:"tags"`
	Relations        PostRelationList `db:"relations" json:"relations"`
	// TODO: Add source
	// Source      []string         `db:"source" json:"source"`
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

type PostTag struct {
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	PostID    int       `db:"post_id" json:"post_id"`
	TagID     string    `db:"tag_id" json:"tag_id"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
