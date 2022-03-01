package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goddes",
	Short: "A simple app to create heavy internet traffic",
	Long: `The app is designed to have maximum performance while sending a lot of requests on server.
With it you can easily generate a heavy load on the targeted server.
It has outperformed most of python apps known to me,
and is slightly faster than another similar app written in Golang.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
