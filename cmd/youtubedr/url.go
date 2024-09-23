package main

import (
	"fmt"
	
	"github.com/spf13/cobra"
)

// urlCmd significa a Url do comando

var urlCmd = &cobra.Command{
	Use:     "url",
	Short:   "Única saída de stream-url para o vídeo",
	Args:    cobra.ExactArgs(1),
	Run:     func(_ *cobra.Command, args[]string) {
		video, format, err := getVideoWithFormat(args[0])
		exitOnError(err)

		url, err := downloaderr.GetStreamURL(video, format)
		exitOnError(err)

		fmt.Println(url)
	},
}

func init() {
	addVideoSelectionFlags(urlCmd.Flags())
	rootCmd.AddCommand(urlCmd)
}