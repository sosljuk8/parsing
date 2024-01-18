package main

import (
	"github.com/spf13/cobra"
	"log/slog"
	"parsing/orm"
	"parsing/parse"
	"parsing/parse/plugin"
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

var crawlerCmd = &cobra.Command{
	Use:   "crawler",
	Short: "Run crawler jobs",
	Run: func(cmd *cobra.Command, args []string) {

		// create database ORM client
		c := orm.NewClient()
		defer c.Close()

		// create high-level database store
		s := orm.NewStore(c)

		// initialize jobs for crawler
		jobs := plugin.NewJobs()
		jobs.Add(plugin.NewHydacJob())

		// create crawler instance
		p := parse.NewCrawler(s, jobs)

		// run crawler until stopped
		p.Run()
	},
}
