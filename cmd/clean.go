// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"github.com/fanux/sealos/install"
	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Simplest way to clean your kubernets HA cluster",
	Long:  `sealos clean`,
	Run: func(cmd *cobra.Command, args []string) {
		beforeNodes:=install.ParseIPs(install.NodeIPs)
		c := &install.SealConfig{}
		c.Load("")
		install.BuildClean(beforeNodes)
		if len(beforeNodes) > 0{
			var resultNodes []string
			//去掉共同数据添加数据到结果集中去
			for _,node:=range c.Nodes{
				if !install.StrSliceContains(beforeNodes,node){
					resultNodes = append(resultNodes,node)
				}
			}
			c.Nodes = resultNodes
			c.Dump("")
		}
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)

	// Here you will define your flags and configuration settings.
	cleanCmd.Flags().StringSliceVar(&install.NodeIPs, "node", []string{}, "kubernetes multi-nodes ex. 192.168.0.5-192.168.0.5")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cleanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cleanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
