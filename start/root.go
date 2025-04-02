package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "web-cli",
	Short: "Read Atom feeds",
	Long: `Just a small CLI application.
        Read Atom feeds`,
}

var cmdLs = &cobra.Command{
	Use:   "ls",
	Short: "List news",
	Long:  `List first 5 news`,
	Run: func(cmd *cobra.Command,
		args []string) {
		inner()
	},
}

var cmdAbout = &cobra.Command{
	Use:   "about",
	Short: "About this program",
	Run: func(cmd *cobra.Command,
		args []string) {
		fmt.Println(" ** This is a small CLI application to read Atom feeds.")
	},
}

func Exec() {
	rootCmd.AddCommand(cmdLs)
	rootCmd.AddCommand(cmdAbout)
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
