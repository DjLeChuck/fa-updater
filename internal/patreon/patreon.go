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
	if "" != patreon.sessionId {
		return patreon.sessionId
	}

	logger.Info("Please, look at the cookies on the Patreon page and copy the value of the one named \"session_id\" in the clipboard (CTRL+C or ⌘+C), then press ENTER. It should looks like a random string: LC2A4j7WAJe4cjR5Oeicycf4YmlEfQsNB_yqwYiWuh8")
	fmt.Scanln()

	patreon.sessionId = clipboard.ReadString()

	return patreon.sessionId
}
