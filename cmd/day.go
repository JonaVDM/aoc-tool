package cmd

import (
	"fmt"
	"os"
	"text/template"
	"time"

	"github.com/jonavdm/aoc-tool/tpl"
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
		dayStr := fmt.Sprintf("day%02d", day)
		pwd, err := os.Getwd()
		cobra.CheckErr(err)

		data := tpl.DayTempl{
			Day:  day,
			Year: year,
		}

		// Create the folder
		cobra.CheckErr(createDir(dayStr))

		// Create the day file
		mainFile, err := os.Create(fmt.Sprintf("%s/%s/%s.go", pwd, dayStr, dayStr))
		cobra.CheckErr(err)
		defer mainFile.Close()

		dayTemplate := template.Must(template.New("day").Parse(string(tpl.DayTemplate())))
		err = dayTemplate.Execute(mainFile, data)
		cobra.CheckErr(err)

		// Create the test file
		testFile, err := os.Create(fmt.Sprintf("%s/%s/%s_test.go", pwd, dayStr, dayStr))
		cobra.CheckErr(err)
		defer testFile.Close()

		testTemplate := template.Must(template.New("test").Parse(string(tpl.DayTestTemplate())))
		err = testTemplate.Execute(testFile, data)
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(dayCmd)

	today := time.Now()

	dayCmd.Flags().IntP("year", "y", today.Year(), "The year")
	dayCmd.Flags().IntP("day", "d", today.Day(), "The day")
}
