package cmd

import (
	"github.com/spf13/cobra"
)

const rootCmdName = "kookbook"

func NewRootCmd() *cobra.Command {
	cmdInstance := &cobra.Command{
		Use:           rootCmdName,
		Short:         rootCmdName,
		Long:          rootCmdName,
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmdInstance.AddCommand(
		NewServeCmd(),
	)

	return cmdInstance
}
