/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/kavish-p/ewm-cli/oslc"
	"github.com/spf13/cobra"
)

// workItemTypeCmd represents the type command
var workItemTypeCmd = &cobra.Command{
	Use:   "type",
	Short: "Get Work Item types in Project Area",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		context, _ := cmd.Flags().GetString("context")
		oslc.GetWorkItemType(context)
	},
}

func init() {
	getCmd.AddCommand(workItemTypeCmd)
	workItemTypeCmd.PersistentFlags().String("context", "", "context ID of project area")
	workItemTypeCmd.MarkPersistentFlagRequired("context")
}
