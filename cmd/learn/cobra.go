package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "learn",
	Short: "测试程序",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("测试")
		return nil
	},
}

func Cobra() {
	_ = rootCommand.Execute()
}
