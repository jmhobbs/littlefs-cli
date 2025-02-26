package path

import "strings"

type Type uint8

const (
	BLOCK Type = iota
	FILESYSTEM
)

type Path struct {
	Type Type
	// The path of the block device or file
	Path string
	// The path on a block device
	VolumePath string
}

func Parse(input string) Path {
	split := strings.SplitN(input, ":", 2)
	if len(split) == 1 {
		return Path{Type: FILESYSTEM, Path: input}
	}
	return Path{Type: BLOCK, Path: split[0], VolumePath: split[1]}
}
