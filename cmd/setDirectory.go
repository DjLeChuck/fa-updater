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
	Use:   "setDirectory [path]",
	Short: "Define the directory which contains your assets",
	Long:  "Define the directory which contains your assets. It will be used by the updateAssets command to get newer versions if exists.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		curDir := viper.GetString("assetsDirectory")
		isForced, _ := cmd.Flags().GetBool("force")

		if curDir != "" && !isForced {
			logger.Infof("The assets directory is already configured: %s", curDir)
			logger.Info("Please, use the flag --force flag if you want to override the configuration.")
			return
		}

		dir := args[0]
		err := config.CheckDirectory(dir)
		if nil != err {
			logger.Fatalf(err, "Cannot validate directory \"%s\"", dir)
		}

		viper.Set("assetsDirectory", dir)
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
