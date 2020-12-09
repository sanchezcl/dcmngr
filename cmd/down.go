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
	"log"
	"os/exec"
)

// downCmd represents the down command
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Stop and remove containers, networks, images, and volumes",
	Long: `Stops containers and removes containers, networks, volumes, and images
created by 'up'.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Shutting down ...")

		c := exec.Command("docker-compose", "down")
		_ , err := c.CombinedOutput()
		if err != nil {
			fmt.Printf("%sFail%s\n", support.ColorRed, support.ColorReset)
			log.Fatalf("cmd.Run() failed with %s\n", err)
		}

		fmt.Printf("%sDone%s\n", support.ColorGreen, support.ColorReset)
	},
}

func init() {
	rootCmd.AddCommand(downCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
