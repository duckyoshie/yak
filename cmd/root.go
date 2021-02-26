package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// root command
var rootCmd = &cobra.Command{
	Use:   "yak",
	Short: "A little CLI",
	Run: func(cmd *cobra.Command, args []string) {
		yakLogo := `
		_____  _____
	        \   \  /   /
		 \   \/   /             ___   ____
		  \      /             |   | /   /
		   |    |   _______    |   |/   /
		   |    |  /       \   |       /
		   |    | |   |  |  |  |   |\  \
		   |____|  \_____|__|  |___| \__\
		`
		fmt.Println(yakLogo)
		fmt.Println("Welcome to the Yak CLI. This is a personal project just for exploring Go.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
