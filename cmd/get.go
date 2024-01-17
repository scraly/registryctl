package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/ovh/go-ovh/ovh"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get your OVHcloud Managed Private Registries",
	Long:  `Get your OVHcloud Managed Private Registries.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Retrieve OVHcloud credentials needed to call API
		OVH_ENDPOINT := viper.GetString("OVH_ENDPOINT")
		OVH_APPLICATION_KEY := viper.GetString("OVH_APPLICATION_KEY")
		OVH_APPLICATION_SECRET := viper.GetString("OVH_APPLICATION_SECRET")
		OVH_CONSUMER_KEY := viper.GetString("OVH_CONSUMER_KEY")
		OVH_CLOUD_PROJECT_SERVICE := viper.GetString("OVH_CLOUD_PROJECT_SERVICE")

		// Check environment variables are filled
		if OVH_ENDPOINT == "" || OVH_APPLICATION_KEY == "" || OVH_APPLICATION_SECRET == "" || OVH_CONSUMER_KEY == "" || OVH_CLOUD_PROJECT_SERVICE == "" {
			err := fmt.Errorf("Error: set OVHcloud environment variables in .env file: OVH_ENDPOINT, OVH_APPLICATION_KEY, OVH_APPLICATION_SECRET, OVH_CONSUMER_KEY, OVH_CLOUD_PROJECT_SERVICE")
			panic(err)
		}

		client, _ := ovh.NewClient(
			OVH_ENDPOINT,
			OVH_APPLICATION_KEY,
			OVH_APPLICATION_SECRET,
			OVH_CONSUMER_KEY,
		)

		//fmt.Printf("[DEBUG] Will read cloud project registries for project: %s", OVH_CLOUD_PROJECT_SERVICE)

		regs := []CloudProjectContainerRegistry{}

		endpoint := fmt.Sprintf(
			"/cloud/project/%s/containerRegistry",
			url.PathEscape(OVH_CLOUD_PROJECT_SERVICE),
		)
		err := client.Get(endpoint, &regs)
		if err != nil {
			fmt.Errorf("Error calling GET %s:\n\t %q", endpoint, err)
		}

		var regList = make([][]string, len(regs))

		for _, reg := range regs {
			// Add the registry in the list of registries
			myRegistry := []string{reg.Name, reg.Status, reg.Region, reg.Url, reg.Version, reg.CreatedAt}
			regList = append(regList, myRegistry)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Status", "Region", "URL", "Version", "Created_At"})
		table.SetAutoWrapText(false)
		table.SetAutoFormatHeaders(true)
		table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.SetCenterSeparator("")
		table.SetColumnSeparator("")
		table.SetRowSeparator("")
		table.SetHeaderLine(false)
		table.SetTablePadding("\t") // pad with tabs
		table.SetNoWhiteSpace(true)
		table.AppendBulk(regList) // Add Bulk Data
		table.Render()
	},
}

type CloudProjectContainerRegistry struct {
	CreatedAt string `json:"createdAt"`
	Id        string `json:"id"`
	Name      string `json:"name"`
	ProjectID string `json:"projectID"`
	Region    string `json:"region"`
	Size      int64  `json:"size"`
	Status    string `json:"status"`
	UpdatedAt string `json:"updatedAt"`
	Url       string `json:"url"`
	Version   string `json:"version"`
}

func init() {
	ovhMprCmd.AddCommand(getCmd)
}
