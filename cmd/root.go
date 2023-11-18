package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dosu",
	Short: "dosu is a command line osu beatmap downloader",
	Long: `Dosu is an osu beatmap donwloader running in the command line.
Just setup your osu path and dosu will download all your beatmaps directly in it !`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
