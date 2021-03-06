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
)

// shCmd represents the sh command
var (
	isAdmin bool
	shCmd   = &cobra.Command{
		Use:   "sh [service name]",
		Short: "Get a shell inside a container.",
		Long:  `Get a shell inside a container.

Examples:
dcmngr sh : 		get shell in the default container configured in .dcmngr.yml
dcmngr sh redis: 	get shell in the redis service`,
		Args: cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if !support.FileExist(ConfigFileName) {
				fmt.Println("yaml config file doesn't exist, please create one with:")
				fmt.Println("\t $ dcmngr genyml")
				fmt.Println("Or provide one with the flag --config [route to the config file]")
				return
			}

			alwaysAdmin := viper.GetBool("sh_always_admin")

			commandArgs := []string{"exec"}
			if !(isAdmin || alwaysAdmin) {
				commandArgs = append(commandArgs, "-u", "1000")
			}

			if len(args) <= 0 {
				defaultContainer := strings.Trim(viper.GetString("sh_default_service"), " ")
				if defaultContainer == "" {
					support.PrintError("sh_default_service not configured in .dcmgr.yaml")
					return
				}
				commandArgs = append(commandArgs, defaultContainer)
			} else {
				commandArgs = append(commandArgs, args...)
			}

			commandArgs = append(commandArgs, "bash")

			c := exec.Command("docker-compose", commandArgs...)
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			c.Stdin = os.Stdin
			c.Run()
			fmt.Printf("container sh... %s%s%s\n", support.ColorGreen, "end", support.ColorReset)
		},
	}
)

func init() {
	rootCmd.AddCommand(shCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// shCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	shCmd.Flags().BoolVarP(&isAdmin, "admin", "a", false, "connect as admin")
}
