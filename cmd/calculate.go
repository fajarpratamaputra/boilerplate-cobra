/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"log"
	"top-ranking-worker/config"
	"top-ranking-worker/infra"
	"top-ranking-worker/lineup"
	"top-ranking-worker/writer"
)

// calculateCmd represents the calculate command
var calculateCmd = &cobra.Command{
	Use:   "calculate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := config.Load(); err != nil {
			log.Fatal(err)
		}

		ctx := context.Background()

		menu := cmd.Flag("menu").Value.String()

		wrt := writer.NewWriter()
		mongoDb, err := infra.NewMongoDatabase(ctx)
		if err != nil {
			log.Fatal(err)
		}

		if err = lineup.Calculate(ctx, wrt, mongoDb, menu); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(calculateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	calculateCmd.PersistentFlags().String("menu", "fyp", "Menu that need to be summarized")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// calculateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
