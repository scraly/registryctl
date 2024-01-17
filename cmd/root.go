package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "registryctl",
	Short: "OCI/Docker Registries CLI",
	Long:  `OCI/Docker Registries CLI allows you to manage your various OCI/Docker images registries (Docker Hub, OVHcloud Managed Private Registry, Artifactory, Google Container Registry, Harbor, Docker registries…)`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Read the OVHcloud credentials
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error: create a .env file and add your environment variables")
		panic(err)
	}

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
