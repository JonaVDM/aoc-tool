package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

// dayCmd represents the day command
var dayCmd = &cobra.Command{
	Use:   "day",
	Short: "Init a day of AOC",
	Run: func(cmd *cobra.Command, args []string) {
		filesCmd.Run(cmd, args)
		inputCmd.Run(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(dayCmd)

	today := time.Now()

	dayCmd.PersistentFlags().IntP("year", "y", today.Year(), "The year")
	dayCmd.PersistentFlags().IntP("day", "d", today.Day(), "The day")

	dayCmd.AddCommand(filesCmd)
	dayCmd.AddCommand(inputCmd)
}
