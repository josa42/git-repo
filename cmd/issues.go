package cmd

import (
	"github.com/josa42/git-repo/utils"
	"github.com/spf13/cobra"
)

var issuesCmd = &cobra.Command{
	Use:  "issues",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		arguments := map[string]interface{}{}
		utils.GetRepo().Open("issues", arguments)
	},
}

func init() {
	rootCmd.AddCommand(issuesCmd)
}
