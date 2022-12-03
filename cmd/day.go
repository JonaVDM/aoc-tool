package cmd

import (
	"time"

	"github.com/jonavdm/aoc-tool/gen"
	"github.com/spf13/cobra"
)

// dayCmd represents the day command
var dayCmd = &cobra.Command{
	Use:   "day",
	Short: "Init a day of AOC",
	Run: func(cmd *cobra.Command, args []string) {
		year, err := cmd.Flags().GetInt("year")
		cobra.CheckErr(err)
		day, err := cmd.Flags().GetInt("day")
		cobra.CheckErr(err)

		cobra.CheckErr(gen.GenerateTemplates(year, day))
	},
}

func init() {
	rootCmd.AddCommand(dayCmd)

	today := time.Now()

	dayCmd.Flags().IntP("year", "y", today.Year(), "The year")
	dayCmd.Flags().IntP("day", "d", today.Day(), "The day")
}
