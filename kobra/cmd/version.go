package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version [no options!]",
	Short: "Prints cobra version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Merhaba yalan dunya iste versiyon 29")
	},
}
