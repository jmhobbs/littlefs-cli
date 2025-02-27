package block

import "os"

func OpenFileDevice(path string, blockSize, blockCount int64) (*FileDevice, error) {
	f, err := os.OpenFile(path, os.O_RDWR, 0) // TODO: mode
	if err != nil {
		return nil, err
	}

	// sensible defaults
	if blockSize == 0 {
		blockSize = 512
	}

	if blockCount == 0 {
		blockCount = 2048
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

// todo: we should not do this unless it was a "create"
func (b *FileDevice) Prepare() error {
	// pre-allocate space
	_, err := b.file.Seek((b.blocks-1)*b.blockSize, 0)
	if err != nil {
		return err
	}
	_, err = b.file.Write([]byte{0})
	if err != nil {
		return err
	}
	_, err = b.file.Seek(0, 0)
	return err
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
