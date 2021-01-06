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
	"github.com/sanchezcl/dcmngr/support"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/exec"
	"strings"
	"time"
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up [SERVICE...]",
	Short: "Create and start containers",
	Long: `Builds, (re)creates, starts, and attaches to containers for a service
Unless they are already running, this command also starts any linked services.
If it's ran without arguments starts the defaults configured in the yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			args = viper.GetStringSlice("up_default_containers")
			fmt.Printf(
				"Starting default containers from yml: %s%s%s\n",
				support.ColorGreen,
				strings.Join(args, ","),
				support.ColorReset)
			time.Sleep(500*time.Millisecond)
		}
		args = append([]string{"up", "-d"}, args...)
		c := exec.Command("docker-compose", args...)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()
	},
}

func init() {
	rootCmd.AddCommand(upCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//upCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
