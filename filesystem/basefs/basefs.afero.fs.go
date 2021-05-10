package basefs

import (
	"os"
	"time"

	"github.com/spf13/afero"
)

func (fs BaseFS) Create(name string) (afero.File, error) {
	return fs.source.Create(name)
}

func (fs BaseFS) Mkdir(name string, perm os.FileMode) error {
	return fs.source.Mkdir(name, perm)
}

func (fs BaseFS) MkdirAll(path string, perm os.FileMode) error {
	return fs.source.MkdirAll(path, perm)
}

func (fs BaseFS) Open(name string) (afero.File, error) {
	return fs.source.Open(name)
}

func (fs BaseFS) OpenFile(name string, flag int, perm os.FileMode) (afero.File, error) {
	return fs.source.OpenFile(name, flag, perm)
}

func (fs BaseFS) Remove(name string) error {
	return fs.source.Remove(name)
}

func (fs BaseFS) RemoveAll(path string) error {
	return fs.source.RemoveAll(path)
}

func (fs BaseFS) Rename(oldname, newname string) error {
	return fs.source.Rename(oldname, newname)
}

func (fs BaseFS) Stat(name string) (os.FileInfo, error) {
	return fs.source.Stat(name)
}

func (fs BaseFS) Name() string { return "GoNFS_BaseFS" }

func (fs BaseFS) Chmod(name string, mode os.FileMode) error {
	return fs.source.Chmod(name, mode)
}

func (fs BaseFS) Chown(name string, uid, gid int) error {
	return fs.source.Chown(name, uid, gid)
}

func (fs BaseFS) Chtimes(name string, atime time.Time, mtime time.Time) error {
	return fs.source.Chtimes(name, atime, mtime)
}
