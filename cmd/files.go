package cmd

import (
	"fmt"

	"github.com/jonavdm/aoc-tool/gen"
	"github.com/spf13/cobra"
)

// filesCmd represents the files command
var filesCmd = &cobra.Command{
	Use:   "files",
	Short: "Clone the day template over",
	Run: func(cmd *cobra.Command, args []string) {
		year, err := cmd.Flags().GetInt("year")
		cobra.CheckErr(err)
		day, err := cmd.Flags().GetInt("day")
		cobra.CheckErr(err)

		cobra.CheckErr(gen.GenerateTemplates(year, day))
		fmt.Println("Files have been generated!")
		fmt.Println("\nNow just add the following entry to main.go:")
		fmt.Printf("{%d, day%02d.Run, \"day%02d\"}\n", day, day, day)
	},
}

func init() {
	dayCmd.AddCommand(filesCmd)
}
