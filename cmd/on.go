/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/derickr/go-litra-driver"
	"github.com/spf13/cobra"
)

// onCmd represents the on command
var onCmd = &cobra.Command{
	Use:   "on",
	Short: "Turns on your Litra light",
	Long: `Turns on your Litra light. Optionally you can set brightness and temperature.
For example:

litra on --brightness 100 --temperature 100`,
	Run: func(cmd *cobra.Command, args []string) {
		ld, err := litra.New()
		if err != nil {
			fmt.Printf("Could not find litra light: %v", err)
			return
		}

		brightness, err := cmd.Flags().GetInt("brightness")
		if err != nil {
			fmt.Printf("Could not get brightness: %v, keeping existing value", err)
			brightness = 0
		}
		if brightness > 0 {
			fmt.Printf("Setting brightness to %d\n", brightness)
			ld.SetBrightness(brightness)
		}

		temperature, err := cmd.Flags().GetInt("temperature")
		if err != nil {
			fmt.Printf("Could not get temperature: %v, keeping existing value", err)
			temperature = 0
		}
		if temperature > 0 {
			fmt.Printf("Setting temperature to %d\n", temperature)
			ld.SetTemperature((int16)(temperature))
		}

		defer ld.Close()
		ld.TurnOn()
	},
}

func init() {
	rootCmd.AddCommand(onCmd)

	onCmd.Flags().IntP("brightness", "b", 0, "Brightness value between 1 and 100")
	onCmd.Flags().IntP("temperature", "t", 0, "Temperature value between 2700 (cool) and 6500 (warm)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// onCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// onCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
