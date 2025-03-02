package block

import (
	"fmt"
	"os"
)

func CreateFileDevice(path string, blockSize, blockCount int64) (*FileDevice, error) {
	if blockSize == 0 {
		blockSize = 512
	}
	if blockCount == 0 {
		return nil, fmt.Errorf("block count must be specified when creating an image")
	}

	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	// pre-allocate space
	_, err = f.Seek((blockCount*blockSize)-1, 0)
	if err != nil {
		return nil, err
	}
	_, err = f.Write([]byte{0})
	if err != nil {
		return nil, err
	}
	_, err = f.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	return &FileDevice{
		file:      f,
		blockSize: blockSize,
		blocks:    blockCount,
	}, nil
}

func OpenFileDevice(path string, blockSize, blockCount int64) (*FileDevice, error) {
	f, err := os.OpenFile(path, os.O_RDWR, 0) // TODO: optional read only mode
	if err != nil {
		return nil, err
	}

	// sensible default
	if blockSize == 0 {
		blockSize = 512
	}

	if blockCount == 0 {
		fstat, err := f.Stat()
		if err != nil {
			f.Close()
			return nil, err
		}
		if fstat.Size()%blockSize != 0 {
			f.Close()
			return nil, fmt.Errorf("file size %d is not a multiple of block size %d", fstat.Size(), blockSize)
		}
		blockCount = fstat.Size() / blockSize
	}

	return &FileDevice{
		file:      f,
		blockSize: blockSize,
		blocks:    blockCount,
	}, nil
}

type FileDevice struct {
	file      *os.File
	blockSize int64
	blocks    int64
}

func (b *FileDevice) Close() error {
	return b.file.Close()
}

func (bd *FileDevice) Size() int64 {
	return bd.blockSize * bd.blocks
}

func (bd *FileDevice) WriteBlockSize() int64 {
	return bd.blockSize
}

func (bd *FileDevice) EraseBlockSize() int64 {
	return bd.blockSize
}

func (bd *FileDevice) EraseBlocks(start, count int64) error {
	_, err := bd.file.WriteAt(make([]byte, 0, (start+count*bd.blockSize-1)), start*bd.blockSize)
	return err
}

func (bd *FileDevice) ReadAt(p []byte, off int64) (n int, err error) {
	return bd.file.ReadAt(p, off)
}

func (bd *FileDevice) WriteAt(p []byte, off int64) (n int, err error) {
	return bd.file.WriteAt(p, off)
}
