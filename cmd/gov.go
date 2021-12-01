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
	"github.com/spf13/cobra"
)

func init() {
	queryCmd.AddCommand(governanceQueryCmd)
	txCmd.AddCommand(governanceTxCmd)
	// TODO: add additional governance commands here
}

// governanceQueryCmd represents the governance command
var governanceQueryCmd = &cobra.Command{
	Use:     "governance",
	Aliases: []string{"gov"},
	Short:   "A brief description of your command",
}

// governanceTxCmd represents the governance command
var governanceTxCmd = &cobra.Command{
	Use:     "governance",
	Aliases: []string{"gov"},
	Short:   "A brief description of your command",
}