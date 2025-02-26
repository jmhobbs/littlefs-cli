package lfs

import (
	"io/fs"

	"tinygo.org/x/tinyfs"
	"tinygo.org/x/tinyfs/littlefs"
)

// Enough of a wrapper for littlefs.LFS that we can use it with fs.WalkDir
type LFS_Sysfs struct {
	fs *littlefs.LFS
}

func NewLFS_Sysfs(fs *littlefs.LFS) LFS_Sysfs {
	return LFS_Sysfs{fs}
}

type lfs_File struct {
	path string
	f    tinyfs.File
	fs   *littlefs.LFS
}

func (f lfs_File) Stat() (fs.FileInfo, error) {
	return f.fs.Stat(f.path)
}

func (f lfs_File) Read(b []byte) (int, error) {
	return f.f.Read(b)
}
func (f lfs_File) Close() error {
	return f.f.Close()
}
func (f lfs_File) ReadDir(n int) ([]fs.DirEntry, error) {
	dirents := []fs.DirEntry{}

	files, err := f.f.Readdir(n)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		dirents = append(dirents, fs.FileInfoToDirEntry(file))
	}

	return dirents, nil
}

func (l LFS_Sysfs) Open(name string) (fs.File, error) {
	tfs, err := l.fs.Open(name)
	if err != nil {
		return nil, err
	}
	return lfs_File{name, tfs, l.fs}, err
}
