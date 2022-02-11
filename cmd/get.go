package cmd

import (
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get resources such as contexts, workflow types, work item types, ..",
}

func init() {
	rootCmd.AddCommand(getCmd)

}
