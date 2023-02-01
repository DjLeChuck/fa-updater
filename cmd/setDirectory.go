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
	"github.com/djlechuck/fa-updater/internal/config"
	"github.com/djlechuck/fa-updater/internal/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// setDirectoryCmd define the directory which contains the FA assets
var setDirectoryCmd = &cobra.Command{
	Use:   "setDirectory [type] [path]",
	Short: "Define the directory which contains your files",
	Long: `Define the directory which contains your files.

[type] can be one of the following:
	* dungeondraft: will be used by the updateAssets command to get newer versions if exists
	* tokens: will be used by the updateTokens command to get newer versions if exists`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		dirType := args[0]

		var configKey string

		switch dirType {
		case "tokens":
			configKey = "tokens.directory"
		case "dungeondraft":
			configKey = "dungeondraft.assets-directory"
		default:
			logger.Fatalf(nil, "Unknown directory type \"%s\"", dirType)
		}

		curDir := viper.GetString(configKey)
		isForced, _ := cmd.Flags().GetBool("force")

		if curDir != "" && !isForced {
			logger.Infof("The directory is already configured: %s", curDir)
			logger.Info("Please, use the flag --force flag if you want to override the configuration")
			return
		}

		dir := args[1]
		err := config.CheckDirectory(dir)
		if nil != err {
			logger.Fatalf(err, "Cannot validate directory \"%s\"", dir)
		}

		viper.Set(configKey, dir)
		err = viper.WriteConfig()
		if nil != err {
			logger.Fatal(err, "Error while saving configuration")
		}
	},
}

func init() {
	rootCmd.AddCommand(setDirectoryCmd)

	setDirectoryCmd.Flags().BoolP("force", "f", false, "Force directory override if already set")
}
