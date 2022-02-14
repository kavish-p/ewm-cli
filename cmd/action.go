package cmd

import (
	"github.com/kavish-p/ewm-cli/oslc"
	"github.com/spf13/cobra"
)

var actionCmd = &cobra.Command{
	Use:   "action",
	Short: "Perform workflow action on a work item",

	Run: func(cmd *cobra.Command, args []string) {
		taskID, _ := cmd.Flags().GetString("taskID")
		action, _ := cmd.Flags().GetString("action")

		oslc.PerformAction(taskID, action)
	},
}

func init() {
	rootCmd.AddCommand(actionCmd)

	actionCmd.PersistentFlags().String("taskID", "", "ID of task work item")
	actionCmd.PersistentFlags().String("action", "", "workflow action to perform on work item")

	actionCmd.MarkPersistentFlagRequired("taskID")
	actionCmd.MarkPersistentFlagRequired("action")

}
