package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var contextCmd = &cobra.Command{
	Use:   "context user TODO",
	Short: "bzctl context",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("%v", args)
		fmt.Println("bzctl context")
	},
}

// use stringi cobranin davranisini etkiliyor
var userCmd = &cobra.Command{
	Use:   "user register login TODO ",
	Short: "bzctl context user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("%v", args)
		fmt.Println("bzctl context user")
	},
}
