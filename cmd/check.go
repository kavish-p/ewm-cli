package cmd

import (
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check generated reports and auto-create defects",
}

func init() {
	rootCmd.AddCommand(checkCmd)

}
