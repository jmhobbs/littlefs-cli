package lfs

import (
	"os"

	"github.com/jmhobbs/littlefs-cli/block"
	"github.com/jmhobbs/littlefs-cli/path"

	"tinygo.org/x/tinyfs/littlefs"
)

func open(file path.Path, flag int) (*os.File, *littlefs.LFS, error) {
	f, err := os.OpenFile(file.Path, flag, 0666)
	if err != nil {
		return nil, nil, err
	}

	// todo: bs/bc for images
	bd, err := block.New(f, 512, 2048) // *blockSize, *blockCount)
	if err != nil {
		f.Close()
		return nil, nil, err
	}

	lfs := littlefs.New(bd)
	lfs.Configure(&littlefs.Config{
		CacheSize:     uint32(512), //bd.blocksize),
		LookaheadSize: uint32(512), //bd.blocksize),
		BlockCycles:   100,         // todo: wtf is this
	})

	err = lfs.Mount()
	if err != nil {
		f.Close()
		return nil, nil, err
	}

	return f, lfs, nil
}

type WithLFSFunc func(*littlefs.LFS) error

func WithReadOnly(file path.Path, cb WithLFSFunc) error {
	return with(os.O_RDONLY, file, cb)
}

func WithReadWrite(file path.Path, cb WithLFSFunc) error {
	return with(os.O_RDWR, file, cb)
}

func with(flag int, file path.Path, cb func(*littlefs.LFS) error) error {
	f, lfs, err := open(file, flag)
	if err != nil {
		return err
	}
	defer lfs.Unmount()
	defer f.Close()

	return cb(lfs)
}
