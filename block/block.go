package block

import "os"

type Device struct {
	Name      string
	Blocks    int
	BlockSize int
}

func New(f *os.File, blockSize, blockCount int64) (*FileBlockDevice, error) {
	// pre-allocate space
	/*
		_, err := f.Seek(blockSize*blockCount-1, 0)
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
	*/

	return &FileBlockDevice{
		file:       f,
		blocksize:  blockSize,
		blockcount: blockCount,
	}, nil
}

type FileBlockDevice struct {
	file       *os.File
	blocksize  int64
	blockcount int64
}

func (bd *FileBlockDevice) Size() int64 {
	return bd.blocksize * bd.blockcount
}

func (bd *FileBlockDevice) WriteBlockSize() int64 {
	return bd.blocksize
}

func (bd *FileBlockDevice) EraseBlockSize() int64 {
	return bd.blocksize
}

func (bd *FileBlockDevice) EraseBlocks(start, count int64) error {
	_, err := bd.file.WriteAt(make([]byte, 0, (start+count*bd.blocksize-1)), start*bd.blocksize)
	return err
}

func (bd *FileBlockDevice) ReadAt(p []byte, off int64) (n int, err error) {
	return bd.file.ReadAt(p, off)
}

func (bd *FileBlockDevice) WriteAt(p []byte, off int64) (n int, err error) {
	return bd.file.WriteAt(p, off)
}
