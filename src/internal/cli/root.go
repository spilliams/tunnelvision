package cli

import "github.com/spf13/cobra"

func newRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "root",
		Aliases: []string{"r"},
		Short:   "commands relating to a single terraform root",
	}

	cmd.AddCommand(newGraphRootCommand())

	return cmd
}

func newGraphRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "graph",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	return cmd
}
