package filesystem

import (
	"os"
	"path/filepath"
	"sort"

	"github.com/go-git/go-billy/v5"
	"github.com/spf13/afero"
)

type BillyFilesystem interface {
	Readlink(link string) (string, error)
	Lstat(filename string) (os.FileInfo, error)
	CapabilityCheck(capabilities billy.Capability) bool
	Symlink(target, link string) error
	// billy.Change
	Lchown(name string, uid, gid int) error
}

type FS interface {
	afero.Fs
	afero.Lstater
	afero.Linker
	BillyFilesystem
}

func ReadDir(fs FS, path string) ([]os.FileInfo, error) {
	f, err := fs.Open(path)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	sort.Slice(list, func(i, j int) bool { return list[i].Name() < list[j].Name() })
	return list, nil
}

// func Readlink(fs afero.Fs, link string) (string, error) {
func Readlink(fs FS, link string) (string, error) {
	return fs.Readlink(link)
}

// func Lstat(fs afero.Fs, filename string) (os.FileInfo, error) {
func Lstat(fs FS, filename string) (os.FileInfo, error) {
	return fs.Lstat(filepath.Clean(filename))
}

func CapabilityCheck(fs FS, capabilities billy.Capability) bool {
	// fsCaps := Capabilities(fs)
	// return fsCaps&capabilities == capabilities
	return true
}

const (
	defaultDirectoryMode = 0755
	// defaultCreateMode    = 0666
)

func createDir(fs FS, fullpath string) error {
	dir := filepath.Dir(fullpath)
	if dir != "." {
		if err := fs.MkdirAll(dir, defaultDirectoryMode); err != nil {
			return err
		}
	}
	return nil
}

func Symlink(fs FS, target, link string) error {
	if err := createDir(fs, link); err != nil {
		return err
	}
	return os.Symlink(target, link)
}
