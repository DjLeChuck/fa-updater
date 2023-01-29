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

package grabber

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/cavaliergopher/grab/v3"
	"github.com/djlechuck/fa-updater/internal/data"
	"github.com/spf13/viper"
)

func GrabPack(sessionId string, packs []data.AssetsPack) {
	dir := viper.GetString("assetsDirectory")

	for _, pack := range packs {
		downloadPack(dir, sessionId, pack)
	}
}

// removeFile deletes the empty/corrupted file from assets directory
func removeFile(resp *grab.Response) {
	_ = os.Remove(resp.Filename)
}

// downloadPack downloads the given pack into the assets directory
func downloadPack(dir string, sessionId string, pack data.AssetsPack) {
	// Create client
	req, err := http.NewRequest("GET", pack.Path, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Cookie", fmt.Sprintf("session_id=%s;", sessionId))

	client := grab.NewClient()
	grabReq := &grab.Request{
		HTTPRequest: req,
		Filename:    dir,
	}

	// Start download
	fmt.Printf("Downloading %v...\n", pack.Name)
	resp := client.Do(grabReq)
	fmt.Printf("  %v\n", resp.HTTPResponse.Status)

	if resp.HTTPResponse.StatusCode != 200 {
		fmt.Fprintln(
			os.Stderr, "Cannot access to the URL", pack.Path, ". Please ensure the given cookie is correct.",
		)

		removeFile(resp)

		os.Exit(1)
	}

	// Start UI loop
	t := time.NewTicker(time.Second)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			fmt.Printf(
				"  transferred %v / %v bytes (%.2f%%)\n",
				resp.BytesComplete(),
				resp.Size(),
				100*resp.Progress(),
			)

		case <-resp.Done:
			// Download is complete
			break Loop
		}
	}

	// Check for errors
	if err = resp.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Download failed:", err)

		removeFile(resp)

		os.Exit(1)
	}

	fmt.Printf("Download saved to %v \n\n", resp.Filename)
}
