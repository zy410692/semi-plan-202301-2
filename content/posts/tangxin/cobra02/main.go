package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	err := root.Execute()
	if err != nil {
		panic(err)
	}
}

var root = &cobra.Command{
	Use:   "greeting",
	Short: "打招呼",
	Run: func(cmd *cobra.Command, args []string) {
		greeting(name, age)
	},
}

var (
	name string
	age  int
)

func init() {
	root.Flags().StringVarP(&name, "name", "n", "", "名字")
	root.Flags().IntVarP(&age, "age", "a", 20, "名字")
}

func greeting(name string, age int) {
	fmt.Printf("你好 %s 先生, 你今天 %d 岁\n", name, age)
}
