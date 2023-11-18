package cmd

import (
	"path/filepath"

	"github.com/POMPOSAN/dosu/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setCmd)

	setCmd.Flags().StringP("path", "p", "", "osu! directory path")
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Change dosu settigs !",
	Long: `Add or update your dosu settings.
You can for example change your osu path here.`,
	Run: set,
}

func set(cmd *cobra.Command, args []string) {

	s := &utils.Settings{}
	s.Load()

	p, err := cmd.Flags().GetString("path")
	if err != nil {
		panic(err)
	}

	if p != "" {
		pf := filepath.ToSlash(p)
		s.OsuPath = pf
	}

	err = s.Save()
	if err != nil {
		panic(err)
	}
}
