package cmd

import (
	"github.com/josa42/git-repo/utils"
	"github.com/spf13/cobra"
)

var prCmd = &cobra.Command{
	Use:  "pr",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		arguments := map[string]interface{}{}
		utils.GetRepo().Open("pr", arguments)
	},
}

func init() {
	rootCmd.AddCommand(prCmd)
}
