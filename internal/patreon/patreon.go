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

package patreon

import (
	"fmt"

	"github.com/djlechuck/fa-updater/internal/clipboard"
	"github.com/djlechuck/fa-updater/internal/logger"
)

type Patreon struct {
	sessionId string
}

func (patreon *Patreon) GetSessionId() string {
	if patreon.sessionId != "" {
		return patreon.sessionId
	}

	logger.Info("Please, look at the cookies on the Patreon page and copy the value of the one named \"session_id\" in the clipboard (CTRL+C or ⌘+C), then press ENTER. It should looks like a random string: LC2A4j7WAJe4cjR5Oeicycf4YmlEfQsNB_yqwYiWuh8")
	fmt.Scanln()

	patreon.sessionId = clipboard.ReadString()

	return patreon.sessionId
}
