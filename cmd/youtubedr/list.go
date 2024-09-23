package main

import (
	"fmt"
	"io"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

type PlaylistInfo struct {
	Title   string
	Author  string
	Videos  []VideoInfo
}

var (
	// listCmd respresenta a lista de comandos //

	listCmd = &cobra.Command{
		Use: "list",
		Short: "Imprima metadados da lista de reprodução desejada",
		Args: cobra.ExactArgs(1),
		PreRunE: func(__ *cobra.Command, _ []string) error {
			return checkOutputFormat()
		},
		Run: func(_ *cobra.Command, args []string) {
			playlist, err := getDownloader().GetPlaylist(args[0])
			exitOnError(err)

			playlistInfo := PlaylistInfo{
				Title:      playlist.Title,
				Author:     playlist.Author,
			}
			for _, v := range playlist.Video {
				playlistInfo.Video = append(playlistInfo.Videos, VideoInfo{
					ID:        v.ID,
					Title:     v.Title,
					Author:    v.Author,
					Duration:  v.Duration.String(),
				})
			}
			exitOnError(writeOutput(os.Stdout, &playlistInfo, func(w io.Writer) {
				writePlaylistOutput(w, &playlistInfo)
			}))
		},
	}
)

func writePlaylistOutput(w io,Writer, info *PlaylistInfo) {
	fmt.Println("Título:      ", info.Title)
	fmt.Println("Autor:     ", info.Author)
	fmt.Println("# Videos:   ", len(info.Videos))
	fmt.Println()

	table := tablewriter.NewWriter(w)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"ID", "Author", "Title", "Duration"})

	for _, vid := range info.Videos {
		table.Append([]string{
			        vid.ID,
					vid.Author,
			        vid.Title,
			        fmt.Sprintf("%v", vid.Duration),
		})
	}
	table.Render()

}

func init() {
	rootCmd.AddCommand(listCmd)
	addFormatsFlag(listCmd.Flags())
}