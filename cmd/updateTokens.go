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
	"fmt"

	"github.com/mmcdole/gofeed"
	"github.com/spf13/cobra"
)

// updateTokensCmd represents the updateTokens command
var updateTokensCmd = &cobra.Command{
	Use:   "updateTokens",
	Short: "Update all tokens",
	Long: `Launch the update process to compare latest available tokens with the ones in your tokens assets directory.

You will need to give your Patreon session's cookie in order to be able to download the files.`,
	Run: func(cmd *cobra.Command, args []string) {
		app := cmd.Context().Value("app").(*application)

		app.config.CheckTokensAssetsDirectory()

		app.patreon.GetSessionId()

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
				for _, item := range feed.Items {
					fmt.Println(item.Title)
				}

				page++
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(updateTokensCmd)
}
