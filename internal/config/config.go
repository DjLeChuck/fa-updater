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

var thumbnailsVersions []string
var downloadedPacks []string

type Dungeondraft struct {
	AssetsDirectory    string
	ThumbnailsVersions map[string]bool
}

type Tokens struct {
	AssetsDirectory string
	DownloadedPacks map[string]bool
}

type Config struct {
	Dungeondraft *Dungeondraft
	Tokens       *Tokens
}

var ErrInvalidAssetsDirectory = errors.New("not a valid directory")

func (cfg *Config) CheckDungeondraftAssetsDirectory() {
	err := checkDirectory(viper.GetString("dungeondraft.assets-directory"))
	if nil != err {
		logger.Fatal(err, "Cannot get assets directory")
	}
}

func (cfg *Config) CheckTokensAssetsDirectory() {
	err := checkDirectory(viper.GetString("tokens.directory"))
	if nil != err {
		logger.Fatal(err, "Cannot get tokens directory")
	}
}

func (cfg *Config) AddDungeondraftThumbnailVersion(version string) bool {
	return addValueInList(
		version, "dungeondraft.thumbnails-versions", &thumbnailsVersions, &cfg.Dungeondraft.ThumbnailsVersions,
	)
}

func (cfg *Config) AddTokensDownloadedPack(pack string) bool {
	return addValueInList(pack, "tokens.packs", &downloadedPacks, &cfg.Tokens.DownloadedPacks)
}

func (cfg *Config) Save() {
	err := viper.WriteConfig()
	if nil != err {
		logger.Fatal(err, "Error while saving configuration")
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

func addValueInList(value string, configKey string, list *[]string, listMap *map[string]bool) bool {
	*list = initExistingList(configKey, listMap)

	if (*listMap)[value] {
		return false
	}

	*list = append(*list, value)
	(*listMap)[value] = true

	viper.Set(configKey, list)

	return true
}

func initExistingList(configKey string, configMap *map[string]bool) []string {
	list := viper.GetStringSlice(configKey)

	for _, item := range list {
		(*configMap)[item] = true
	}

	return list
}
