package jarvisnews

import (
	"github.com/zhs007/jarviscore"
)

// makeURLKey - make url key
func makeURLKey(url string) string {
	return jarviscore.AppendString("url:", url)
}
