package cli

import "github.com/spf13/cobra"

func newModuleCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "module",
		Aliases: []string{"m"},
		Short:   "commands relating to a single terraform module",
	}

	cmd.AddCommand(newGraphModuleCommand())

	return cmd
}

func newGraphModuleCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "graph",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	return cmd
}
