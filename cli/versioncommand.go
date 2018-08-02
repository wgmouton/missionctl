package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func versionCmd(version Version) *cobra.Command {
	return &cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version.GetVersion())
		},
	}
}

var setCmd = &cobra.Command{
	Use: "set",
}

var incrementCmd = &cobra.Command{
	Use: "increment",
}

func incrementCommand(version Version) {
	var incrementMajorCmd = &cobra.Command{
		Use: "major",
	}

	var incrementMinorCmd = &cobra.Command{
		Use: "minor",
	}

	var incrementBugFixCmd = &cobra.Command{
		Use: "bugfix",
	}

	var incrementLabelCmd = &cobra.Command{
		Use: "label",
	}

	incrementCmd.AddCommand(incrementMajorCmd)
	incrementCmd.AddCommand(incrementMinorCmd)
	incrementCmd.AddCommand(incrementBugFixCmd)
	incrementCmd.AddCommand(incrementLabelCmd)

}

func VersionCommand(version Version) *cobra.Command {
	incrementCommand(version)

	_versionCmd := versionCmd(version)
	_versionCmd.AddCommand(setCmd)
	_versionCmd.AddCommand(incrementCmd)
	return _versionCmd
}
