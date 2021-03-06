// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/getlunaform/lunaform/client/providers"
	"github.com/spf13/cobra"
)

var (
	tfProviderConfigurationListProviderNameFlag string
)

// tfProviderConfigurationListCmd represents the tfProviderConfigurationList command
var tfProviderConfigurationListCmd = &cobra.Command{
	Use:   "list-configurations",
	Short: "List provider configurations",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		params := providers.NewListProviderConfigurationsParams().
			WithProviderName(tfProviderConfigurationListProviderNameFlag)
		providers, err := lunaformClient.Providers.ListProviderConfigurations(
			params, authHandler)
		if err == nil {
			handleOutput(cmd, providers.Payload, useHal, err)
		} else {
			handlerError(err)
		}
	},
}

func init() {
	flags := tfProviderConfigurationListCmd.Flags()
	flags.StringVar(&tfProviderConfigurationListProviderNameFlag,
		"provider-name", "", "Provider Name")
	tfProviderConfigurationListCmd.MarkFlagRequired("provider-name")
}
