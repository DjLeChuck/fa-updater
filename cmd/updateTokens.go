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
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/djlechuck/fa-updater/internal/data"
	"github.com/djlechuck/fa-updater/internal/grabber"
	"github.com/djlechuck/fa-updater/internal/logger"
	"github.com/mmcdole/gofeed"
	"github.com/spf13/cobra"
)

var re = regexp.MustCompile(`https://www\.patreon\.com/file\?h=[0-9]+&amp;i=[0-9]+`)

// updateTokensCmd represents the updateTokens command
var updateTokensCmd = &cobra.Command{
	Use:   "updateTokens",
	Short: "Update all tokens",
	Long: `Launch the update process to compare the latest available tokens with the ones in your tokens assets directory.

You will need to give your Patreon session's cookie in order to be able to download the files.`,
	Run: func(cmd *cobra.Command, args []string) {
		app := cmd.Context().Value(AppContextKey).(*application)

		app.config.CheckTokensAssetsDirectory()

		sessionId := app.patreon.GetSessionId()
		hideProgress, _ := cmd.Flags().GetBool("no-progress")

		fp := gofeed.NewParser()
		page := 1
		hasItems := true

		for ok := true; ok; ok = hasItems {
			feed, _ := fp.ParseURL(
				fmt.Sprintf(
					"https://www.forgotten-adventures.net/product-category/tokens/feed/?paged=%d", page,
				),
			)

			if nil == feed {
				hasItems = false
			} else {
				logger.Infof("Processing page %d...", page)

				for _, item := range feed.Items {
					if ignoreItem(item) {
						continue
					}

					link, err := extractPatreonLink(item)
					if nil != err {
						logger.Error(err, "")
						continue
					}

					file := data.PatreonFile{
						Name: item.Title,
						Path: link,
					}

					if app.config.AddTokensDownloadedPack(file.Name) {
						grabber.GrabFile(sessionId, file, hideProgress)
					}
				}

				page++
			}
		}

		app.config.Save()
	},
}

func init() {
	rootCmd.AddCommand(updateTokensCmd)

	updateTokensCmd.Flags().BoolP("no-progress", "n", false, "Hide pack download progression")
}

func ignoreItem(item *gofeed.Item) bool {
	// Ignore bundles
	return strings.Contains(item.Title, "Bundle")
}

func extractPatreonLink(item *gofeed.Item) (string, error) {
	link := re.FindString(item.Description)
	if link == "" {
		return "", errors.New("cannot find Patreon link")
	}

	return strings.ReplaceAll(link, "&amp;", "&"), nil
}
