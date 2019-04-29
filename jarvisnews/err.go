package jarvisnews

import "errors"

var (
	// ErrNoCrawlerConfig - no crawler config
	ErrNoCrawlerConfig = errors.New("no no crawler config")
	// ErrNoAnkaDBConfig - no ankadb config
	ErrNoAnkaDBConfig = errors.New("no ankadb config")
	// ErrNoHTTPServerAddr - no http server address
	ErrNoHTTPServerAddr = errors.New("no http server address")
)
