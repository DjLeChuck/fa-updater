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

package dungeondraft

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/djlechuck/fa-updater/internal/config"
	"github.com/djlechuck/fa-updater/internal/logger"
	"github.com/djlechuck/fa-updater/internal/unzip"
)

func UnzipThumbnails(src string) error {
	return unzip.Unzip(src, getDataDirectory())
}

// getDataDirectory returns the directory storing Dungeondraft data
func getDataDirectory() string {
	// macOS: ~/Library/Application Support/Dungeondraft/
	// Windows: C:\Users\USERNAME\AppData\Roaming\Dungeondraft\
	// Linux: ~/.local/share/Dungeondraft/

	configDir, err := os.UserConfigDir()
	if nil != err {
		logger.Fatal(err, "Cannot detect user config directory")
	}

	var dir string
	osName := runtime.GOOS
	switch osName {
	case "windows":
		// cascade
	case "darwin":
		dir = filepath.Join(configDir, "Dungeondraft")
	case "linux":
		homeDir, err := os.UserHomeDir()
		if nil != err {
			logger.Fatal(err, "Cannot detect user home directory")
		}

		dir = filepath.Join(homeDir, ".local/share/Dungeondraft")
	default:
		logger.Fatalf(nil, "Cannot work with OS %s", osName)
	}

	err = config.CheckDirectory(dir)
	if nil != err {
		logger.Fatalf(err, "Cannot validate directory \"%s\"", dir)
	}

	return dir
}
