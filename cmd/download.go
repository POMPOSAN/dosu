package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/POMPOSAN/dosu/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(downloadCmd)

	downloadCmd.Flags().IntP("id", "i", 0, "Beatmap id")
}

var downloadCmd = &cobra.Command{
	Use:   "dl",
	Short: "Download the beatmap corresponding to the given id.",
	Long:  `Download the beatmap corresponding to the given id. Downloads from beatconnect.io`,

	Run: download,
}

func download(cmd *cobra.Command, args []string) {

	id, err := cmd.Flags().GetInt("id")
	if err != nil {
		panic(err)
	}

	s := &utils.Settings{}
	s.Load()

	if s.OsuPath == "" {
		fmt.Println("OSU PATH ERROR")
		fmt.Println("osu path: ", s.OsuPath)
	}

	client := &http.Client{}

	url := fmt.Sprintf("https://beatconnect.io/b/%d", id)

	fmt.Println(utils.DownloadStyle.Render("Downloading..."))

	res, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Println(utils.ExtractingStyle.Render("Extracting..."))

	filename := res.Header.Get("Content-Disposition")[23 : len(res.Header.Get("Content-Disposition"))-2]
	path := filepath.Join(filepath.FromSlash(s.OsuPath), "/Songs/", filename)

	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(utils.SuccessStyle.Render("Sucess !"))
}
