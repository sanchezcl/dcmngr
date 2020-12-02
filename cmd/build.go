/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"github.com/sanchezcl/dcmngr/support"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/exec"
	"strings"
	"time"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build or rebuild services",
	Long:  `Build or rebuild services`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			args = viper.GetStringSlice("build_default_containers")
			fmt.Printf(
				"Building default containers from yml: %s%s%s\n",
				support.ColorGreen,
				strings.Join(args, ","),
				support.ColorReset)
			time.Sleep(500*time.Millisecond)
		}
		args = append([]string{"build"}, args...)
		c := exec.Command("docker-compose", args...)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
