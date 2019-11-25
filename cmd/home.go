package cmd

import (
	"github.com/josa42/git-repo/utils"
	"github.com/spf13/cobra"
)

var homeCmd = &cobra.Command{
	Use:  "home",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		arguments := map[string]interface{}{}
		utils.GetRepo().Open("home", arguments)
	},
}

func init() {
	rootCmd.AddCommand(homeCmd)
}
