package cmd

import (
	"github.com/josa42/git-repo/utils"
	"github.com/spf13/cobra"
)

var compareCmd = &cobra.Command{
	Use:  "compare",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		arguments := map[string]interface{}{}
		utils.GetRepo().Open("compare", arguments)
	},
}

func init() {
	rootCmd.AddCommand(compareCmd)
}
