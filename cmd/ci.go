package cmd

import (
	"github.com/josa42/git-repo/utils"
	"github.com/spf13/cobra"
)

var ciCmd = &cobra.Command{
	Use:  "ci",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		// TODO refactor flags handling
		f := func(name string) bool {
			b, _ := cmd.Flags().GetBool(name)
			return b
		}

		arguments := map[string]interface{}{
			"--appveyor":  f("appveyor"),
			"--bitbucket": f("bitbucket"),
			"--circle":    f("circle"),
			"--github":    f("github"),
			"--jenkins":   f("jenkins"),
			"--travis":    f("travis"),
		}
		utils.GetRepo().Open("ci", arguments)
	},
}

func init() {
	rootCmd.AddCommand(ciCmd)

	ciCmd.Flags().BoolP("appveyor", "", false, "")
	ciCmd.Flags().BoolP("bitbucket", "", false, "")
	ciCmd.Flags().BoolP("circle", "", false, "")
	ciCmd.Flags().BoolP("github", "", false, "")
	ciCmd.Flags().BoolP("jenkins", "", false, "")
	ciCmd.Flags().BoolP("travis", "", false, "")
}
