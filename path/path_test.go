package path_test

import (
	"testing"

	"github.com/jmhobbs/littlefs-cli/path"
	"github.com/stretchr/testify/assert"
)

func Test_ParsePath(t *testing.T) {
	t.Parallel()
	tests := []struct {
		Path     string
		Input    string
		Expected path.Path
	}{
		{
			"Disk block device",
			"/dev/disk6:/path/to/file",
			path.Path{Type: path.BLOCK, Path: "/dev/disk6", VolumePath: "/path/to/file"},
		},
		{
			"Filesystem",
			"/dev/disk6",
			path.Path{Type: path.FILESYSTEM, Path: "/dev/disk6"},
		},
		{
			"Image block device",
			"~/Desktop/disk.img:/",
			path.Path{Type: path.BLOCK, Path: "~/Desktop/disk.img", VolumePath: "/"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Path, func(t *testing.T) {
			t.Parallel()
			actual := path.Parse(tt.Input)
			assert.Equal(t, tt.Expected, actual)
		})
	}
}
