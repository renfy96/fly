package fly

import (
	"fmt"
	"github.com/renfy96/fly/internal/command/create"
	"github.com/renfy96/fly/internal/command/project"
	"github.com/renfy96/fly/internal/command/run"
	"github.com/renfy96/fly/internal/command/upgrade"
	"github.com/renfy96/fly/internal/command/wire"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "fly",
	Example: "fly new demo-api",
	Short:   " ******** **       **    **\n/**///// /**      //**  ** \n/**      /**       //****  \n/******* /**        //**   \n/**////  /**         /**   \n/**      /**         /**   \n/**      /********   /**   \n//       ////////    // ",
	Version: fmt.Sprintf(" ******** **       **    **\n/**///// /**      //**  ** \n/**      /**       //****  \n/******* /**        //**   \n/**////  /**         /**   \n/**      /**         /**   \n/**      /********   /**   \n//       ////////    // fly %s - Copyright (c) 2025 fly\nReleased under the MIT License.\n\n", "v1"),
}

func init() {
	rootCmd.AddCommand(project.NewCmd)
	rootCmd.AddCommand(create.CreateCmd)
	rootCmd.AddCommand(wire.WireCmd)
	rootCmd.AddCommand(run.RunCmd)
	rootCmd.AddCommand(upgrade.UpgradeCmd)
	// 新增application
	create.CreateCmd.AddCommand(create.CreateApplicationCmd)
	create.CreateCmd.AddCommand(create.CreateBCCmd)
	create.CreateCmd.AddCommand(create.CreateAllCmd)
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
