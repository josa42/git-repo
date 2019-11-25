package cmd

import (
	"github.com/josa42/git-repo/utils"
	"github.com/spf13/cobra"
)

var commitsCmd = &cobra.Command{
	Use:  "commits",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		arguments := map[string]interface{}{}
		utils.GetRepo().Open("commits", arguments)
	},
}

func init() {
	rootCmd.AddCommand(commitsCmd)
}
