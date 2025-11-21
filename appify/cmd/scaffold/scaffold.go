package scaffold

import (
	"os"

	"github.com/KhoalaS/Appify/embeds"
	"github.com/spf13/cobra"
)

var withTypescript bool

var ScaffoldCmd = &cobra.Command{
	Use:   "scaffold",
	Short: "Generates an example config and userscripts.",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := os.CopyFS("./", embeds.ExampleFiles)
		if err != nil {
			return err
		}

		if withTypescript {
			err = os.CopyFS("./", embeds.TypescriptFiles)
		}

		return err
	},
}

func init() {
	ScaffoldCmd.Flags().BoolVarP(&withTypescript, "typescript", "t", false, "Add Typescript support")
}
