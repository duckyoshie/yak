package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func init() {
	// init
	rootCmd.AddCommand(fsCmd)

	fsCmd.Flags().StringVarP(&mkdir, "mkdir", "d", "", "Creates a new directory")
	fsCmd.Flags().StringVarP(&mkf, "mkf", "f", "", "Creates a new file.")
	fsCmd.Flags().StringVarP(&cd, "chdir", "c", "", "Changes directory.")
	fsCmd.Flags().StringVarP(&move, "move", "m", "", "Moves files to a different directory")
}

// flags
var mkdir string
var mkf string
var cd string
var move string
var copy string

var fsCmd = &cobra.Command{
	Use:   "fs",
	Short: "File system operators",
	Run: func(cmd *cobra.Command, args []string) {
		// checks use of flags and uses them to build functionality

		// mkdir
		mkdir, _ := cmd.Flags().GetString("mkdir")
		os.Mkdir(mkdir, os.ModePerm)

		//mkf
		mkf, _ := cmd.Flags().GetString("mkf")
		if mkf != "" {
			file, err := os.Create(mkf)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Successfully created " + file.Name())
			}
		} else {
			return
		}

		//cd
		cd, _ := cmd.Flags().GetString("cd")
		err := os.Chdir(filepath.Join(filepath.Dir(cd), cd))
		if err != nil {
			fmt.Println(err)
		}

	},
}
