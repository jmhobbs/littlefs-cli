package block

import (
	"strings"

	"tinygo.org/x/tinyfs"
)

type Device interface {
	tinyfs.BlockDevice

	Close() error
}

func Open(path string, defaultBlockSize, defaultBlockCount int64) (Device, error) {
	// TODO: better detection, cross platform
	if strings.HasPrefix(path, "/dev/") {
		return OpenHardwareDevice(path)
	}
	return OpenFileDevice(path, defaultBlockSize, defaultBlockCount)
}

func Create(path string, blockSize, blockCount int64) (Device, error) {
	if strings.HasPrefix(path, "/dev/") {
		return OpenHardwareDevice(path)
	}
	return CreateFileDevice(path, blockSize, blockCount)
}
