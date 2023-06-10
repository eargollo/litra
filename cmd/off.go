/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/derickr/go-litra-driver"
	"github.com/spf13/cobra"
)

// offCmd represents the off command
var offCmd = &cobra.Command{
	Use:   "off",
	Short: "Turns off Litra light",
	Long: `Turns off your Litra light. For example:

litra off`,
	Run: func(cmd *cobra.Command, args []string) {
		ld, err := litra.New()
		if err != nil {
			fmt.Printf("Could not find litra light: %v", err)
			return
		}
		defer ld.Close()
		ld.TurnOff()
	},
}

func init() {
	rootCmd.AddCommand(offCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// offCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// offCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
