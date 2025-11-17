package scaffold

import (
	"os"

	"github.com/KhoalaS/Appify/embeds"
	"github.com/spf13/cobra"
)

var ScaffoldCmd = &cobra.Command{
	Use:   "scaffold",
	Short: "Generates an example config and userscripts.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return os.CopyFS("./", embeds.ExampleFiles)
	},
}
