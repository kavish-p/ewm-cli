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
	"io/ioutil"
	"log"
	"strconv"

	"github.com/kavish-p/ewm-cli/oslc"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

var pytestCmd = &cobra.Command{
	Use:   "pytest",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		reportPath, _ := cmd.Flags().GetString("report")

		report, err := ioutil.ReadFile(reportPath)
		if err != nil {
			log.Fatal(err)
		}
		failedTests := gjson.Get(string(report), "tests.#(outcome=failed)#").Array()
		failedTestsNum := len(failedTests)

		log.Println("Found " + strconv.FormatInt(int64(failedTestsNum), 10) + " failed tests in Pytest JSON report located at " + reportPath)

		for i := 0; i < failedTestsNum; i++ {
			// log.Println(failedTests[i])
			testDesc := gjson.Get(failedTests[i].String(), "nodeid")
			errorMessage := gjson.Get(failedTests[i].String(), "call.crash.message")

			oslc.CreateDefect("Failed Functional Test: "+testDesc.Str, errorMessage.Str)
			log.Println("Defect Logged Successfully")
		}
	},
}

func init() {
	checkCmd.AddCommand(pytestCmd)

	pytestCmd.PersistentFlags().String("report", "", "json report of Pytest execution")
	pytestCmd.MarkPersistentFlagRequired("report")
}
