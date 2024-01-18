package main

import (
	"github.com/spf13/cobra"
	"log/slog"
	"parsing/orm"
)

// RootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "home",
	Short: "A brief description of your application",
	Long:  `A longer description that spans multiple lines`,
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate database schema",
	Long:  `Note this command must not be run in parallel with processing commands`,
	Run: func(cmd *cobra.Command, args []string) {

		err := orm.Migrate()
		if err != nil {
			slog.Error("Migrate", err)
		}
	},
}
