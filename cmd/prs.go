package cmd

import (
	"github.com/josa42/git-repo/utils"
	"github.com/spf13/cobra"
)

var prsCmd = &cobra.Command{
	Use:  "prs",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		arguments := map[string]interface{}{}
		utils.GetRepo().Open("prs", arguments)
	},
}

func init() {
	rootCmd.AddCommand(prsCmd)
}
