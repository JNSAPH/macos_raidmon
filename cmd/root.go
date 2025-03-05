package cmd

import (
	"fmt"
	"runtime"

	"github.com/JNSAPH/macos_raidmon/core"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var devFlag bool

var rootCmd = &cobra.Command{
	Use:   "moss",
	Short: "Moss: Macos Server",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if runtime.GOOS != "darwin" {
			logrus.Fatal("This application is only supported on macOS")
		}

		if devFlag {
			logrus.SetLevel(logrus.DebugLevel)
			logrus.Debug("Running in development mode")
		}

		core.SetupLogger()
		core.SetupConfig()
		core.SetupSystray()

		// Stop the program from exiting
		select {}
	},
}

func init() {
	// Register the -dev flag with the root command
	rootCmd.Flags().BoolVarP(&devFlag, "dev", "d", false, "Run in development mode")

	// Register a Config path flag with the root command
	rootCmd.Flags().StringVarP(&core.ConfigPath, "config", "c", fmt.Sprintf("%v/%v", core.DefaultConfigPath, core.DefaultConfigName), "Path to the config.yaml")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
