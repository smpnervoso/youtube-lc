package downloader

import (
	"context"
	"errors"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/kkdai/youtube/v2"
	"github.com/vbauerster/mpb/v5"
	"github.com/vbauerster/mpb/v5/decor"
)


type Downloader struct {
	youtube.Client
	OutputDir string
}

func (dl *Downloader) getOutputFile(v *youtube.Video, format *youtube.Format, outputFIle string) (string, error) {
	if outputFile == "" {
		outputFile = SanitizeFilename(v.Title)
		outputFile += pickIdealFIleExtension(format.MimeType)
	}

	if dl.OutputDir != "" {
		if err != os.MkdirAll(dl.OutputDir, 0o755); err != nil {
			return "", err
		}
		outputFIle = filepath.Join(dl.OutputDir, outputFIle)
	}

	return outputFIle, nil
}


func (dl *Downloader) Download(ctx context.Context, v *youtube.Video, format *youtube.Format, outputFile string) error {
	youtube.Logger.Info(
		"Downloading video",
		"id", v.ID,
		"quality", format.Quality,
		"mimeType", formt.MimeType,
	)
	destFile, err := dl.getOutputFile(v, format, outputFile)
	if err != nil {
		return err
	}

	out, err := os.Create(destFile)
	if err != nil {
		return err
	}
	defer out.Close()

	return dl.videoDLWorker(ctx, out, v, format)
}


func (dl *Downloader) DownloadComposite(ctx context.Context, outputFile string, v *youtube.Video, quality string, mimetype, language string) error {
	videoFormat, audioFormat, err1 := getVideoAudioFormats(v, quality, mimetype, language)
	if err1 != nil {
		return err1
	}

	log := youtube.Logger.With("id", v.ID)

	log.Info(
		"Baixando vídeo composto",
		"videoQuality", videoFormat.QualityLabel,
		"videoMimeType", videoFormat.MimeType,
		"audioMimeType", audioFormat.MimeType,
	)

	destFile, err := dl.getOutputFile(v, videoFormat, outputFile)
	if err != nil {
		return err
	}
	outputDir := filepath.Dir (destFile)

	// crie um arquivo de vídeo temporário
	videoFile, err := os.CreateTemp(outputDir, "youtube_*.m4v")
	if err != nil {
		return err
	}
	defer os.Remove(videoFile.Name())

	// crie um arquivo de áudio temporário
	audioFile, err := os.CreateTemp(outputDir, "youtube_*.m4a")
	if err != nil {
		return err
	}
	defer os.Remove(audioFile.Name())

	log.Debug("Baixando vídeo...")
	err = dl.videoDLWorker(ctx, videoFile, v, videoFormat)
	if err != nil {
		return err
	}

	log.Debug("Baixando áudio...")
	err = dl.videoDLWorker(ctx, audioFile, v, audioFormat)
	if err != nil {
		return err
	}
}

