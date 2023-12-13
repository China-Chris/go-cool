package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var CmdRoot = &cobra.Command{
	Use:     "cool",
	Example: "cool new demo-api",
	Short:   "This command is cool!",
	Long:    fmt.Sprintf("\n ______   ______   ______   __          \n/_____/\\ /_____/\\ /_____/\\ /_/\\         \n\\:::__\\/ \\:::_ \\ \\\\:::_ \\ \\\\:\\ \\        \n \\:\\ \\  __\\:\\ \\ \\ \\\\:\\ \\ \\ \\\\:\\ \\       \n  \\:\\ \\/_/\\\\:\\ \\ \\ \\\\:\\ \\ \\ \\\\:\\ \\____  \n   \\:\\_\\ \\ \\\\:\\_\\ \\ \\\\:\\_\\ \\ \\\\:\\/___/\\ \n    \\_____\\/ \\_____\\/ \\_____\\/ \\_____\\/ \n                                        \n"),

	Run: func(cmd *cobra.Command, args []string) {
		// 根命令的执行逻辑
	},
}

func main() {
	if err := CmdRoot.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
