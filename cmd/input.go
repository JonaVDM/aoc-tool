package cmd

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jonavdm/aoc-tool/aoc"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// inputCmd represents the input command
var inputCmd = &cobra.Command{
	Use:   "input",
	Short: "Download the input file for a day, defaults to this day and this year",
	Run: func(cmd *cobra.Command, args []string) {
		year, err := cmd.Flags().GetInt("year")
		cobra.CheckErr(err)
		day, err := cmd.Flags().GetInt("day")
		cobra.CheckErr(err)
		token := viper.GetString("token")

		if token == "" {
			fmt.Println("Token is not set! Please set it with the setToken command")
			return
		}

		inp, err := aoc.DownloadInput(year, day, token)
		cobra.CheckErr(err)

		createDir("inputs")
		file := fmt.Sprintf("inputs/%02d", day)
		err = os.WriteFile(file, inp, 0644)
		cobra.CheckErr(err)

		fmt.Println("Downloaded the input file!")
	},
}

func init() {
	rootCmd.AddCommand(inputCmd)

	today := time.Now()

	inputCmd.Flags().IntP("year", "y", today.Year(), "The year")
	inputCmd.Flags().IntP("day", "d", today.Day(), "The year")
}

func createDir(name string) error {
	_, err := os.Stat(name)

	if errors.Is(err, os.ErrNotExist) {
		return os.Mkdir(name, os.ModePerm)
	}

	return nil
}
