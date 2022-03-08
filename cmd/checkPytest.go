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
	"strings"

	"github.com/kavish-p/ewm-cli/oslc"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

var pytestCmd = &cobra.Command{
	Use:   "pytest",
	Short: "Check Pytest JSON output file and create a defect workitem on IBM EWM for each failed test case",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		reportPath, _ := cmd.Flags().GetString("report")

		addDesc := cmd.Flags().Lookup("addDesc").Changed

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

			cleanErrorMsg := strings.Replace(errorMessage.Str, "\n", " ", -1)

			if addDesc {
				addDescString, _ := cmd.Flags().GetString("addDesc")
				cleanErrorMsg = cleanErrorMsg + "\n\n" + addDescString
			}

			oslc.CreateDefect("Failed Functional Test: "+testDesc.Str, cleanErrorMsg)
		}
	},
}

func init() {
	checkCmd.AddCommand(pytestCmd)

	pytestCmd.PersistentFlags().String("report", "", "json report of Pytest execution")
	pytestCmd.MarkPersistentFlagRequired("report")

	pytestCmd.PersistentFlags().String("addDesc", "", "additional description to add to defect")
}
