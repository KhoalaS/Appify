package generate

import (
	"github.com/KhoalaS/Appify/embeds"
	"github.com/KhoalaS/Appify/pkg/core"
	"github.com/spf13/cobra"
)

var configPath string

var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate the Android project.",
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := core.ReadConfigFromFile("./__test__/config.json")

		if err != nil {
			return err
		}

		err = core.RenderTemplate(*config, embeds.TemplateFolder, embeds.AppCodeFolder)
		if err != nil {
			return err

		}

		return nil
	},
}

func init() {
	// Local flag (only for `sub`)
	GenerateCmd.Flags().StringVarP(&configPath, "config", "c", "./config.json", "The path to the project configuration.")
}
