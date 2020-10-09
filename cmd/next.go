package cmd

/*
Copyright © 2019 Graham Plata <graham.plata@gmail.com>

*/

import (
	"github.com/grahamplata/sixers/handlers"
	"github.com/spf13/cobra"
)

// nextCmd represents the next command
var nextCmd = &cobra.Command{
	Use:   "next",
	Short: "Gets the next available sixers game date and time.",
	Long:  "Gets the next available sixers game date and time.",
	Run: func(cmd *cobra.Command, args []string) {
		handlers.Next(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(nextCmd)
}
