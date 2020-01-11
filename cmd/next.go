/*
Copyright © 2019 Graham Plata <graham.plata@gmail.com>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

// nextCmd represents the next command
var nextCmd = &cobra.Command{
	Use:   "next",
	Short: "Gets the next available sixers game date and time.",
	Long:  "Gets the next available sixers game date and time.",
	Run: func(cmd *cobra.Command, args []string) {
		currentDate := time.Now()
		fmt.Println(currentDate)
	},
}

func init() {
	rootCmd.AddCommand(nextCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nextCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nextCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
