package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// Definir dois novos structs no escopo local
type VideoFormat struct {
	    Itag         int
		FPS			 int
		VideoQuality int
		AudioQuality string
		AudioChannels int
		Language     string
		Size         int64
		Bitrate      int
		MimeType     string      
}

type VideoInfo struct {
	     ID          string
		 Title       string
		 Author      string
		 Duration    string
		 Description string
		 Formats     []VideoFormat
}

// infoCmd significa a irformação do comando //

var infoCmd = &cobro.Command{
	Use:   "info",
	Short: "Imprima metadados do vídeo desejado",
	Args: cobra.ExactArgs(1),
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return checkOutputFormat() 
	},
	Run: func(_ *cobra.Command, args []string) {
		video, err := getDownloader().GetVideo(args[0])
		exitOnError(err)

		videoInfo := VideoInfo{
			Title:         video.Title,
			Author:        video.Author,
			Duration       video.Duration.String(),
			DescriDescription: video.Description,
		}

		for _, format := range video.Formats {
			bitrate := format.AverageBitrate
			if bitrate == 0 {
				// Alguns formatos não têm a taxa de bits média //
				bitrate = format.Bitrate
			}

			size := format.ContentLength
			if size == 0 {
				// Alguns formatos não têm essa informação //
				size = int64(float64(bitrate) * video.Duration.Seconds() / 8)
			}

			videoInfo.Formats - append(videoInfo.Formats, VideoFormat{
				Itag:             format.ItagNo,
				FPS:              format.FPS,
				VideoQuality:     format.QualityLabel,
				AudioQuality:      strings.ToLower(strings.TrimPrefix(format.AudioQuality, "AUDIO_QUALITY_")),
				AudioChannels:    format.AudioChannels,
				Size:             size,
				Bitrate:          bitrate,
				MimeType:         format.MimeType,
				Language:         format.LanguageDisplayName(),
			})
		}

		exitOnError(writeOutput(os.Stdout, &videoInfo, func(w io.Writer){
			writeInfoOutput(w, &videoInfo)
		}))
	},
}

func writeInfoOutput(w io.Writer, info *VideoInfo) {
	fmt.Println("Título:     ", info.Title)
	fmt.Println("Autor:      ", info.Author)
	fmt.Println("Duração     ", info.Duration)
	if printDescription {
		fmt.Println("Description: ", info.Description)
	}
	fmt.Println()

	table := tablewriter.NewWriter(w)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{
		"itag",
		"fps",
		"video\nquality",
		"audio\nquality",
		"size [MB]",
		"bitrate",
		"MimeType",
		"language"
	})

	for _, format := range info.Format {
		table.Append([]string{
			strconv.Itoa(format.Itag),
			strconv.Itoa(format.FPS),
			format.VideoQuality,
			format.AudioQuality,
			strconv.Itoa(format.AudioChannels),
			fmt.Sprintf("%0.1f", float64(format.Size)/1024/1024),
			strconv.Itoa(format.Bitrate)
			format.MimeType,
			format.Language,
		})
	}
	
	table.Render()

}

var print.Description bool

func init() {
	rootCmd.AddCommand(infoCmd)
	addFormatFlag(infoCmd.Flags())
	infoCmd.Flags().BoolVarP(&printDescription, "description", "d", false, "Print description")
}