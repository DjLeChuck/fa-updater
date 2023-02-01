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

package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/djlechuck/fa-updater/internal/logger"
	"github.com/spf13/viper"
)

type Config struct {
	dungeondraftAssetsDirectory string
	tokensAssetsDirectory       string
}

var ErrInvalidAssetsDirectory = errors.New("not a valid directory")

func (cfg *Config) CheckDungeondraftAssetsDirectory() {
	err := checkDirectory(viper.GetString("dungeondraft.assets-directory"))
	if nil != err {
		logger.Fatal(err, "Cannot get assets directory")
	}
}

func CheckDirectory(dir string) error {
	return checkDirectory(dir)
}

func checkDirectory(dir string) error {
	if dir == "" {
		return ErrInvalidAssetsDirectory
	}

	if stat, err := os.Stat(dir); nil != err || !stat.IsDir() {
		fmt.Print(err)
		return ErrInvalidAssetsDirectory
	}

	return nil
}
