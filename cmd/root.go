/*
Copyright © 2023 DjLeChuck <djlechuck@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/djlechuck/fa-updater/internal/config"
	"github.com/djlechuck/fa-updater/internal/logger"
	"github.com/djlechuck/fa-updater/internal/patreon"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type application struct {
	patreon *patreon.Patreon
	config  *config.Config
}

var cfgFile string

type contextKey string

const appContextKey = contextKey("app")

var (
	version string
	date    string
	commit  string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fa-updater",
	Short: "Easily keep up-to-date your FA files",
	Long: `FA updater is a simple tool that will allow you to keep your FA files up to date. It will analyze the versions you have and check if more recent ones exist in order to download them easily.

First, be sure to define the directory which contains your assets (fa-updater setDirectory), then launch the update process (fa-updater updateAssets).`,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	Version: fmt.Sprintf("FA udpater (c) 2023 DjLeChuck -- v%s - %s - %s\n", version, date, commit),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	app := &application{
		patreon: &patreon.Patreon{},
		config: &config.Config{
			Dungeondraft: &config.Dungeondraft{
				ThumbnailsVersions: make(map[string]bool),
			},
			Tokens: &config.Tokens{
				DownloadedPacks: make(map[string]bool),
			},
		},
	}
	ctx := context.WithValue(context.Background(), appContextKey, app)

	err := rootCmd.ExecuteContext(ctx)
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	logger.Init()

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.fa-updater.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".fa-updater" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".fa-updater")
		viper.SafeWriteConfig()
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		logger.Fatalf(err, "Error reading config file %s", viper.ConfigFileUsed())
	}
}
