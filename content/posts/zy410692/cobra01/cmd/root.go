package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var name string
var age string

var rootCmd = &cobra.Command{
	Use: "hello",
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			fmt.Println("你好，今年 " + fmt.Sprintf("%s", age) + " 岁")
		} else {
			fmt.Println(fmt.Sprintf("%s", name) + " 你好，今年 " + fmt.Sprintf("%s", age) + " 岁")
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
	rootCmd.Flags().StringVarP(&name, "name", "", "", "")
	rootCmd.Flags().StringVarP(&age, "age", "", "20", "")
}
