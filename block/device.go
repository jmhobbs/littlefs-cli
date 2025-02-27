package block

import (
	"strings"

	"tinygo.org/x/tinyfs"
)

type Device interface {
	tinyfs.BlockDevice

	Prepare() error
	Close() error
}

func Open(path string, defaultBlockSize, defaultBlockCount int64) (Device, error) {
	// TODO: better detection, cross platform
	if strings.HasPrefix(path, "/dev/") {
		return OpenHardwareDevice(path)
	}

	// todo: detect block count for images

	return OpenFileDevice(path, defaultBlockSize, defaultBlockCount)
}
