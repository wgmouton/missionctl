package main

import "github.com/spf13/cobra"
import (
	"strings"
)

var rootToolCmd = &cobra.Command{
	Use:     "tools",
	Aliases: []string{"tool"},
	Example: "tool [toolname] [toolargs]",
	Short:   "Execute one of your defined tools with their arguments.",
}

func buildToolsCmd(name string, tool Tool) *cobra.Command {
	lName := strings.ToLower(name)

	cmd := &cobra.Command{
		Use:     lName,
		Aliases: []string{name, tool.Binary},
		Run: func(cmd *cobra.Command, args []string) {
			Cmd(tool.Binary, strings.Join(tool.Args, " "), strings.Join(args, " "))(true)
		},
	}

	return cmd
}

func toolCommand(tools *MissionTools) *cobra.Command {

	for key, value := range tools.Tools {
		rootToolCmd.SuggestFor = append(rootToolCmd.SuggestFor, key)
		rootToolCmd.AddCommand(buildToolsCmd(key, value))
	}

	return rootToolCmd
}
