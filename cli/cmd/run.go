package cmd

import (
	"errors"
	//"fmt"

	"github.com/spf13/cobra"
	"github.com/devhindo/plan/cli/core"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Execute the command",
	Long:  `Execute the command`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("filename is required")
		}
		return nil
	},
	Run: run,
}

func run(cmd *cobra.Command, args []string) {
	core.RUN(args[0])
}

func init() {
	rootCmd.AddCommand(runCmd)
}