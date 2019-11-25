package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:    "completion",
	Args:   cobra.NoArgs,
	Hidden: true,
}

func init() {
	rootCmd.AddCommand(completionCmd)

	completionCmd.AddCommand(&cobra.Command{
		Use:  "bash",
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			rootCmd.GenBashCompletion(os.Stdout)
		},
	})

	completionCmd.AddCommand(&cobra.Command{
		Use:  "zsh",
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			rootCmd.GenZshCompletion(os.Stdout)
		},
	})

	completionCmd.AddCommand(&cobra.Command{
		Use:  "powershell",
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			rootCmd.GenPowerShellCompletion(os.Stdout)
		},
	})
}
