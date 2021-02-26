package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// initialise the function
func init() {
	rootCmd.AddCommand(echoCmd)
}

var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "Echos text",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Requires arguement [text]")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var txt string

		// loop through string
		for i := 0; i < len(args); i++ {
			txt += args[i] + " " //a space between words
		}
		fmt.Println(txt)
	},
}
