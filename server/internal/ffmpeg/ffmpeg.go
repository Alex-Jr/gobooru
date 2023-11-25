package ffmpeg

import (
	"bytes"
	"fmt"
	"os/exec"
)

type FFMPEGModuleConfig struct {
}

type FFMPEGModule interface {
	GenerateGifThumb(src string, dst string) error
	GenerateImageThumb(src string, dst string) error
	GenerateVideoThumb(src string, dst string) error
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
