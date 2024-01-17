/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// ovhMprCmd represents the ovhMpr command
var ovhMprCmd = &cobra.Command{
	Use:   "ovhMpr",
	Short: "Manage your OVHcloud Managed Private Registries",
	Long:  `Manage (get) your OVHcloud Managed Private Registries.`,
}

func init() {
	rootCmd.AddCommand(ovhMprCmd)
}
