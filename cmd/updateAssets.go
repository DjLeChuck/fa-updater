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
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/djlechuck/fa-updater/internal/clipboard"
	"github.com/djlechuck/fa-updater/internal/config"
	"github.com/djlechuck/fa-updater/internal/data"
	"github.com/djlechuck/fa-updater/internal/grabber"
	"github.com/djlechuck/fa-updater/internal/logger"
	"github.com/djlechuck/fa-updater/internal/pack"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const PatreonPostLink = "https://www.patreon.com/posts/56375276"
const PatreonPackLinkPrefix = "https://www.patreon.com/file?h=17713082"

// updateAssetsCmd represents the updateAssets command
var updateAssetsCmd = &cobra.Command{
	Use:   "updateAssets",
	Short: "Launch the update process",
	Long: `Launch the update process to compare latest available packs with the ones in your assets directory.

First, you will need to get the Patreon page content, then give your Patreon session's cookie in order to be able to download the files.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := config.CheckConfigAssetsDirectory()
		if nil != err {
			logger.Fatal(err, "Cannot get assets directory")
		}

		logger.Infof(
			"Go on %s with your browser. Display the source of the page (CTRL+U or ⌘ +U) and copy it in the clipboard (CTRL+A and CTRL+C or ⌘ +A and ⌘ +C), then go back here and press ENTER.",
			PatreonPostLink,
		)
		fmt.Scanln()

		// Parse clipboard content
		doc, err := htmlquery.Parse(bytes.NewReader(clipboard.ReadBytes()))
		if nil != err {
			logger.Fatalf(err, "Cannot parse Patreon post")
		}

		// XPath all the file URLs and get packs
		list := htmlquery.Find(doc, fmt.Sprintf("//a[starts-with(@href, '%s')]", PatreonPackLinkPrefix))
		if len(list) == 0 {
			logger.Fatal(nil, "Cannot find any packs. Ensure you have correctly copy the page source code.")
		}

		var packs []data.AssetsPack
		thumbnailsPartTwoPassed := false
		for _, n := range list {
			if thumbnailsPartTwoPassed {
				break
			}

			name := strings.Trim(htmlquery.InnerText(n), " ")
			isThumbnails := strings.HasPrefix(name, "THUMBNAILS_")

			if !isThumbnails {
				name = fmt.Sprintf("FA_%s.dungeondraft_pack", name)
			}

			packs = append(
				packs, data.AssetsPack{
					Name:       name,
					Path:       htmlquery.SelectAttr(n, "href"),
					Thumbnails: isThumbnails,
					IsLocal:    false,
				},
			)

			thumbnailsPartTwoPassed = strings.HasPrefix(name, "THUMBNAILS_Part2_")
		}

		logger.Infof("%d packs found. Comparing to your assets directory...", len(packs))
		dir := viper.GetString("assetsDirectory")

		files, err := ioutil.ReadDir(dir)
		if err != nil {
			logger.Fatal(err, "Cannot read assets directory")
		}

		var localPacks []data.AssetsPack
		for _, file := range files {
			localPacks = append(
				localPacks, data.AssetsPack{
					Name:    file.Name(),
					Path:    filepath.Join(dir, file.Name()),
					IsLocal: true,
				},
			)
		}

		newPacks := pack.PackDiff(packs, localPacks)

		if len(newPacks) == 0 {
			logger.Info("All your packs are already up-to-date!")
			os.Exit(0)
		}

		logger.Infof("There are %d packs to download.", len(newPacks))
		logger.Info("Please, look at the cookies on the Patreon page and copy the value of the one named \"session_id\" in the clipboard (CTRL+C or ⌘ +C), then press ENTER. It should looks like a random string: LC2A4j7WAJe4cjR5Oeicycf4YmlEfQsNB_yqwYiWuh8")
		fmt.Scanln()

		hideProgress, _ := cmd.Flags().GetBool("no-progress")
		grabber.GrabPack(clipboard.ReadString(), newPacks, hideProgress)

		oldPacks := pack.PackDiff(localPacks, packs)

		if 0 < len(oldPacks) {
			logger.Info("Removing old packs...")

			for _, oldPack := range oldPacks {
				logger.Info(oldPack.Name)
				err := os.Remove(oldPack.Path)
				if nil != err {
					logger.Error(err, "Error while removing the file")
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(updateAssetsCmd)

	updateAssetsCmd.Flags().BoolP("no-progress", "n", false, "Hide pack download progression")
}
