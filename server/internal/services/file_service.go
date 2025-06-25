package services

import (
	"crypto/md5"
	"fmt"
	"gobooru/internal/ffmpeg"
	"gobooru/internal/models"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

type FileServiceConfig struct {
	FFMPEGModule ffmpeg.FFMPEGModule
	BASE_PATH    string
}

type FileService interface {
	// DeleteDirectory(filePath string) error
	// Exists(path string) (bool, error)
	HandleUpload(fileHeader *multipart.FileHeader) (models.File, error)
	// HandleFileUpload(fileHeader *multipart.FileHeader) (*models.File, error)
	// HandleTempFileUpload(fileHeader *multipart.FileHeader) (*models.File, error)
	// MarkToDelete(filePath string) error
}

type fileService struct {
	ffmpegModule ffmpeg.FFMPEGModule
	basePath     string
}

func NewFileService(c FileServiceConfig) FileService {
	return fileService{
		ffmpegModule: c.FFMPEGModule,
		basePath:     c.BASE_PATH,
	}
}

func (f fileService) HandleUpload(fileHeader *multipart.FileHeader) (models.File, error) {
	src, err := fileHeader.Open()
	if err != nil {
		return models.File{}, fmt.Errorf("fileHeader.Open: %w", err)
	}
	defer src.Close()

	MD5String, err := f.GenerateMD5(src)
	if err != nil {
		return models.File{}, fmt.Errorf("f.GenerateMD5: %w", err)
	}

	ext := strings.Replace(filepath.Ext(fileHeader.Filename), ".", "", 1)

	if ext == "" {
		return models.File{}, fmt.Errorf("ext is empty")
	}

	folderPath := fmt.Sprintf("%s/posts/%s/%s/%s", f.basePath, MD5String[:2], MD5String[2:4], MD5String)
	filePath := fmt.Sprintf("%s/%s.%s", folderPath, "original", ext)
	thumbPath := fmt.Sprintf("%s/%s.%s", folderPath, "250", "webp")

	src.Seek(0, io.SeekStart)

	err = f.SaveFile(filePath, src)
	if err != nil {
		return models.File{}, fmt.Errorf("f.SaveFile: %w", err)
	}

	err = f.GenerateThumb(filePath, thumbPath)
	if err != nil {
		log.Println(fmt.Errorf("f.GenerateThumb: %w", err))
		thumbPath = fmt.Sprintf("%s/%s", folderPath, "error.webp")
	}

	fileMetadata, err := f.ffmpegModule.ExtractMetadata(filePath)
	if err != nil {
		return models.File{}, fmt.Errorf("f.ffmpegModule.ExtractMetadata: %w", err)
	}

	file := models.File{
		MD5:              MD5String,
		FilePath:         strings.Replace(filePath, f.basePath, "", 1),
		FileExt:          ext,
		ThumbPath:        strings.Replace(thumbPath, f.basePath, "", 1),
		FileSize:         int(fileHeader.Size),
		FileOriginalName: fileHeader.Filename,
		Width:            fileMetadata.Width,
		Height:           fileMetadata.Height,
		Duration:         fileMetadata.Duration,
	}

	return file, nil
}

func (f fileService) DeleteDirectory(filePath string) error {
	err := os.RemoveAll(f.basePath + filePath)

	if err != nil {
		return fmt.Errorf("os.RemoveAll: %w", err)
	}

	return nil
}

func (f fileService) Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, fmt.Errorf("os.Stat: %w", err)
}

func (f fileService) SaveFile(dstPath string, src io.Reader) error {
	folderPath := filepath.Dir(dstPath)

	isExist, err := f.Exists(folderPath)
	if err != nil {
		return fmt.Errorf("f.Exists: %w", err)
	}

	if isExist {
		return nil
	}

	err = os.MkdirAll(folderPath, 0750)
	if err != nil {
		return fmt.Errorf("os.MkdirAll: %w", err)
	}

	dst, err := os.Create(dstPath)
	if err != nil {
		return fmt.Errorf("os.Create: %w", err)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return fmt.Errorf("io.Copy: %w", err)
	}

	return nil
}

func (f fileService) GenerateMD5(src io.Reader) (string, error) {
	dst := md5.New()

	_, err := io.Copy(dst, src)
	if err != nil {
		return "", fmt.Errorf("io.Copy: %w", err)
	}

	return fmt.Sprintf("%x", dst.Sum(nil)), nil
}

func (f fileService) GenerateThumb(src string, dst string) error {
	var err error

	ext := strings.Replace(filepath.Ext(src), ".", "", 1)

	exist, err := f.Exists(dst)
	if err != nil {
		return fmt.Errorf("f.Exists: %w", err)
	}

	if exist {
		return nil
	}

	switch ext {
	case
		"jpeg",
		"jpg",
		"png",
		"webp":
		err = f.ffmpegModule.GenerateImageThumb(src, dst)
	case
		"gif":
		err = f.ffmpegModule.GenerateGifThumb(src, dst)
	case
		"mov",
		"mp4",
		"webm":
		err = f.ffmpegModule.GenerateVideoThumb(src, dst)
	case "swf":
		return nil
	default:
		return fmt.Errorf("%s is not supported", ext)
	}

	if err != nil {
		return fmt.Errorf("ffmpegModule.Generate*Thumb: %w", err)
	}

	return nil
}
