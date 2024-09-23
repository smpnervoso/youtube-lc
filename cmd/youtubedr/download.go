package main

import (
	    "context"
		"fmt"
		"log"
		"os/exec"
		"strings"

		"github.com/spf13/cobra"
)

 // downloadCmd representa o comando de download // 
 var downloadCmd = &cobra.Command {
	Use:           "download",
	Short:         "Baixar um vídeo do YouTube",
	Example:       `youtubedr -o "Campaign Diary".mp4 https://www.youtube.com/watch\?v\=XbNghLqsVwU`,
	Args:          cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		exitOnError(download(args[0]))
	},
 }

 var (
	    ffmpegCheck error
		outputFile  string
		outputDir   string
 )

 func init() {
	     rootCmd.AddCommand(downloadCmd)

		 downloadCmd.Flags().StringVarP(&outputFile, "filename", "o", "", "O arquivo de saída, o padrão, é gerado pelo título do vídeo.")
		 downloadCmd.Flags().StringVarP(&outputDir, "directory", "d", ".", "O diretório de saída")
		 addVideoSelectionFlags(downloadCmd.Flags())
}

func download(id string) error {
	     video, format, err := getVideoWithFormat(id)
		 if err != nil {
			return err

		 }

		 log.Println("baixar no diretório", outputDir)

		 if strings.HasPrefix(outputQuality, "hd"){
			if err := checkFFMPEG(); err != nil {
				return err
		}
		 return downloader.DownloadComposite(context,context.Background(), outputFile, video, outputQuality, mimetype, language)
	}
	
	return downloader.Download(context.Background(), video, format, outputFile)

}

func checkFFMPEG() error {
	fmt.Println("check ffmpeg is installed....")
	if err := exec.Command("ffmpeg", "-version").Run(); err != nil {
		ffmpegCheck = fmt.Errorf("Verifique se o ffmpegCheck foi instalado corretamente")
	}

	return ffmpegCheck
}