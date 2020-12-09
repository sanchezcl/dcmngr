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
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type dcmngrYaml struct {
	BashDefaultContainerName string   `yaml:"bash_default_container_name"`
	BuildDefaultContainers   []string `yaml:"build_default_containers"`
	UpDefaultContainers      []string `yaml:"up_default_containers"`
}

// genymlCmd represents the genyml command
var genymlCmd = &cobra.Command{
	Use:   "genyml",
	Short: "Generate .dcmngr.yml file",
	Long:  `Generate .dcmngr.yml file to set the defaults containers for up, build and shell commands
based on existent docker-compose.yml in your project`,
	Run: func(cmd *cobra.Command, args []string) {
		if support.FileExist(ConfigFileName) {
			if !support.AskForConfirmation(fmt.Sprintf("%s already exists overwrite?", ConfigFileName)) {
				fmt.Println("bye!")
				return
			}
		}

		var containerNames []string
		if _, err := os.Stat("docker-compose.yml"); os.IsNotExist(err) {
			fmt.Printf("docker-compose.yml does not exist\n")
		}

		containerNames = extractContainerNamesFromYml()
		ymlExport := dcmngrYaml{
			BashDefaultContainerName: "",
			BuildDefaultContainers:   containerNames,
			UpDefaultContainers:      containerNames,
		}
		ymlExportBytes, err := yaml.Marshal(&ymlExport)
		if err != nil {
			fmt.Errorf("Error marshaling yml %s", err)
		}
		err = ioutil.WriteFile(ConfigFileName, ymlExportBytes, 0644)
		fmt.Printf("\n%s%s%s created.\n\n", support.ColorGreen, ConfigFileName, support.ColorReset)
		fmt.Println("Your docker-compose.yml was parsed and all the services")
		fmt.Println("included in as default containers to up and build. You")
		fmt.Printf("can change this list in the %s\n", ConfigFileName)
	},
}

func init() {
	rootCmd.AddCommand(genymlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genymlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genymlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func extractContainerNamesFromYml() []string {
	ymlFile, err := ioutil.ReadFile("docker-compose.yml")
	if err != nil {
		log.Fatalf("can't open file...\n")
	}
	var dc interface{}
	err = yaml.Unmarshal(ymlFile, &dc)
	if err != nil {
		log.Fatalf("Unmarshal: %v\n", err)
	}
	servicesMap := dc.(map[interface{}]interface{})["services"]
	var cn []string
	for k := range servicesMap.(map[interface{}]interface{}) {
		cn = append(cn, k.(string))
	}
	return cn
}
