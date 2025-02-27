package lfs

import (
	"os"

	"github.com/jmhobbs/littlefs-cli/block"
	"github.com/jmhobbs/littlefs-cli/path"

	"tinygo.org/x/tinyfs"
	"tinygo.org/x/tinyfs/littlefs"
)

type Config struct {
	BlockSize int64
	Blocks    int64
}

// Open a filesystem and file on it
// Caller is responsible for calling Close() and Unmount() on returned values
func OpenPath(file path.Path, flag int, blockSize, blockCount int64) (block.Device, *littlefs.LFS, tinyfs.File, error) {
	bd, lfs, err := Open(file.Path, flag&os.O_RDWR, blockSize, blockCount)
	if err != nil {
		return nil, nil, nil, err
	}

	f, err := lfs.OpenFile(file.VolumePath, flag)
	if err != nil {
		lfs.Unmount()
		bd.Close()
		return nil, nil, nil, err
	}

	return bd, lfs, f, nil
}

// Open a filesystem
// Caller is responsible for calling Close() and Unmount() on returned values
func Open(file string, flag int, defaultBlockSize, defaultBlockCount int64) (block.Device, *littlefs.LFS, error) {
	bd, err := block.Open(file, defaultBlockSize, defaultBlockCount)
	if err != nil {
		return nil, nil, err
	}
	if err := bd.Prepare(); err != nil {
		bd.Close()
		return nil, nil, err
	}

	lfs := littlefs.New(bd)
	lfs.Configure(&littlefs.Config{
		CacheSize:     512,
		LookaheadSize: 512,
		BlockCycles:   100,
	})

	err = lfs.Mount()
	if err != nil {
		bd.Close()
		return nil, nil, err
	}

	return bd, lfs, nil
}

type WithLFSFunc func(*littlefs.LFS) error

func WithReadOnly(file path.Path, defaultBlockSize, defaultBlockCount int64, cb WithLFSFunc) error {
	return with(os.O_RDONLY, file, defaultBlockSize, defaultBlockCount, cb)
}

func WithReadWrite(file path.Path, defaultBlockSize, defaultBlockCount int64, cb WithLFSFunc) error {
	return with(os.O_RDWR, file, defaultBlockSize, defaultBlockCount, cb)
}

func with(flag int, file path.Path, defaultBlockSize, defaultBlockCount int64, cb func(*littlefs.LFS) error) error {
	f, lfs, err := Open(file.Path, flag, defaultBlockSize, defaultBlockCount)
	if err != nil {
		return err
	}
	defer lfs.Unmount()
	defer f.Close()

	return cb(lfs)
}
