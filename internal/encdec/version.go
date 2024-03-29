package encdec

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of encdec",
	Long:  `All software has versions. This is encdec's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("encdec v0.0.2")
	},
}
