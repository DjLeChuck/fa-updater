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
	"github.com/djlechuck/fa-updater/internal/logger"
	"github.com/spf13/viper"
)

func GrabPack(sessionId string, packs []data.AssetsPack, hideProgress bool) {
	dir := viper.GetString("assetsDirectory")

	for _, pack := range packs {
		downloadPack(dir, sessionId, pack, hideProgress)
	}
}

// removeFile deletes the empty/corrupted file from assets directory
func removeFile(resp *grab.Response) {
	_ = os.Remove(resp.Filename)
}

// downloadPack downloads the given pack into the assets directory
func downloadPack(dir string, sessionId string, pack data.AssetsPack, hideProgress bool) {
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
	logger.Infof("Downloading %v...", pack.Name)
	resp := client.Do(grabReq)
	logger.Info(resp.HTTPResponse.Status)

	if resp.HTTPResponse.StatusCode != 200 {
		removeFile(resp)

		logger.Fatalf(nil, "Cannot access to the URL%s. Please ensure the given cookie is correct.", pack.Path)
	}

	// Start UI loop
	if hideProgress {
		logger.Info("Please wait...")
	} else {
		t := time.NewTicker(10 * time.Second)
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
	}

	// Check for errors
	if err = resp.Err(); err != nil {
		removeFile(resp)

		logger.Fatal(err, "Download failed")
	}

	logger.Infof("Download saved to %s", resp.Filename)
}
