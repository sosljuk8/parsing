package main

import (
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "home",
	Short: "A brief description of your application",
	Long:  `A longer description that spans multiple lines`,
}
