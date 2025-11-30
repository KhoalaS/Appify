package build

import (
	"errors"
	"os"

	"github.com/KhoalaS/Appify/pkg/core"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// TODO: Custom adb errors

var configPath string
var appVariant core.AppVariant = core.DebugAppVariant

var BuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the app",
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := core.ReadConfigFromFile(configPath)
		if err != nil {
			return err
		}

		_, err = os.Stat(config.ProjectDirectory)
		if err != nil {
			return errors.New("project directory does not exist")
		}

		adb := &core.ADB{}

		err = adb.Connect()
		if err != nil {
			log.Err(err).Msg("could not connect to adb")
			return err
		}

		err = adb.BuildApp(config.ProjectDirectory, appVariant)
		if err != nil {
			log.Err(err).Msg("could not build the app")
			return err
		}

		err = adb.InstallApp(config.ProjectDirectory, appVariant)
		if err != nil {
			log.Err(err).Msg("could not install the app")
			return err
		}

		err = adb.StartApp(config.PackageName)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	BuildCmd.Flags().StringVarP(&configPath, "config", "c", "./config.json", "The path to the project configuration.")
	BuildCmd.Flags().VarP(&appVariant, "type", "t", `The app variant, either "release" or "debug".`)
}
