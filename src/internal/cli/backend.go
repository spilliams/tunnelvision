package cli

import "github.com/spf13/cobra"

func newBackendCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "backend",
		Aliases: []string{"b"},
		Short:   "commands relating to a single backend provider",
	}

	cmd.AddCommand(newGraphBackendCommand())
	cmd.AddCommand(newListBackendCommand())
	cmd.AddCommand(newSearchBackendCommand())
	cmd.AddCommand(newShowBackendCommand())

	return cmd
}

func newGraphBackendCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "graph",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	return cmd
}

func newListBackendCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	return cmd
}

func newSearchBackendCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "search",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	return cmd
}

func newShowBackendCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	return cmd
}
