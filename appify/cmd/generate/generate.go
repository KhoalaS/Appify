package cmd

import (
	"github.com/KhoalaS/Appify/embeds"
	"github.com/KhoalaS/Appify/pkg/core"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at https://gohugo.io/documentation/`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := core.ReadConfigFromFile("./__test__/config.json")

		if err != nil {
			panic(err)
		}

		err = core.RenderTemplate(*config, embeds.TemplateFolder, embeds.AppCodeFolder)
		if err != nil {
			panic(err)
		}
	},
}
