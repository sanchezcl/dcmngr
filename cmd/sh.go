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
)

// shCmd represents the sh command
var (
	isAdmin bool
	shCmd   = &cobra.Command{
		Use:   "sh [container name]",
		Short: "Get a shell inside a container.",
		Long:  `Get a shell inside a container.`,
		Run: func(cmd *cobra.Command, args []string) {
			if !support.FileExist(ConfigFileName) {
				fmt.Println("yaml config file doesn't exist, please create one with:")
				fmt.Println("\t $ dcmngr genyml")
				fmt.Println("Or provide one with the flag --config [route to the config file]")
				return
			}

			defaultShContainerName := viper.GetString("sh_default_container_name")
			alwaysAdmin := viper.GetBool("sh_always_admin")

			args = []string{"exec", "-it"}
			if !(isAdmin || alwaysAdmin) {
				args = append(args, "-u 1000")
			}
			args = append(args, []string{defaultShContainerName, "bash"}...)

			c := exec.Command("docker", args...)
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
