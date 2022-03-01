package cmd

import (
	"github.com/kiIIer/goddes/hell"
	"github.com/spf13/cobra"
)

var bombardCmd = &cobra.Command{
	Use:   "hell",
	Short: "Command to make the hell on site",
	Long: `This command launches multiple go routines,
which all simultaneously start sending requests on provided url.
It is recommended to use approximately cpu*500 go routines to maximise the dmg.

But feel free to experiment.`,
	Run: func(cmd *cobra.Command, args []string) {
		gophersCount, _ := cmd.Flags().GetInt("gophers")
		method, _ := cmd.Flags().GetString("method")
		url, _ := cmd.Flags().GetString("url")

		army := hell.Hell{
			RawUrl:       url,
			GophersCount: gophersCount,
			Method:       method,
		}

		army.Start()
	},
}

func init() {
	rootCmd.AddCommand(bombardCmd)

	bombardCmd.Flags().String("url", "", "Url to be targeted")
	bombardCmd.Flags().Int("gophers", 500, "Number of goroutines to be used to make hell")
	bombardCmd.Flags().String("method", "GET", "Method to be used when raining hell")
	bombardCmd.MarkFlagRequired("url")
}
