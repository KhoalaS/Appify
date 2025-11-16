package main

import (
	appify "github.com/KhoalaS/Appify"
	"github.com/KhoalaS/Appify/pkg/core"
)

func main() {

	config, err := core.ReadConfigFromFile("./__test__/config.json")

	if err != nil {
		panic(err)
	}

	err = core.RenderTemplate(*config, appify.TemplateFolder)
	if err != nil {
		panic(err)
	}

}
