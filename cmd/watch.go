/*
Copyright © 2020 Carlos Sanchez

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
	"github.com/sanchezcl/dcmngr/models"
	"github.com/sanchezcl/dcmngr/support"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/exec"
)

// watchCmd represents the watch command
var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//check if parameter is set and vaild in the yml | not error
		var settings models.DcmngrYaml
		viper.Unmarshal(&settings)
		wc := &settings.WatchConfigs
		if wc.Command == "" || wc.Service == "" {
			fmt.Printf("Can't execute watch command missing configs in %s\n", ConfigFileName)
			fmt.Println("please complete the configs to use this feature...")
			return
		}

		execArgs := []string{"exec", wc.Service, wc.Command}
		execArgs = append(execArgs, wc.Args...)

		c := exec.Command("docker-compose", execArgs...)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Stdin = os.Stdin
		c.Run()
		fmt.Printf("And Now His Watch Is %sEnded%s... \n", support.ColorGreen, support.ColorReset)
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// watchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// watchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
