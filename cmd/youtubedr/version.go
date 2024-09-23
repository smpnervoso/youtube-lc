package main

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	version string
	commir string
	date string
)

// versionCmd significa a versão do do comando

var versionCmd = &cobra.Command{
	Use:  "version",
	Short: "Prints version informação",
	Run: func(cobra.Command, _ []string) {
		fmt.Println("Version:    ", version)
		fmt.Println("Commit:     ", commit)
		fmt.Println("Date:       ", date)
		fmt.Println("Go Version: ", runtime.Version())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}