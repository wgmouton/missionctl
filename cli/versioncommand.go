package main

import (
	"flag"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
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
	Run: func(cmd *cobra.Command, args []string) {
		strconv.ParseUint(cmd.Flag("major").Value.String(), 8, 8)

	},
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

var fMajor = flag.Uint("major", 0, "major xx")

func VersionCommand(version Version) *cobra.Command {
	incrementCommand(version)

	setCmd.Flags().Uint8("major", version.Major, "help")
	setCmd.Flags().Uint8("minor", version.Minor, "help")
	setCmd.Flags().Uint8("patch", version.Patch, "help")
	setCmd.Flags().String("label", version.Label.getLabel(), "help")
	setCmd.Flags().String("meta", version.Meta, "help")

	_versionCmd := versionCmd(version)
	_versionCmd.AddCommand(setCmd)
	_versionCmd.AddCommand(incrementCmd)
	return _versionCmd
}
