package ffmpeg

import (
	"bytes"
	"fmt"
	"gobooru/internal/models"
	"os/exec"
	"regexp"
	"strconv"
)

type FFMPEGModuleConfig struct {
}

type FFMPEGModule interface {
	GenerateGifThumb(src string, dst string) error
	GenerateImageThumb(src string, dst string) error
	GenerateVideoThumb(src string, dst string) error
	ExtractMetadata(src string) (models.FileMetadata, error)
}

type ffmpegModule struct {
}

func NewFfmpegModule() FFMPEGModule {
	return ffmpegModule{}
}

func (f ffmpegModule) GenerateGifThumb(src string, dst string) error {
	cmd := exec.Command(
		"ffmpeg",
		"-i", src,
		"-vf", "scale=250:250:force_original_aspect_ratio=decrease",
		"-vframes", "1",
		dst,
	)
	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		return fmt.Errorf("cmd.Run: %w", err)
	}

	return nil
}

func (f ffmpegModule) GenerateImageThumb(src string, dst string) error {
	cmd := exec.Command(
		"ffmpeg",
		"-i", src,
		"-vf", "scale=250:250:force_original_aspect_ratio=decrease",
		dst,
	)
	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		return fmt.Errorf("cmd.Run: %w", err)
	}

	return nil
}

func (f ffmpegModule) GenerateVideoThumb(src string, dst string) error {
	cmd := exec.Command(
		"ffmpeg",
		"-i", src,
		"-vf", "thumbnail=1000,scale=250:250:force_original_aspect_ratio=decrease",
		"-frames:v", "1", dst,
	)

	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		return fmt.Errorf("cmd.Run: %w", err)
	}

	return nil
}

var widthRe = regexp.MustCompile(`width=(\d+)`)
var heightRe = regexp.MustCompile(`height=(\d+)`)
var durationRe = regexp.MustCompile(`duration=(\d+\.\d+)`)

func (f ffmpegModule) ExtractMetadata(src string) (models.FileMetadata, error) {
	cmd := exec.Command(
		"ffprobe",
		"-v", "error",
		"-select_streams", "v:0",
		"-show_entries", "stream=width,height,duration:format=duration",
		"-of", "default=nw=1",
		src,
	)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return models.FileMetadata{}, fmt.Errorf("cmd.Run: %w %s", err, stderr.String())
	}

	outBytes := out.Bytes()

	widthMatch := widthRe.FindSubmatch(outBytes)
	heightMatch := heightRe.FindSubmatch(outBytes)
	durationMatch := durationRe.FindSubmatch(outBytes)

	width := 0
	height := 0
	duration := 0.0

	if len(widthMatch) == 2 {
		width, err = strconv.Atoi(string(widthMatch[1]))
		if err != nil {
			return models.FileMetadata{}, fmt.Errorf("strconv.Atoi: %w", err)
		}
	}

	if len(heightMatch) == 2 {
		height, err = strconv.Atoi(string(heightMatch[1]))
		if err != nil {
			return models.FileMetadata{}, fmt.Errorf("strconv.Atoi: %w", err)
		}
	}

	if len(durationMatch) == 2 {
		duration, err = strconv.ParseFloat(string(durationMatch[1]), 32)
		if err != nil {
			return models.FileMetadata{}, fmt.Errorf("strconv.ParseFloat: %w", err)
		}
	}

	return models.FileMetadata{
		Width:    width,
		Height:   height,
		Duration: int(duration * 100),
	}, nil
}
