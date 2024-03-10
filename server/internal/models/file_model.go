package models

type File struct {
	FileExt          string
	FilePath         string
	FileSize         int
	FileOriginalName string
	MD5              string
	ThumbPath        string
	Width            int
	Height           int
	Duration         int
}

type FileMetadata struct {
	Width    int
	Height   int
	Duration int
}
