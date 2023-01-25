package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

var config string

var rootCmd = &cobra.Command{
	Use: "hello",
	Run: func(cmd *cobra.Command, args []string) {
		t := map[string]interface{}{}
		buffer, err := ioutil.ReadFile(config)
		err = yaml.Unmarshal(buffer, &t)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v", t)

		data, err := json.Marshal(t)

		err = ioutil.WriteFile("config.json", data, 0644)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&config, "config", "c", "", "")

}
