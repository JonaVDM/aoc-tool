package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// setTokenCmd represents the setToken command
var setTokenCmd = &cobra.Command{
	Use:   "setToken",
	Short: "Set the token from AOC",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("token", args[0])
		err := viper.WriteConfig()
		cobra.CheckErr(err)

		fmt.Println("Updated the config file!")
	},
}

func init() {
	rootCmd.AddCommand(setTokenCmd)
}
