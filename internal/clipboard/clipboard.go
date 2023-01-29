package clipboard

import (
	"strings"

	"golang.design/x/clipboard"
)

func ReadBytes() []byte {
	return readBytes()
}

func ReadString() string {
	return strings.Trim(string(readBytes()), " ")
}

func readBytes() []byte {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	cb := clipboard.Read(clipboard.FmtText)
	clipboard.Write(clipboard.FmtText, []byte(""))

	return cb
}
